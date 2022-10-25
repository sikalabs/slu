package copy_to_cloud

import (
	"os"

	"log"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/s3_utils"
	"github.com/sikalabs/slu/utils/vault_s3_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "copy-to-cloud <local source file> <cloud alias>",
	Short:   "Copy local file to temporary cloud storage",
	Aliases: []string{"ctc"},
	Args:    cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		sourceFile := args[0]
		targetAlias := args[1]

		f, err := os.Open(sourceFile)
		if err != nil {
			log.Fatalln(err)
		}

		accessKey, secretKey, region,
			endpoint, bucketName, err := vault_s3_utils.GetS3Secrets("secret/data/slu/cloud-copy")
		if err != nil {
			log.Fatalln(err)
		}

		err = s3_utils.Upload(
			accessKey,
			secretKey,
			region,
			endpoint,
			bucketName,
			targetAlias,
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
