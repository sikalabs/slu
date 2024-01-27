package prune_environments

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

type Config struct {
	ApiUrl    string
	Token     string
	ProjectID int
}

func LoadConfig(config *Config, path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func PruneEnvironments(path string) {
	var config Config
	err := LoadConfig(&config, path)
	if err != nil {
		panic(err)
	}
	g, err := gitlab.NewClient(
		config.Token,
		gitlab.WithBaseURL(config.ApiUrl),
	)
	if err != nil {
		panic(err)
	}
	for {
		envs, _, err := g.Environments.ListEnvironments(
			config.ProjectID,
			&gitlab.ListEnvironmentsOptions{},
		)
		if err != nil {
			panic(err)
		}
		if len(envs) == 0 {
			break
		}
		for _, e := range envs {
			fmt.Println("stop & delete environment:", e.Name)
			_, _, err = g.Environments.StopEnvironment(config.ProjectID, e.ID)
			if err != nil {
				panic(err)
			}
			_, err = g.Environments.DeleteEnvironment(config.ProjectID, e.ID)
			if err != nil {
				panic(err)
			}
		}
	}
}
