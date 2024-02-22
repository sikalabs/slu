package download

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/3rdparty/nginx_log_generator"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "nginx-log-generator",
	Aliases: []string{"nlg"},
	Short:   "Nginx log generator (from kscarlett/nginx-log-generator)",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		nginx_log_generator.NginxLogGenerator()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
