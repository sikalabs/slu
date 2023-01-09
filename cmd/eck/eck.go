package argocd

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "eck",
	Short: "ECK (Elastic Cloud on Kubernetes) Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
