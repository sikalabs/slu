package gitignore

import "io/ioutil"

var GitignoreBase = `# Mac
.DS_Store

# Editor
.vscode
.idea

# Generic
*.log
*.backup
`

var GitignoreTerraform = `# Terraform
.terraform
*.tfstate
.terraform.tfstate.lock.info
*.tfvars
!*.EXAMPLE.tfvars
`

func CreateGitignore(content string) {
	err := ioutil.WriteFile(".gitignore", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
