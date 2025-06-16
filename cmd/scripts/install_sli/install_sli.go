package install_sli

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"runtime"
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
		encryptedToken              = "eqLoHqyn0Sq9t+SRafoI/XJKA4ePdUn4DylKVfn2tQMvG06WxRyAN6Bdj90aCQ5VX4/E6t+jHOB1awwnfDX1SNISAeYk7bAL5+2whjL++kUxQMAHKELZt+hegcRYOGFQcpDLnfXrXlELu+Ox4CK1ttUXjAi9f+Xdew=="
		encryptedAssetIDDarwinArm64 = "ZeX1YA38SMaPx/JlCpwBYxwWeRYPcyg5yXU28hhvIhttqLJopQ=="
		encryptedAssetIDLinuxAmd64  = "AgSK+LhuapiuN8hVJGZbCyAeYq4SVuWNS3/IwQQWdgQaxlKbxQ=="
	)

	assetId := ""

	if flagOs == "darwin" {
		assetId = decrypt(encryptedAssetIDDarwinArm64, password)
	} else if flagOs == "linux" {
		assetId = decrypt(encryptedAssetIDLinuxAmd64, password)
	} else {
		log.Fatalf("Unsupported OS: %s. Supported OS are: darwin, linux.", flagOs)
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
