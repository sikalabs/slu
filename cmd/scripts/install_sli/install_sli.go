package install_sli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"syscall"

	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
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
		encryptedToken = "UGRvP1IN2AODY9ArML9cly62m4RCrKWXB01+kOTDBOG68bXaQJvfgHTaNvqevT0qGDYR6b+kJsZemfbS2phT2c0Vc0M/BdbrTa44lh6hPTfz0PuQcVfVAgEk2O/uwaIHUTm5A7Z4p9JmD9aR4GQoWF+2SgzBMxkfpoDUR8Z0iszC6857lJ+2VYw="
	)

	token, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(password, encryptedToken)
	if err != nil {
		log.Fatalf("Failed to decrypt token: %v", err)
	}

	assetId, err := getAssetID(FlagOS, token)
	if err != nil {
		log.Fatalf("Failed to get asset ID: %v", err)
	}

	tar_gz_utils.WebTarGzToBin(
		"https://api.github.com/repos/sikalabs/sli/releases/assets/"+assetId,
		"sli",
		map[string]string{
			"Accept":        "application/octet-stream",
			"Authorization": "Bearer " + token,
		},
		"./sli",
	)
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
