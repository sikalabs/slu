package upload

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/sikalabs/slu/cmd/root"
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

		accessKey, secretKey, region,
			endpoint, bucketName := s3_utils.GetS3SecretsFromVaultOrEnvOrDie("secret/data/slu/upload")

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
