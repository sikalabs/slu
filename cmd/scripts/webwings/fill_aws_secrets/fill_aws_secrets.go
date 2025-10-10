package get_images_from_env_for_values_yaml

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/webwings"
	"github.com/spf13/cobra"
)

var FlagRegion string

var Cmd = &cobra.Command{
	Use:   "fill-aws-secrets",
	Short: "Fill AWS Secrets",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fill_aws_secrets()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagRegion,
		"region",
		"r",
		"eu-central-1",
		"AWS Region",
	)
}

func fill_aws_secrets() {
	ctx := context.Background()

	// Initialize AWS config
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))
	if err != nil {
		panic(err)
	}
	smClient := secretsmanager.NewFromConfig(cfg)

	// List all secrets
	input := &secretsmanager.ListSecretsInput{
		MaxResults: aws.Int32(100),
	}
	result, err := smClient.ListSecrets(ctx, input)
	if err != nil {
		panic(err)
	}

	// Iterate over all secrets and check if they are empty
	for _, secret := range result.SecretList {
		secretValue, _ := smClient.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
			SecretId: secret.Name,
		})

		if secretValue.SecretString == nil || *secretValue.SecretString == "" || strings.TrimSpace(*secretValue.SecretString) == "" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Enter value for secret %s: ", *secret.Name)
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)

			_, err := smClient.PutSecretValue(ctx, &secretsmanager.PutSecretValueInput{
				SecretId:     secret.Name,
				SecretString: aws.String(value),
			})
			if err != nil {
				fmt.Println("Error updating secret value for", *secret.Name)
			} else {
				fmt.Println("Updated secret value for", *secret.Name)
			}
		}
	}

	fmt.Println("Finished updating secrets.")
}
