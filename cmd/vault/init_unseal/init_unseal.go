package init_unseal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/spf13/cobra"
)

var FlagFile string
var FlagNamespace string
var FlagSelector string
var FlagKeyShares int
var FlagKeyThreshold int

var Cmd = &cobra.Command{
	Use:   "init-unseal",
	Short: "Initialize and unseal Vault running in Kubernetes",
	Long: `Initialize Vault (vault operator init) on the first Vault pod, save the init JSON
(unseal keys and root token) to a local file and unseal all Vault pods.`,
	Example: `slu vault init-unseal
slu vault init-unseal --file vault_init.local.json --namespace vault`,
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		vaultInitUnseal(FlagFile, FlagNamespace, FlagSelector, FlagKeyShares, FlagKeyThreshold)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"vault_init.local.json",
		"Path to save vault init JSON",
	)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"vault",
		"Namespace where Vault is running",
	)
	Cmd.Flags().StringVarP(
		&FlagSelector,
		"selector",
		"l",
		"app.kubernetes.io/name=vault,component=server",
		"Label selector of Vault server pods",
	)
	Cmd.Flags().IntVarP(
		&FlagKeyShares,
		"key-shares",
		"s",
		5,
		"Number of key shares to generate during Vault initialization",
	)
	Cmd.Flags().IntVarP(
		&FlagKeyThreshold,
		"key-threshold",
		"t",
		3,
		"Number of key shares required to unseal Vault",
	)
}

func vaultInitUnseal(file, namespace, selector string, keyShares, keyThreshold int) {
	if keyShares < 1 || keyThreshold < 1 {
		log.Fatalf("Key shares and threshold must be greater than 0")
	}
	if keyThreshold > keyShares {
		log.Fatalf("Key threshold cannot be greater than key shares")
	}
	if _, err := os.Stat(file); err == nil {
		log.Fatalf("File %s already exists, refusing to overwrite it", file)
	}

	podNames := getPods(namespace, selector)
	if len(podNames) == 0 {
		log.Fatalf("No Vault pods found in namespace %s (selector %s)", namespace, selector)
	}

	vaultKeys := vaultInit(podNames[0], file, namespace, keyShares, keyThreshold)
	unsealKeys, threshold := extractVaultKeys(vaultKeys)

	for _, podName := range podNames {
		log.Printf("Unsealing pod %s", podName)
		unsealPod(podName, namespace, unsealKeys, threshold)
	}

	log.Printf("Vault initialization and unsealing completed successfully. Keys saved to %s\n", file)
}

func getPods(namespace, selector string) []string {
	cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-l", selector, "-o", "jsonpath={.items[*].metadata.name}")
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing kubectl command: %v output %s", err, string(output))
	}
	return strings.Fields(string(output))
}

func vaultInit(pod, file, namespace string, keyShares, keyThreshold int) []byte {
	log.Printf("Executing vault init on pod %s in namespace %s\n", pod, namespace)
	cmd := exec.Command(
		"kubectl", "exec", pod, "-n", namespace, "--",
		"vault", "operator", "init", "-format=json",
		"-key-shares", fmt.Sprintf("%d", keyShares),
		"-key-threshold", fmt.Sprintf("%d", keyThreshold),
	)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing vault init command: %v\n", err)
	}

	err = os.WriteFile(file, output, 0600)
	if err != nil {
		log.Printf("[WARNING] Error writing output to file: %v", err)
		log.Printf("Output: %s\n", string(output))
	}

	return output
}

func extractVaultKeys(jsonData []byte) ([]string, int) {
	var response struct {
		UnsealKeysB64   []string `json:"unseal_keys_b64"`
		UnsealThreshold int      `json:"unseal_threshold"`
	}

	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return response.UnsealKeysB64, response.UnsealThreshold
}

func unsealPod(podName, namespace string, vaultKeys []string, threshold int) {
	for i, key := range vaultKeys {
		if i >= threshold {
			break
		}
		cmd := exec.Command("kubectl", "exec", podName, "-n", namespace, "--", "vault", "operator", "unseal", key)
		cmd.Env = os.Environ()
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Error unsealing pod %s: %v\nOutput: %s", podName, err, output)
		}
	}
	waitForPodReady(podName, namespace)
}

func waitForPodReady(pod, namespace string) {
	cmd := exec.Command("kubectl", "wait", "pod", pod, "-n", namespace, "--for=condition=Ready", "--timeout=60s")
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error waiting for pod to be ready: %v\nOutput: %s", err, output)
	}
}
