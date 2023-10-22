package google_drive_utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func Upload(clientId, clientSecret, accessToken, fileToUpload string) {
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
