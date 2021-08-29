package install_bin_tool

var Tools = []Tool{
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
}

func hashicorpUrlTemplate(name string) string {
	return "https://releases.hashicorp.com/" + name + "/{{.Version}}/" + name +
		"_{{.Version}}_{{.Os}}_{{.Arch}}.zip"
}
