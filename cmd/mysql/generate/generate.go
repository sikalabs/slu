package drop

import (
	mysqlcmd "github.com/sikalabs/slu/cmd/mysql"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate data in MySQL",
	Aliases: []string{"g", "gen"},
}

func init() {
	mysqlcmd.MysqlCmd.AddCommand(Cmd)
}
