package k8s_scripts

import (
	"fmt"
	"os"

	"github.com/sikalabs/slu/utils/exec_utils"
)

func ConnectFromVault(vaultAddress string, clusterName string) {
	tmpKubeconfig := "/tmp/kubeconfig"
	secretPath := fmt.Sprintf("secret/kubeconfigs/%s", clusterName)

	// Vault login with OIDC
	fmt.Println("Logging in to Vault...")
	err := exec_utils.ExecOut("vault", "login",
		"-address", vaultAddress,
		"-method=oidc")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging in to Vault: %v\n", err)
		os.Exit(1)
	}

	// Copy kubeconfig from Vault
	fmt.Printf("Downloading kubeconfig for cluster '%s' from Vault...\n", clusterName)
	err = exec_utils.ExecOut("slu", "vault", "copy-file-from-vault",
		"--vault-address", vaultAddress,
		"--secret-path", secretPath,
		"--file-path", tmpKubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying kubeconfig from Vault: %v\n", err)
		os.Exit(1)
	}

	// Add kubeconfig
	fmt.Println("Adding kubeconfig...")
	err = exec_utils.ExecOut("slu", "k8s", "kubeconfig", "add", "-p", tmpKubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding kubeconfig: %v\n", err)
		// Clean up temp file before exiting
		os.Remove(tmpKubeconfig)
		os.Exit(1)
	}

	// Clean up temp file
	fmt.Println("Cleaning up temporary files...")
	err = os.Remove(tmpKubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not remove temporary kubeconfig file: %v\n", err)
	}

	fmt.Println("Successfully connected to cluster!")
}
