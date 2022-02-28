package install_bin_tool

import "github.com/sikalabs/slu/utils/github_utils"

var Tools = []Tool{
	{
		Name:           "install-slu",
		GetVersionFunc: func() string { return "v0.1.0" },
		UrlTemplate:    "https://github.com/sikalabs/install-slu/releases/download/{{.Version}}/install-slu_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "tergum",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "tergum") },
		UrlTemplate:    "https://github.com/sikalabs/tergum/releases/download/{{.Version}}/tergum_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "training-cli",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("ondrejsika", "training-cli") },
		UrlTemplate:    "https://github.com/ondrejsika/training-cli/releases/download/{{.Version}}/training-cli_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "statica",
		GetVersionFunc: func() string { return "v0.4.0" },
		UrlTemplate:    "https://github.com/vojtechmares/statica/releases/download/{{.Version}}/statica_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name: "kubectl",
		// Get Version: `curl -L -s https://dl.k8s.io/release/stable.txt`
		GetVersionFunc: func() string { return "v1.22.1" },
		UrlTemplate:    "https://dl.k8s.io/release/{{.Version}}/bin/{{.Os}}/{{.Arch}}/kubectl",
	},
	{
		Name:           "minikube",
		GetVersionFunc: func() string { return "latest" },
		UrlTemplate:    "https://storage.googleapis.com/minikube/releases/{{.Version}}/minikube-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "helm",
		GetVersionFunc: func() string { return "v3.6.3" },
		SourcePath:     "{{.Os}}-amd64/helm",
		UrlTemplate:    "https://get.helm.sh/helm-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:           "docker",
		SourcePath:     "docker/docker",
		GetVersionFunc: func() string { return "20.10.8" },
		UrlTemplate:    "https://download.docker.com/{{.OsDocker}}/static/stable/{{.ArchDocker}}/docker-{{.Version}}.tgz",
	},
	{
		Name:           "docker-compose",
		GetVersionFunc: func() string { return "1.29.2" },
		UrlTemplate:    "https://github.com/docker/compose/releases/download/{{.Version}}/docker-compose-{{.Os|capitalize}}-{{.ArchDocker}}",
	},
	{
		Name:        "mcli",
		UrlTemplate: "https://dl.minio.io/client/mc/release/{{.Os}}-{{.Arch}}/mc",
	},
	{
		Name:           "terraform",
		GetVersionFunc: func() string { return "1.0.5" },
		UrlTemplate:    hashicorpUrlTemplate("terraform"),
	},
	{
		Name:           "terraform13",
		GetVersionFunc: func() string { return "0.13.7" },
		SourcePath:     "terraform",
		UrlTemplate:    hashicorpUrlTemplate("terraform"),
	},
	{
		Name:           "vault",
		GetVersionFunc: func() string { return "1.8.2" },
		UrlTemplate:    hashicorpUrlTemplate("vault"),
	},
	{
		Name:           "doctl",
		GetVersionFunc: func() string { return "1.64.0" },
		UrlTemplate:    "https://github.com/digitalocean/doctl/releases/download/v{{.Version}}/doctl-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:           "skaffold",
		GetVersionFunc: func() string { return "latest" },
		UrlTemplate:    "https://storage.googleapis.com/skaffold/releases/{{.Version}}/skaffold-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "glab",
		GetVersionFunc: func() string { return "1.20.0" },
		UrlTemplate:    "https://github.com/profclems/glab/releases/download/v{{.Version}}/glab_{{.Version}}_{{.Os|capitalize}}_{{.ArchDocker}}.tar.gz",
		SourcePath:     "bin/glab",
	},
	{
		Name:           "alertmanager",
		Aliases:        []string{"am"},
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("prometheus", "alertmanager") },
		UrlTemplate:    "https://github.com/prometheus/alertmanager/releases/download/{{.Version}}/alertmanager-{{.Version|removev}}.{{.Os}}-{{.Arch}}.tar.gz",
		SourcePath:     "alertmanager-{{.Version|removev}}.{{.Os}}-{{.Arch}}/alertmanager",
	},
	{
		Name:           "amtool",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("prometheus", "alertmanager") },
		UrlTemplate:    "https://github.com/prometheus/alertmanager/releases/download/{{.Version}}/alertmanager-{{.Version|removev}}.{{.Os}}-{{.Arch}}.tar.gz",
		SourcePath:     "alertmanager-{{.Version|removev}}.{{.Os}}-{{.Arch}}/amtool",
	},
}

func hashicorpUrlTemplate(name string) string {
	return "https://releases.hashicorp.com/" + name + "/{{.Version}}/" + name +
		"_{{.Version}}_{{.Os}}_{{.Arch}}.zip"
}
