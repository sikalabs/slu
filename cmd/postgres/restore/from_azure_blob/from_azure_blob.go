package from_azure_blob

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	postgresrestorecmd "github.com/sikalabs/slu/cmd/postgres/restore"
	"github.com/sikalabs/slu/utils/postgres"

	"github.com/spf13/cobra"
)

var CmdFlagContainerName string
var CmdFlagAccountName string
var CmdFlagAccountKey string
var CmdFlagSourcePrefix string
var CmdFlagSourceSuffix string

var Cmd = &cobra.Command{
	Use:   "from-azure-blob",
	Short: "Restore Postgres database from Azure Blob Storage",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		postgres.RestoreFromAzureBlob(
			postgrescmd.PostgresCmdFlagHost,
			postgrescmd.PostgresCmdFlagPort,
			postgrescmd.PostgresCmdFlagUser,
			postgrescmd.PostgresCmdFlagPassword,
			postgresrestorecmd.PostgresRestoreFlagName,
			CmdFlagContainerName,
			CmdFlagAccountName,
			CmdFlagAccountKey,
			CmdFlagSourcePrefix,
			CmdFlagSourceSuffix,
		)
	},
}

func init() {
	postgresrestorecmd.PostgresRestoreCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagContainerName,
		"container-name",
		"c",
		"",
		"Azure Blob Storage container name",
	)
	Cmd.MarkFlagRequired("container-name")

	Cmd.Flags().StringVarP(
		&CmdFlagAccountName,
		"account-name",
		"a",
		"",
		"Azure Blob Storage account name",
	)
	Cmd.MarkFlagRequired("account-name")

	Cmd.Flags().StringVarP(
		&CmdFlagAccountKey,
		"account-key",
		"k",
		"",
		"Azure Blob Storage account key",
	)
	Cmd.MarkFlagRequired("account-key")

	Cmd.Flags().StringVarP(
		&CmdFlagSourcePrefix,
		"source-prefix",
		"r",
		"",
		"Source backup file prefix in Azure Blob Storage. e.g. prod_db",
	)
	Cmd.MarkFlagRequired("source-prefix")

	Cmd.Flags().StringVarP(
		&CmdFlagSourceSuffix,
		"source-suffix",
		"s",
		"",
		"Source backup file suffix in Azure Blob Storage. e.g. full.gz",
	)
	Cmd.MarkFlagRequired("source-suffix")
}
