package upload

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/lib/vault_upload"
	"github.com/sikalabs/slu/utils/s3_utils"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "upload <file>",
	Short: "Upload file to S3",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		filePath := args[0]
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatalln(err)
		}

		accessKeyVault, secretKeyVault, regionVault,
			endpointVault, bucketNameVault, _ := vault_upload.GetUploadSecrets()

		// Access Key
		var accessKey string
		accessKeyEnv := os.Getenv("SLU_UPLOAD_ACCESS_KEY")
		if accessKeyVault != "" {
			accessKey = accessKeyVault
		}
		if accessKeyEnv != "" {
			accessKey = accessKeyEnv
		}
		if accessKey == "" {
			log.Fatalln("SLU_UPLOAD_ACCESS_KEY is empty")
		}

		// Secret Key
		var secretKey string
		secretKeyEnv := os.Getenv("SLU_UPLOAD_ACCESS_KEY")
		if secretKeyVault != "" {
			secretKey = secretKeyVault
		}
		if accessKeyEnv != "" {
			secretKey = secretKeyEnv
		}
		if accessKey == "" {
			log.Fatalln("SLU_UPLOAD_SECRET_KEY is empty")
		}

		// Region
		var region string
		regionEnv := os.Getenv("SLU_UPLOAD_REGION")
		if regionVault != "" {
			region = regionVault
		}
		if regionEnv != "" {
			region = regionEnv
		}

		// Endpoint
		var endpoint string
		endpointEnv := os.Getenv("SLU_UPLOAD_ENDPOINT")
		if endpointVault != "" {
			endpoint = endpointVault
		}
		if endpointEnv != "" {
			endpoint = endpointEnv
		}

		// Region, Endpoint Validation
		if region == "" && endpoint == "" {
			log.Fatalln("SLU_UPLOAD_REGION and SLU_UPLOAD_ENDPOINT are empty")
		}

		// Secret Key
		var bucketName string
		bucketNameEnv := os.Getenv("SLU_UPLOAD_BUCKET_NAME")
		if bucketNameVault != "" {
			bucketName = bucketNameVault
		}
		if bucketNameEnv != "" {
			bucketName = bucketNameEnv
		}
		if bucketName == "" {
			log.Fatalln("SLU_UPLOAD_BUCKET_NAME is empty")
		}

		key := time_utils.NowForFileName() + "_" + filePath

		err = s3_utils.Upload(
			accessKey,
			secretKey,
			region,
			endpoint,
			bucketName,
			key,
			f,
		)
		if err != nil {
			log.Fatalln(err)
		}
		url, err := s3_utils.GetObjectPresignUrl(
			accessKey,
			secretKey,
			region,
			endpoint,
			bucketName,
			key,
			24*time.Hour,
		)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Key:              %s\n", key)
		fmt.Printf("Download (1 day): %s\n", url)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
