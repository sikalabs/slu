package remove_bucket

import (
	parent_cmd "github.com/sikalabs/slu/cmd/s3"
	"github.com/sikalabs/slu/utils/s3_utils"

	"github.com/spf13/cobra"
)

var FlagAccessKey string
var FlagSecretKey string
var FlagRegion string
var FlagEndpoint string
var FlagBucketName string

var Cmd = &cobra.Command{
	Use:     "remove-bucket",
	Short:   "Remove S3 Bucket with objects in it",
	Aliases: []string{"rb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		s3_utils.DeleteBucketWithObjects(
			FlagAccessKey, FlagSecretKey, FlagRegion,
			FlagEndpoint, FlagBucketName)
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
}
