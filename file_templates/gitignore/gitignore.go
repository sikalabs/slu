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
*.local.sh
*.local.yml
*.local.yaml
*.local.json
*.local.txt
`

var GitignoreTerraform = `# Terraform
.terraform
*.tfstate
.terraform.tfstate.lock.info
*.tfvars
!*.EXAMPLE.tfvars
override.tf

# Infracost
.infracost
.infracost-reports
`

var GitignoreNodeJS = `# NodeJS
node_modules
`

var GitignoreNextJs = `# NodeJS
.next
out
`

func CreateGitignore(content string) {
	err := ioutil.WriteFile(".gitignore", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
