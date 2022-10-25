package copy_from_cloud

import (
	"log"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/s3_utils"
	"github.com/sikalabs/slu/utils/vault_s3_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "copy-from-cloud <cloud alias> <local target file>",
	Short:   "Copy local file from temporary cloud storage",
	Aliases: []string{"cfc"},
	Args:    cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		sourceAlias := args[0]
		targetFilePath := args[1]

		accessKey, secretKey, region,
			endpoint, bucketName, _ := vault_s3_utils.GetS3Secrets("secret/data/slu/cloud-copy")

		err := s3_utils.DownloadToFile(
			accessKey,
			secretKey,
			region,
			endpoint,
			bucketName,
			sourceAlias,
			targetFilePath,
		)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
