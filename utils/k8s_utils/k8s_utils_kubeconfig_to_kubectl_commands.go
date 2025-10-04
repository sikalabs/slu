package k8s_utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/sikalabs/slu/internal/error_utils"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func KubeconfigToKubectlCommandsOrDie(kubeconfigPath string) {
	config, err := loadKubeconfig(kubeconfigPath)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	ctxName := config.CurrentContext
	ctx, ok := config.Contexts[ctxName]
	error_utils.HandleNotOK(ok, fmt.Sprintf("Context %q not found", ctxName))

	cluster, ok := config.Clusters[ctx.Cluster]
	error_utils.HandleNotOK(ok, fmt.Sprintf("Cluster %q not found", ctx.Cluster))

	user, ok := config.AuthInfos[ctx.AuthInfo]
	error_utils.HandleNotOK(ok, fmt.Sprintf("AuthInfo %q not found", ctx.AuthInfo))

	fmt.Println("# Commands to replicate kubeconfig setup (inline cert/key data, no files)")

	// Cluster: server + tls options
	fmt.Printf("kubectl config set-cluster %s --server=%s", ctx.Cluster, cluster.Server)
	if cluster.InsecureSkipTLSVerify {
		fmt.Printf(" --insecure-skip-tls-verify=true")
	}
	fmt.Println()

	// Cluster CA data (embed directly)
	if len(cluster.CertificateAuthorityData) > 0 {
		caB64 := base64.StdEncoding.EncodeToString(cluster.CertificateAuthorityData)
		fmt.Printf("kubectl config set clusters.%s.certificate-authority-data '%s'\n", ctx.Cluster, caB64)
	} else if cluster.CertificateAuthority != "" {
		// Read existing CA file and embed it (still no new files created)
		b, err := os.ReadFile(cluster.CertificateAuthority)
		if err != nil {
			log.Fatalf("Failed to read certificate-authority file %q: %v", cluster.CertificateAuthority, err)
		}
		fmt.Printf("kubectl config set clusters.%s.certificate-authority-data '%s'\n", ctx.Cluster, base64.StdEncoding.EncodeToString(b))
	}

	// User credentials
	printedSetCred := false
	if user.Token != "" || (user.Username != "" && user.Password != "") {
		fmt.Printf("kubectl config set-credentials %s", ctx.AuthInfo)
		if user.Token != "" {
			fmt.Printf(" --token=%s", user.Token)
		}
		if user.Username != "" && user.Password != "" {
			fmt.Printf(" --username=%s --password=%s", user.Username, user.Password)
		}
		fmt.Println()
		printedSetCred = true
	}

	// Client cert/key (embed directly)
	if len(user.ClientCertificateData) > 0 {
		ccB64 := base64.StdEncoding.EncodeToString(user.ClientCertificateData)
		fmt.Printf("kubectl config set users.%s.client-certificate-data '%s'\n", ctx.AuthInfo, ccB64)
	} else if user.ClientCertificate != "" {
		b, err := os.ReadFile(user.ClientCertificate)
		if err != nil {
			log.Fatalf("Failed to read client certificate %q: %v", user.ClientCertificate, err)
		}
		fmt.Printf("kubectl config set users.%s.client-certificate-data '%s'\n", ctx.AuthInfo, base64.StdEncoding.EncodeToString(b))
	}

	if len(user.ClientKeyData) > 0 {
		ckB64 := base64.StdEncoding.EncodeToString(user.ClientKeyData)
		fmt.Printf("kubectl config set users.%s.client-key-data '%s'\n", ctx.AuthInfo, ckB64)
	} else if user.ClientKey != "" {
		b, err := os.ReadFile(user.ClientKey)
		if err != nil {
			log.Fatalf("Failed to read client key %q: %v", user.ClientKey, err)
		}
		fmt.Printf("kubectl config set users.%s.client-key-data '%s'\n", ctx.AuthInfo, base64.StdEncoding.EncodeToString(b))
	}

	// If we didn't call set-credentials above, ensure the user exists (empty call)
	if !printedSetCred {
		fmt.Printf("kubectl config set-credentials %s\n", ctx.AuthInfo)
	}

	// Optional: auth-provider or exec plugins aren’t translated here.
	if user.AuthProvider != nil || user.Exec != nil {
		fmt.Println("# NOTE: This kubeconfig uses an auth provider/exec plugin; token refresh logic isn’t reproduced by these commands.")
	}

	// Context + namespace
	fmt.Printf("kubectl config set-context %s --cluster=%s --user=%s", ctxName, ctx.Cluster, ctx.AuthInfo)
	if ctx.Namespace != "" {
		fmt.Printf(" --namespace=%s", ctx.Namespace)
	}
	fmt.Println()

	fmt.Printf("kubectl config use-context %s\n", ctxName)
}

func loadKubeconfig(path string) (*clientcmdapi.Config, error) {
	return clientcmd.LoadFromFile(path)
}
