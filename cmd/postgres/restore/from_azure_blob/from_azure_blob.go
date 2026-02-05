package from_azure_blob

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

		reader := bufio.NewReader(os.Stdin)

		fmt.Print(fmt.Sprintf("Do you really want to restore Postgres database `%s` from Azure Blob Storage?\nThis will overwrite existing database data!\nType 'yes' to confirm: ", postgresrestorecmd.PostgresRestoreFlagName))

		response, _ := reader.ReadString('\n')

		if strings.TrimSpace(response) == "yes" {
			fmt.Print(fmt.Sprintf("Postgres database `%s` will be restored from Azure Blob Storage. ", postgresrestorecmd.PostgresRestoreFlagName))

			for i := 10; i > 0; i-- {
				fmt.Printf("\rWait for %d seconds... cancel using ctrl+c ", i)
				time.Sleep(1 * time.Second)
			}

			fmt.Println("\nStart!")

			err := postgres.RestoreFromAzureBlob(
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
				postgresrestorecmd.PostgresRestoreFlagSSLMode,
			)

			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Printf("Postgres database `%s` restore from Azure Blob Storage cancelled\n", postgresrestorecmd.PostgresRestoreFlagName)
		}
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
