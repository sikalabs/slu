package open

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/browser"
	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "Open ArgoCD in browser",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, _ := k8s.KubernetesClient()

		// Try argocd namespace first
		ingressClient := clientset.NetworkingV1().Ingresses(FlagNamespace)
		ingress, err := ingressClient.Get(context.TODO(), "argocd-server", metav1.GetOptions{})
		if err == nil {
			rule := ingress.Spec.Rules[0]
			url := "https://" + rule.Host + rule.HTTP.Paths[0].Path
			fmt.Println(url)
			browser.OpenURL(url)
			return
		}

		// If argocd namespace fails and we're using default namespace, try openshift-gitops
		if FlagNamespace == "argocd" {
			url, err := getOpenShiftGitOpsURL()
			if err == nil {
				fmt.Println(url)
				browser.OpenURL(url)
				return
			}
		}

		log.Fatal("Could not find ArgoCD in namespace 'argocd' or OpenShift GitOps in namespace 'openshift-gitops'")
	},
}

func getOpenShiftGitOpsURL() (string, error) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	restconfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return "", err
	}

	dynamicClient, err := dynamic.NewForConfig(restconfig)
	if err != nil {
		return "", err
	}

	routeGVR := schema.GroupVersionResource{
		Group:    "route.openshift.io",
		Version:  "v1",
		Resource: "routes",
	}

	route, err := dynamicClient.Resource(routeGVR).Namespace("openshift-gitops").Get(
		context.TODO(),
		"openshift-gitops-server",
		metav1.GetOptions{},
	)
	if err != nil {
		return "", err
	}

	spec, ok := route.Object["spec"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid route spec")
	}

	host, ok := spec["host"].(string)
	if !ok {
		return "", fmt.Errorf("invalid route host")
	}

	return "https://" + host, nil
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
		"ArgoCD Namespace",
	)
}
