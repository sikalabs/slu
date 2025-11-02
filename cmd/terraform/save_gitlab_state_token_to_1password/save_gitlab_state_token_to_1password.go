package save_gitlab_state_token_to_1password

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	parentcmd "github.com/sikalabs/slu/cmd/terraform"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagGitlabDomain string

var Cmd = &cobra.Command{
	Use:   "save-gitlab-state-token-to-1password",
	Short: "Save GitLab Terraform state token to 1Password employee vault",
	Run: func(c *cobra.Command, args []string) {
		// Validate required flag
		if FlagGitlabDomain == "" {
			error_utils.HandleError(fmt.Errorf("--gitlab-domain is required"), "Missing required flag")
		}

		// Check if stdin has data
		fi, err := os.Stdin.Stat()
		if err != nil {
			error_utils.HandleError(err, "Failed to check stdin")
		}

		isPipe := fi.Mode()&os.ModeNamedPipe != 0
		if !isPipe {
			error_utils.HandleError(fmt.Errorf("no input from pipe"), "Token must be provided via stdin")
		}

		// Read token from stdin
		scanner := bufio.NewScanner(os.Stdin)
		var token string
		if scanner.Scan() {
			token = strings.TrimSpace(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			error_utils.HandleError(err, "Failed to read from stdin")
		}

		if token == "" {
			error_utils.HandleError(fmt.Errorf("empty token"), "Token cannot be empty")
		}

		// Construct the item name
		itemName := fmt.Sprintf("GITLAB_TOKEN_TF_STATE_%s", FlagGitlabDomain)

		// Save to 1Password using op CLI
		// Using 'op item create' with template 'Login' and storing token in password field
		fmt.Printf("Saving token to 1Password item: %s\n", itemName)

		err = exec_utils.ExecOut("op", "item", "create",
			"--category", "password",
			"--title", itemName,
			"--vault", "employee",
			fmt.Sprintf("password=%s", token),
		)

		if err != nil {
			error_utils.HandleError(err, "Failed to save token to 1Password")
		}

		fmt.Printf("Successfully saved token to 1Password item: %s in employee vault\n", itemName)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagGitlabDomain,
		"gitlab-domain",
		"g",
		"",
		"GitLab domain (required)",
	)
	Cmd.MarkFlagRequired("gitlab-domain")
}
