package upload

import (
	"os"

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
		access_key := os.Getenv("SLU_UPLOAD_ACCESS_KEY")
		if access_key == "" {
			log.Fatalln("SLU_UPLOAD_ACCESS_KEY is empty")
		}
		secret_key := os.Getenv("SLU_UPLOAD_SECRET_KEY")
		if secret_key == "" {
			log.Fatalln("SLU_UPLOAD_SECRET_KEY is empty")
		}
		region := os.Getenv("SLU_UPLOAD_REGION")
		endpoint := os.Getenv("SLU_UPLOAD_ENDPOINT")
		if region == "" && endpoint == "" {
			log.Fatalln("SLU_UPLOAD_REGION and SLU_UPLOAD_ENDPOINT are empty")
		}
		bucket_name := os.Getenv("SLU_UPLOAD_BUCKET_NAME")
		if bucket_name == "" {
			log.Fatalln("SLU_UPLOAD_BUCKET_NAME is empty")
		}
		s3_utils.Upload(
			access_key,
			secret_key,
			region,
			endpoint,
			bucket_name,
			time_utils.NowForFileName()+"_"+filePath,
			f,
		)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
