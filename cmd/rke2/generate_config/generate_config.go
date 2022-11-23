package generate_config

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/rke2"
	"github.com/sikalabs/slu/utils/rke2_utils"

	"github.com/spf13/cobra"
)

var FlagFile string
var FlagDomain string
var FlagToken string
var FlagTlsSans []string
var FlagIsFirstMaster bool

var Cmd = &cobra.Command{
	Use:     "generate-config",
	Short:   "Generate RKE2 config file",
	Aliases: []string{"genconf", "gen"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(rke2_utils.GenerateMasterConfig(
			FlagDomain,
			FlagToken,
			FlagTlsSans,
			FlagIsFirstMaster,
		))
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"",
		"YAML file",
	)
	Cmd.Flags().StringVarP(
		&FlagDomain,
		"domain",
		"d",
		"",
		"Server domain",
	)
	Cmd.MarkFlagRequired("domain")
	Cmd.Flags().StringVarP(
		&FlagToken,
		"token",
		"t",
		"",
		"Token",
	)
	Cmd.MarkFlagRequired("token")
	Cmd.Flags().StringArrayVarP(
		&FlagTlsSans,
		"tls-san",
		"s",
		[]string{},
		"TLS SAN",
	)
	Cmd.Flags().BoolVar(
		&FlagIsFirstMaster,
		"is-first-master",
		false,
		"Is first master",
	)
}
