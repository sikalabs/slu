package copy_from_s3

import (
	"log"
	"os"

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
	Use:   "copy-to-s3",
	Short: "Copy local file to s3",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		f, err := os.Open(FlagSource)
		if err != nil {
			log.Fatalln(err)
		}

		err = s3_utils.Upload(
			FlagAccessKey,
			FlagSecretKey,
			FlagRegion,
			FlagEndpoint,
			FlagBucketName,
			FlagTarget,
			f,
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
		"Target s3 path",
	)
	Cmd.MarkFlagRequired("target")
	Cmd.Flags().StringVar(
		&FlagSource,
		"source",
		"",
		"Source file path",
	)
	Cmd.MarkFlagRequired("source")
}
