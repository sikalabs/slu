package google_drive_utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sikalabs/slu/utils/vault_google_drive_utils"
	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func Upload(clientId, clientSecret, accessToken, fileToUpload string) {
	if clientId == "" && clientSecret == "" {
		clientId, clientSecret = GetGoogleDriveUploadSecretsFromVaultOrEnvOrDie()
	}

	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{drive.DriveFileScope},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	ctx := context.Background()

	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	client := conf.Client(ctx, token)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	f, err := os.Open(fileToUpload)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer f.Close()

	fileMetadata := &drive.File{
		Name: filepath.Base(fileToUpload),
	}
	file, err := srv.Files.Create(fileMetadata).Media(f).Do()
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	fmt.Printf("File ID: %s\n", file.Id)
}

func GetToken(clientId, clientSecret string) {
	if clientId == "" && clientSecret == "" {
		clientId, clientSecret = GetGoogleDriveUploadSecretsFromVaultOrEnvOrDie()
	}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/drive.file"},
		Endpoint: oauth2.Endpoint{
			TokenURL:      "https://oauth2.googleapis.com/token",
			DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		},
	}

	code, err := conf.DeviceAuth(ctx)
	if err != nil {
		log.Fatalf("Failed to get device and user codes: %v", err)
	}

	fmt.Printf("Visit the URL: %s\n", code.VerificationURI)
	fmt.Printf("And enter the code: %s\n", code.UserCode)

	for {
		token, err := conf.DeviceAccessToken(ctx, code)
		if err == nil {
			fmt.Printf("Got access token: %s\n", token.AccessToken)
			fmt.Printf("Got refresh token: %s\n", token.RefreshToken)
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func GetGoogleDriveUploadSecretsFromVaultOrEnvOrDie() (
	string, string,
) {
	clientIdVault, clientSecretVault,
		_ := vault_google_drive_utils.GetGoogleDriveUploadSecrets("secret/data/slu/google-drive-upload/client")

	// Client ID
	var clientId string
	clientIdEnv := os.Getenv("SLU_GOOGLE_DRIVE_UPLOAD_CLIENT_ID")
	if clientIdVault != "" {
		clientId = clientIdVault
	}
	if clientIdEnv != "" {
		clientId = clientIdEnv
	}
	if clientId == "" {
		log.Fatalln("SLU_GOOGLE_DRIVE_UPLOAD_CLIENT_ID is empty")
	}

	// Client Secret
	var clientSecret string
	clientSecretEnv := os.Getenv("SLU_GOOGLE_DRIVE_UPLOAD_CLIENT_SECRET")
	if clientIdVault != "" {
		clientSecret = clientSecretVault
	}
	if clientIdEnv != "" {
		clientSecret = clientSecretEnv
	}
	if clientSecret == "" {
		log.Fatalln("SLU_GOOGLE_DRIVE_UPLOAD_CLIENT_SECRET is empty")
	}

	return clientId, clientSecret
}
