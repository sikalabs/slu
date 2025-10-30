package gitignore

import "os"

var GitignoreBase = `# Mac
.DS_Store

# Editor
.vscode
.idea
*.swp
*.swo
*~

# Generic
*.log
*.backup
*.local.sh
*.local.yml
*.local.yaml
*.local.json
*.local.txt
Makefile.local.mk
`

var GitignoreTerraform = `# Terraform
.terraform
*.tfplan
*.tfstate
.terraform.tfstate.lock.info
*.tfvars
!*.EXAMPLE.tfvars
override.tf

# Terragrunt
.terragrunt-cache

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

var GitignoreKubernetes = `# Kuberneres
kubeconfig.yml
kubeconfig.*.yml
`

var GitignoreHelm = `# Helm
*.tgz
`

func CreateGitignore(content string) {
	err := os.WriteFile(".gitignore", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
