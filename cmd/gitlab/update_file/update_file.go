package update_file

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	parent_cmd "github.com/sikalabs/slu/cmd/gitlab"
	"github.com/spf13/cobra"
)

var FlagGitlabUrl string
var FlagToken string
var FlagProjectId int
var FlagBranch string
var FlagFile string
var FlagContent string
var FlagCommitterEmail string
var FlagCommitterName string
var FlagCommitMessage string

var Cmd = &cobra.Command{
	Use:   "update-file",
	Short: "Update file in GitLab using API",
	Run: func(cmd *cobra.Command, args []string) {
		gitlabUpdateFile(FlagGitlabUrl, FlagToken, FlagProjectId, FlagBranch, FlagFile, FlagContent, FlagCommitterEmail, FlagCommitterName, FlagCommitMessage)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagGitlabUrl,
		"gitlab-url",
		"u",
		"",
		"Gitlab URL",
	)
	Cmd.MarkFlagRequired("gitlab-url")
	Cmd.Flags().StringVarP(
		&FlagToken,
		"token",
		"t",
		"",
		"GitLab Token",
	)
	Cmd.MarkFlagRequired("token")
	Cmd.Flags().IntVarP(
		&FlagProjectId,
		"project-id",
		"p",
		0,
		"Project ID",
	)
	Cmd.MarkFlagRequired("project-id")
	Cmd.Flags().StringVarP(
		&FlagBranch,
		"branch",
		"b",
		"",
		"Branch",
	)
	Cmd.MarkFlagRequired("project-id")
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"",
		"File",
	)
	Cmd.MarkFlagRequired("file")
	Cmd.Flags().StringVarP(
		&FlagContent,
		"content",
		"c",
		"",
		"Content",
	)
	Cmd.MarkFlagRequired("content")
	Cmd.Flags().StringVarP(
		&FlagCommitterEmail,
		"committer-email",
		"e",
		"",
		"Committer Email",
	)
	Cmd.MarkFlagRequired("committer-email")
	Cmd.Flags().StringVarP(
		&FlagCommitterName,
		"committer-name",
		"n",
		"",
		"Committer Name",
	)
	Cmd.MarkFlagRequired("committer-name")
	Cmd.Flags().StringVarP(
		&FlagCommitMessage,
		"commit-message",
		"m",
		"",
		"Commit Message",
	)
	Cmd.MarkFlagRequired("commit-message")
}

func gitlabUpdateFile(gitlabUrl, token string, projectId int, branch, file, content, email, name, message string) {
	contentCurrent := readGitlabFile(gitlabUrl, token, projectId, branch, file)
	if contentCurrent == content {
		log.Println("Content is the same, skipping")
		return
	}
	rawGitlabUpdateFile(gitlabUrl, token, projectId, branch, file, content, email, name, message)
}

func rawGitlabUpdateFile(gitlabUrl, token string, projectId int, branch, file, content, email, name, message string) {
	url := fmt.Sprintf("%s/api/v4/projects/%d/repository/files/%s", gitlabUrl, projectId, url.QueryEscape(file))
	jsonData := `{
		"branch": "` + branch + `",
		"author_email": "` + email + `",
		"author_name": "` + name + `",
		"content": "` + content + `",
		"commit_message": "` + message + `"
	}`

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Fatalln("Error creating request:", err)
	}

	req.Header.Set("PRIVATE-TOKEN", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error making request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln("Error updating file:", resp.Status)
	}
}

func readGitlabFile(gitlabUrl, token string, projectId int, branch, file string) string {
	url := fmt.Sprintf("%s/api/v4/projects/%d/repository/files/%s?ref=%s", gitlabUrl, projectId, url.QueryEscape(file), branch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("Error creating request:", err)
	}

	req.Header.Set("PRIVATE-TOKEN", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error making request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln("Error updating file:", resp.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln("Error decoding response:", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(data["content"].(string))
	if err != nil {
		log.Fatal("Error decoding string:", err)
	}

	return string(decoded)
}
