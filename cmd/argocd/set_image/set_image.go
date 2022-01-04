package set_image

import (
	"io/ioutil"
	"regexp"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"

	"github.com/spf13/cobra"
)

var FlagFile string
var FlagKey string
var FlagValue string

func replaceImage(s, key, value string) string {
	r := regexp.MustCompile(key + `: +([\w\./:_-]+)`)
	return r.ReplaceAllString(s, key+": "+value)
}

var Cmd = &cobra.Command{
	Use:     "set-image",
	Short:   "Set image in ArgoCD YAML file",
	Aliases: []string{"si"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		data, _ := ioutil.ReadFile(FlagFile)
		s := replaceImage(string(data), FlagKey, FlagValue)
		ioutil.WriteFile(FlagFile, []byte(s), 0644)
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"",
		"YAML file",
	)
	Cmd.MarkFlagRequired("file")
	Cmd.Flags().StringVarP(
		&FlagKey,
		"key",
		"k",
		"",
		"image key in Helm values",
	)
	Cmd.MarkFlagRequired("key")
	Cmd.Flags().StringVarP(
		&FlagValue,
		"value",
		"v",
		"",
		"New image",
	)
	Cmd.MarkFlagRequired("value")
}
