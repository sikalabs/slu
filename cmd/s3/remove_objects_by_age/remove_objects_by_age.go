package remove_objects_by_age

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/s3"
	"github.com/sikalabs/slu/utils/3rdparty/parseduration"
	"github.com/sikalabs/slu/utils/s3_utils"

	"github.com/spf13/cobra"
)

var FlagAccessKey string
var FlagSecretKey string
var FlagRegion string
var FlagEndpoint string
var FlagBucketName string
var FlagAge string

var Cmd = &cobra.Command{
	Use:     "remove-objects-by-age",
	Short:   "Remove S3 objects by age",
	Aliases: []string{"rb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		age, err := parseduration.ParseDuration(FlagAge)
		if err != nil {
			log.Fatalln(err)
		}
		s3_utils.RemoveObjectsByAge(
			FlagAccessKey, FlagSecretKey, FlagRegion,
			FlagEndpoint, FlagBucketName, age)
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
		&FlagAge,
		"age",
		"",
		"Age of objects to delete",
	)
	Cmd.MarkFlagRequired("age")
}
