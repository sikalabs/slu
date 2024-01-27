package skip_stage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

func SkipStageShow(path string) {
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
	variables, _, err := g.ProjectVariables.ListVariables(
		config.ProjectID,
		&gitlab.ListProjectVariablesOptions{},
	)
	if err != nil {
		panic(err)
	}
	for _, v := range variables {
		if !strings.HasPrefix(v.Key, "SKIP_STAGE_") {
			continue
		}
		fmt.Printf("KEY:   %s\nVALUE: %s\n---\n", v.Key, v.Value)
	}
}

func SkipStage(path string, stage string) {
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

	key := "SKIP_STAGE_" + stage
	val := "skip"

	v, _, _ := g.ProjectVariables.GetVariable(config.ProjectID, key, &gitlab.GetProjectVariableOptions{})
	if v == nil {
		g.ProjectVariables.CreateVariable(config.ProjectID, &gitlab.CreateProjectVariableOptions{
			Key:   &key,
			Value: &val,
		})
	} else {
		g.ProjectVariables.UpdateVariable(config.ProjectID, key, &gitlab.UpdateProjectVariableOptions{
			Value: &val,
		})
	}
}

func RemoveSkipStage(path string, stage string) {
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

	key := "SKIP_STAGE_" + stage
	val := "no-skip"

	v, _, _ := g.ProjectVariables.GetVariable(config.ProjectID, key, &gitlab.GetProjectVariableOptions{})
	if v == nil {
		g.ProjectVariables.CreateVariable(config.ProjectID, &gitlab.CreateProjectVariableOptions{
			Key:   &key,
			Value: &val,
		})
	} else {
		g.ProjectVariables.UpdateVariable(config.ProjectID, key, &gitlab.UpdateProjectVariableOptions{
			Value: &val,
		})
	}
}

func RemoveVariableSkipStage(path string, stage string) {
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

	key := "SKIP_STAGE_" + stage

	v, _, _ := g.ProjectVariables.GetVariable(config.ProjectID, key, &gitlab.GetProjectVariableOptions{})
	if v != nil {
		g.ProjectVariables.RemoveVariable(config.ProjectID, key, &gitlab.RemoveProjectVariableOptions{})
	}
}
