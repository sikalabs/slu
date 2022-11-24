package copy_to_s3

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/s3"
	"github.com/sikalabs/slu/utils/s3_utils"

	"github.com/spf13/cobra"
)

var FlagAccessKey string
var FlagSecretKey string
var FlagRegion string
var FlagEndpoint string
var FlagBucketName string
var FlagTarget string
var FlagSource string

var Cmd = &cobra.Command{
	Use:   "copy-from-s3",
	Short: "Copy file from s3 to local",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := s3_utils.DownloadToFile(
			FlagAccessKey,
			FlagSecretKey,
			FlagRegion,
			FlagEndpoint,
			FlagBucketName,
			FlagSource,
			FlagTarget,
		)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagAccessKey,
		"access-key",
		"a",
		"",
		"AWS Access Key",
	)
	Cmd.MarkFlagRequired("access-key")
	Cmd.Flags().StringVarP(
		&FlagSecretKey,
		"secret-key",
		"s",
		"",
		"AWS Secret Key",
	)
	Cmd.MarkFlagRequired("secret-key")
	Cmd.Flags().StringVarP(
		&FlagRegion,
		"region",
		"r",
		"",
		"AWS Region",
	)
	Cmd.Flags().StringVarP(
		&FlagEndpoint,
		"endpoint",
		"e",
		"",
		"Custom S3 Endpoint",
	)
	Cmd.Flags().StringVarP(
		&FlagBucketName,
		"bucket-name",
		"b",
		"",
		"Bucket Name",
	)
	Cmd.MarkFlagRequired("bucket-name")
	Cmd.Flags().StringVar(
		&FlagTarget,
		"target",
		"",
		"Local target path",
	)
	Cmd.MarkFlagRequired("target")
	Cmd.Flags().StringVar(
		&FlagSource,
		"source",
		"",
		"Source s3 path",
	)
	Cmd.MarkFlagRequired("source")
}
