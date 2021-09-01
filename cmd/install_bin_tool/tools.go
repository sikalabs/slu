package install_bin_tool

var Tools = []Tool{
	{
		Name:        "tergum",
		Version:     "v0.12.0",
		UrlTemplate: "https://github.com/sikalabs/tergum/releases/download/{{.Version}}/tergum_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:        "statica",
		Version:     "v0.4.0",
		UrlTemplate: "https://github.com/vojtechmares/statica/releases/download/{{.Version}}/statica_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name: "kubectl",
		// Get Version: `curl -L -s https://dl.k8s.io/release/stable.txt`
		Version:     "v1.22.1",
		UrlTemplate: "https://dl.k8s.io/release/{{.Version}}/bin/{{.Os}}/{{.Arch}}/kubectl",
	},
	{
		Name:        "helm",
		Version:     "v3.6.3",
		SourcePath:  "{{.Os}}-amd64/helm",
		UrlTemplate: "https://get.helm.sh/helm-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:        "docker",
		SourcePath:  "docker/docker",
		Version:     "20.10.8",
		UrlTemplate: "https://download.docker.com/{{.OsDocker}}/static/stable/{{.ArchDocker}}/docker-{{.Version}}.tgz",
	},
	{
		Name:        "docker-compose",
		Version:     "1.29.2",
		UrlTemplate: "https://github.com/docker/compose/releases/download/{{.Version}}/docker-compose-{{.Os|capitalize}}-{{.ArchDocker}}",
	},
	{
		Name:        "mcli",
		UrlTemplate: "https://dl.minio.io/client/mc/release/{{.Os}}-{{.Arch}}/mc",
	},
	{
		Name:        "terraform",
		Version:     "1.0.5",
		UrlTemplate: hashicorpUrlTemplate("terraform"),
	},
	{
		Name:        "vault",
		Version:     "1.8.2",
		UrlTemplate: hashicorpUrlTemplate("vault"),
	},
	{
		Name:        "doctl",
		Version:     "1.64.0",
		UrlTemplate: "https://github.com/digitalocean/doctl/releases/download/v{{.Version}}/doctl-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:        "skaffold",
		Version:     "latest",
		UrlTemplate: "https://storage.googleapis.com/skaffold/releases/{{.Version}}/skaffold-{{.Os}}-{{.Arch}}",
	},
	{
		Name:        "glab",
		Version:     "1.20.0",
		UrlTemplate: "https://github.com/profclems/glab/releases/download/v{{.Version}}/glab_{{.Version}}_{{.Os|capitalize}}_{{.ArchDocker}}.tar.gz",
		SourcePath:  "bin/glab",
	},
}

func hashicorpUrlTemplate(name string) string {
	return "https://releases.hashicorp.com/" + name + "/{{.Version}}/" + name +
		"_{{.Version}}_{{.Os}}_{{.Arch}}.zip"
}
