package install_sli

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"syscall"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/tar_gz_utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var FlagOS string

var Cmd = &cobra.Command{
	Use:     "install-sli",
	Short:   "install sli - sikalabs internal utils",
	Aliases: []string{"isli"},
	Run: func(c *cobra.Command, args []string) {
		fmt.Print("Enter password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println() // Print a newline after password input

		installSli(string(bytePassword), FlagOS)

		fmt.Println("sli installed successfully to current folder! You can try it with `./sli` command.")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&FlagOS,
		"os",
		"o",
		runtime.GOOS,
		"OS",
	)
}

func installSli(password, flagOs string) {
	const (
		encryptedToken = "eqLoHqyn0Sq9t+SRafoI/XJKA4ePdUn4DylKVfn2tQMvG06WxRyAN6Bdj90aCQ5VX4/E6t+jHOB1awwnfDX1SNISAeYk7bAL5+2whjL++kUxQMAHKELZt+hegcRYOGFQcpDLnfXrXlELu+Ox4CK1ttUXjAi9f+Xdew=="
	)

	assetId, err := getAssetID(FlagOS, decrypt(encryptedToken, password))
	if err != nil {
		log.Fatalf("Failed to get asset ID: %v", err)
	}

	tar_gz_utils.WebTarGzToBin(
		"https://api.github.com/repos/sikalabs/sli/releases/assets/"+assetId,
		"sli",
		map[string]string{
			"Accept":        "application/octet-stream",
			"Authorization": "Bearer " + decrypt(encryptedToken, password),
		},
		"./sli",
	)
}

func decrypt(encryptedDataBase64, password string) string {
	key := deriveKey(password)

	encryptedData, err := base64.StdEncoding.DecodeString(encryptedDataBase64)
	if err != nil {
		log.Fatalf("Failed to decode encrypted data: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Failed to create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		log.Fatal("Ciphertext too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	return string(plaintext)
}

func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func getAssetID(os, token string) (string, error) {
	release, err := getReleaseInfo(token)
	if err != nil {
		return "", fmt.Errorf("failed to get release info: %v", err)
	}

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, os) {
			return fmt.Sprintf("%d", asset.ID), nil
		}
	}

	return "", fmt.Errorf("asset not found")
}

type releaseInfo struct {
	Assets []releaseInfoAsset `json:"assets"`
}

type releaseInfoAsset struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func getReleaseInfo(token string) (releaseInfo, error) {
	url := "https://api.github.com/repos/sikalabs/sli/releases/latest"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return releaseInfo{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return releaseInfo{}, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return releaseInfo{}, err
	}

	var release releaseInfo
	if err := json.Unmarshal(body, &release); err != nil {
		return releaseInfo{}, fmt.Errorf("failed to parse response: %v", err)
	}

	return release, nil
}
