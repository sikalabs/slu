package argocd_utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sikalabs/slu/utils/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const argoCDRefreshAnnotation = "argocd.argoproj.io/refresh"

var applicationGVR = schema.GroupVersionResource{
	Group:    "argoproj.io",
	Version:  "v1alpha1",
	Resource: "applications",
}

// ArgoCDRefresh requests a refresh of an ArgoCD Application by setting the
// argocd.argoproj.io/refresh annotation (refreshType is "normal" or "hard").
// ArgoCD removes the annotation once the refresh is done; if wait > 0, it
// waits up to that long for that to happen.
func ArgoCDRefresh(
	ctx context.Context,
	namespace string,
	appName string,
	refreshType string,
	wait time.Duration,
) error {
	restconfig, err := kubeRestConfig()
	if err != nil {
		return err
	}
	dynamicClient, err := dynamic.NewForConfig(restconfig)
	if err != nil {
		return err
	}
	apps := dynamicClient.Resource(applicationGVR).Namespace(namespace)

	patch := fmt.Sprintf(`{"metadata":{"annotations":{%q:%q}}}`, argoCDRefreshAnnotation, refreshType)
	_, err = apps.Patch(ctx, appName, types.MergePatchType, []byte(patch), metav1.PatchOptions{})
	if err != nil {
		return err
	}
	if wait <= 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, wait)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		app, err := apps.Get(ctx, appName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if _, ok := app.GetAnnotations()[argoCDRefreshAnnotation]; !ok {
			return nil
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for ArgoCD to refresh application '%s'", appName)
		case <-ticker.C:
		}
	}
}

func kubeRestConfig() (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
}

func ArgoCDGetInitialPassword(namespace string) string {
	clientset, _, _ := k8s.KubernetesClient()

	secretClient := clientset.CoreV1().Secrets(namespace)

	// Try argocd secret first
	secret, err := secretClient.Get(context.TODO(), "argocd-initial-admin-secret", metav1.GetOptions{})
	if err == nil {
		return string(secret.Data["password"])
	}

	// If argocd namespace fails and we're using default namespace, try openshift-gitops
	if namespace == "argocd" {
		secretClient = clientset.CoreV1().Secrets("openshift-gitops")
		secret, err = secretClient.Get(context.TODO(), "openshift-gitops-cluster", metav1.GetOptions{})
		if err == nil {
			return string(secret.Data["admin.password"])
		}
	}

	log.Fatalf("could not find initial password in namespace '%s' or 'openshift-gitops'", namespace)
	return ""
}

func ArgoCDGetDomain(namespace string) (string, error) {
	clientset, _, err := k8s.KubernetesClient()
	if err != nil {
		return "", err
	}

	// Try argocd ingress first
	ingressClient := clientset.NetworkingV1().Ingresses(namespace)
	ingress, err := ingressClient.Get(context.TODO(), "argocd-server", metav1.GetOptions{})
	if err == nil {
		rule := ingress.Spec.Rules[0]
		return rule.Host, nil
	}

	// If argocd namespace fails and we're using default namespace, try openshift-gitops route
	if namespace == "argocd" {
		domain, err := getOpenShiftGitOpsDomain()
		if err == nil {
			return domain, nil
		}
	}

	return "", fmt.Errorf("could not find ArgoCD in namespace '%s' or OpenShift GitOps in namespace 'openshift-gitops'", namespace)
}

func getOpenShiftGitOpsDomain() (string, error) {
	restconfig, err := kubeRestConfig()
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

	return host, nil
}

func ArgoCDGetDomainOrDie(namespace string) string {
	domain, err := ArgoCDGetDomain(namespace)
	if err != nil {
		log.Fatal(err)
	}
	return domain
}
