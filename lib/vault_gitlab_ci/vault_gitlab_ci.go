package vault_gitlab_ci

import (
	"fmt"
	"strconv"

	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
)

func GetGitlabCiSecrets(name string) (string, string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read("secret/data/slu/gitlab-ci/" + name)
	if err != nil {
		return "", "", err
	}
	if secret == nil {
		return "", "", fmt.Errorf("secret secret/data/slu/gitlab-ci/" + name + " not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("wrong data")
	}
	gitlabUrl, err := getString(data, "GITLAB_URL", true)
	if err != nil {
		return "", "", err
	}
	registrationToken, err := getString(data, "TOKEN", true)
	if err != nil {
		return "", "", err
	}
	return gitlabUrl, registrationToken, nil
}

func getString(data map[string]interface{}, key string, required bool) (string, error) {
	val, ok := data[key]
	if !ok {
		if !required {
			return "", nil
		}
		return "", fmt.Errorf("key \"%s\" not found", key)
	}
	if val == nil {
		if !required {
			return "", nil
		}
		return "", fmt.Errorf("no value for key \"%s\"", key)
	}
	return val.(string), nil
}

func getIntFromString(data map[string]interface{}, key string, required bool) (int, error) {
	val, ok := data[key]
	if !ok {
		if !required {
			return 0, nil
		}
		return 0, fmt.Errorf("key \"%s\" not found", key)
	}
	if val == nil {
		if !required {
			return 0, nil
		}
		return 0, fmt.Errorf("no value for key \"%s\"", key)
	}
	return strconv.Atoi(val.(string))
}
