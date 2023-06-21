package install_bin

import (
	"github.com/sikalabs/slu/utils/github_utils"
	"github.com/sikalabs/slu/utils/http_utils"
)

var Tools = []Tool{
	{
		Name:           "install-slu",
		GetVersionFunc: func() string { return "v0.1.0" },
		UrlTemplate:    "https://github.com/sikalabs/install-slu/releases/download/{{.Version}}/install-slu_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "slu",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "slu") },
		UrlTemplate:    "https://github.com/sikalabs/slu/releases/download/{{.Version}}/slu_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "tergum",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "tergum") },
		UrlTemplate:    "https://github.com/sikalabs/tergum/releases/download/{{.Version}}/tergum_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "gobble",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "gobble") },
		UrlTemplate:    "https://github.com/sikalabs/gobble/releases/download/{{.Version}}/gobble_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "signpost",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "signpost") },
		UrlTemplate:    "https://github.com/sikalabs/signpost/releases/download/{{.Version}}/signpost_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "training-cli",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("ondrejsika", "training-cli") },
		UrlTemplate:    "https://github.com/ondrejsika/training-cli/releases/download/{{.Version}}/training-cli_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "hello-world-server",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "hello-world-server") },
		UrlTemplate:    "https://github.com/sikalabs/hello-world-server/releases/download/{{.Version}}/hello-world-server_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "statica",
		GetVersionFunc: func() string { return "v0.4.0" },
		UrlTemplate:    "https://github.com/vojtechmares/statica/releases/download/{{.Version}}/statica_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name: "kubectl",
		GetVersionFunc: func() string {
			return http_utils.UrlGetToString("https://storage.googleapis.com/kubernetes-release/release/stable.txt")
		},
		UrlTemplate: "https://dl.k8s.io/release/{{.Version}}/bin/{{.Os}}/{{.Arch}}/kubectl",
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
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("docker", "compose") },
		UrlTemplate:    "https://github.com/docker/compose/releases/download/{{.Version}}/docker-compose-{{.Os}}-{{.ArchDocker}}",
	},
	{
		Name:           "mcli",
		GetVersionFunc: func() string { return "" },
		UrlTemplate:    "https://dl.minio.io/client/mc/release/{{.Os}}-{{.Arch}}/mc",
	},
	{
		Name:           "terraform",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("hashicorp", "terraform") },
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
		GetVersionFunc: func() string { return "1.12.0" },
		UrlTemplate:    hashicorpUrlTemplate("vault"),
	},
	{
		Name:           "packer",
		GetVersionFunc: func() string { return "1.8.3" },
		UrlTemplate:    hashicorpUrlTemplate("packer"),
	},
	{
		Name:           "consul",
		GetVersionFunc: func() string { return "1.15.2" },
		UrlTemplate:    hashicorpUrlTemplate("consul"),
	},
	{
		Name:           "doctl",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("digitalocean", "doctl") },
		UrlTemplate:    "https://github.com/digitalocean/doctl/releases/download/{{.Version}}/doctl-{{.Version|removev}}-{{.Os}}-{{.Arch}}.tar.gz",
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
		Name:           "prometheus",
		Aliases:        []string{"prom"},
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("prometheus", "prometheus") },
		UrlTemplate:    "https://github.com/prometheus/prometheus/releases/download/{{.Version}}/prometheus-{{.Version|removev}}.{{.Os}}-{{.Arch}}.tar.gz",
		SourcePath:     "prometheus-{{.Version|removev}}.{{.Os}}-{{.Arch}}/prometheus",
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
	{
		Name:           "lego",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("go-acme", "lego") },
		UrlTemplate:    "https://github.com/go-acme/lego/releases/download/{{.Version}}/lego_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
		SourcePath:     "lego",
	},
	{
		Name:           "rancher",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("rancher", "cli") },
		UrlTemplate:    "https://github.com/rancher/cli/releases/download/{{.Version}}/rancher-{{.Os}}-{{.Arch}}-{{.Version}}.tar.gz",
		SourcePath:     "./rancher-{{.Version}}/rancher",
	},
	{
		Name:           "k9s",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("derailed", "k9s") },
		UrlTemplate:    "https://github.com/derailed/k9s/releases/download/{{.Version}}/k9s_{{.Os}}_{{.Arch}}.tar.gz",
		SourcePath:     "k9s",
	},
	{
		Name:           "slack-cli",
		UrlTemplate:    "https://raw.githubusercontent.com/rockymadden/slack-cli/master/src/slack",
		GetVersionFunc: func() string { return "" },
	},
	{
		Name:           "configboard-cli",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("configboard", "configboard-cli") },
		UrlTemplate:    "https://github.com/configboard/configboard-cli/releases/download/{{.Version}}/configboard-cli_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "viddy",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sachaos", "viddy") },
		UrlTemplate:    "https://github.com/sachaos/viddy/releases/download/{{.Version}}/viddy_{{.Version|removev}}_{{.Os|capitalize}}_{{.Arch}}.tar.gz",
		GetArchFunc:    craneGetArch,
	},
	{
		Name:           "krew",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("kubernetes-sigs", "krew") },
		UrlTemplate:    "https://github.com/kubernetes-sigs/krew/releases/download/{{.Version}}/krew-{{.Os}}_{{.Arch}}.tar.gz",
		SourcePath:     "krew-{{.Os}}_{{.Arch}}",
	},
	{
		Name:           "k6",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("grafana", "k6") },
		UrlTemplate:    "https://github.com/grafana/k6/releases/download/{{.Version}}/k6-{{.Version}}-{{.OsK6}}-{{.Arch}}.zip",
		SourcePath:     "k6-{{.Version}}-{{.OsK6}}-{{.Arch}}/k6",
	},
	{
		Name:           "oc",
		GetVersionFunc: func() string { return "latest" },
		UrlTemplate:    "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/{{.Version}}/openshift-client-{{.Os}}{{.Arch}}.tar.gz",
		GetOsFunc:      openshiftGetOs,
		GetArchFunc:    openshiftGetArch,
	},
	{
		Name:           "openshift-install",
		GetVersionFunc: func() string { return "latest" },
		UrlTemplate:    "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/{{.Version}}/openshift-install-{{.Os}}{{.Arch}}.tar.gz",
		GetOsFunc:      openshiftGetOs,
		GetArchFunc:    openshiftGetArch,
	},
	{
		Name:           "butane",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("coreos", "butane") },
		UrlTemplate:    "https://github.com/coreos/butane/releases/download/{{.Version}}/butane-{{.Arch}}-{{.Os}}",
		GetArchFunc:    butaneGetArchFunc,
		GetOsFunc:      butaneGetOsFunc,
	},
	{
		Name:           "argocd-image-updater",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("argoproj-labs", "argocd-image-updater") },
		UrlTemplate:    "https://github.com/argoproj-labs/argocd-image-updater/releases/download/{{.Version}}/argocd-image-updater-{{.Os}}_{{.Arch}}",
	},
	{
		Name:           "usql",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("xo", "usql") },
		UrlTemplate:    "https://github.com/xo/usql/releases/download/{{.Version}}/usql-{{.Version|removev}}-{{.Os}}-{{.Arch}}.tar.bz2",
	},
	{
		Name:           "reg",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("genuinetools", "reg") },
		UrlTemplate:    "https://github.com/genuinetools/reg/releases/download/{{.Version}}/reg-{{.Os}}-amd64",
	},
	{
		Name:           "crane",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("google", "go-containerregistry") },
		UrlTemplate:    "https://github.com/google/go-containerregistry/releases/download/{{.Version}}/go-containerregistry_{{.Os|capitalize}}_{{.Arch}}.tar.gz",
		GetArchFunc:    craneGetArch,
	},
	{
		Name:           "editorconfig-checker",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("editorconfig-checker", "editorconfig-checker") },
		UrlTemplate:    "https://github.com/editorconfig-checker/editorconfig-checker/releases/download/{{.Version}}/ec-{{.Os}}-{{.Arch}}.tar.gz",
		SourcePath:     "bin/ec-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "thanos",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("thanos-io", "thanos") },
		UrlTemplate:    "https://github.com/thanos-io/thanos/releases/download/{{.Version}}/thanos-{{.Version|removev}}.{{.Os}}-{{.Arch}}.tar.gz",
		SourcePath:     "thanos-{{.Version|removev}}.{{.Os}}-{{.Arch}}/thanos",
	},
	{
		Name:           "kaf",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("birdayz", "kaf") },
		UrlTemplate:    "https://github.com/birdayz/kaf/releases/download/{{.Version}}/kaf_{{.Version|removev}}_{{.Os|capitalize}}_{{.Arch}}.tar.gz",
		GetArchFunc:    craneGetArch,
	},
	{
		Name:           "tflint",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("terraform-linters", "tflint") },
		UrlTemplate:    "https://github.com/terraform-linters/tflint/releases/download/{{.Version}}/tflint_{{.Os}}_{{.Arch}}.zip",
	},
	{
		Name:           "filebeat",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("elastic", "beats") },
		UrlTemplate:    "https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-{{.Version|removev}}-{{.Os}}-{{.Arch}}.tar.gz",
		GetArchFunc:    dockerArch,
		SourcePath:     "filebeat-{{.Version|removev}}-{{.Os}}-{{.Arch}}/filebeat",
	},
	{
		Name:           "nerdctl",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("containerd", "nerdctl") },
		UrlTemplate:    "https://github.com/containerd/nerdctl/releases/download/{{.Version}}/nerdctl-{{.Version|removev}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:           "helmfile",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("helmfile", "helmfile") },
		UrlTemplate:    "https://github.com/helmfile/helmfile/releases/download/{{.Version}}/helmfile_{{.Version|removev}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "kubectx",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("ahmetb", "kubectx") },
		UrlTemplate:    "https://github.com/ahmetb/kubectx/releases/download/{{.Version}}/kubectx_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
		GetArchFunc:    craneGetArch,
	},
	{
		Name:           "kubens",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("ahmetb", "kubectx") },
		UrlTemplate:    "https://github.com/ahmetb/kubectx/releases/download/{{.Version}}/kubens_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
		GetArchFunc:    craneGetArch,
	},
	{
		Name:           "yq",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("mikefarah", "yq") },
		UrlTemplate:    "https://github.com/mikefarah/yq/releases/download/{{.Version}}/yq_{{.Os}}_{{.Arch}}.tar.gz",
		SourcePath:     "yq_{{.Os}}_{{.Arch}}",
	},
	{
		Name:           "argocd-vault-plugin",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("argoproj-labs", "argocd-vault-plugin") },
		UrlTemplate:    "https://github.com/argoproj-labs/argocd-vault-plugin/releases/download/{{.Version}}/argocd-vault-plugin_{{.Version|removev}}_{{.Os}}_{{.Arch}}",
	},
	{
		Name:           "ctop",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("bcicen", "ctop") },
		UrlTemplate:    "https://github.com/bcicen/ctop/releases/download/{{.Version}}/ctop-{{.Version|removev}}-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "caddy",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("caddyserver", "caddy") },
		UrlTemplate:    "https://github.com/caddyserver/caddy/releases/download/{{.Version}}/caddy_{{.Version|removev}}_{{.Os}}_{{.Arch}}.tar.gz",
		GetOsFunc:      openshiftGetOs,
	},
	{
		Name:           "hadolint",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("hadolint", "hadolint") },
		UrlTemplate:    "https://github.com/hadolint/hadolint/releases/download/{{.Version}}/hadolint-{{.Os|capitalize}}-{{.Arch}}",
		GetArchFunc:    hadolintGetArchFunc,
	},
}

func hashicorpUrlTemplate(name string) string {
	return "https://releases.hashicorp.com/" + name + "/{{.Version|removev}}/" + name +
		"_{{.Version|removev}}_{{.Os}}_{{.Arch}}.zip"
}

func openshiftGetOs(x string) string {
	if x == "darwin" {
		return "mac"
	}
	return x
}

func openshiftGetArch(x string) string {
	if x == "amd64" {
		return ""
	}
	if x == "arm64" {
		return "-arm64"
	}
	return x
}

func craneGetArch(x string) string {
	if x == "amd64" {
		return "x86_64"
	}
	return x
}

func butaneGetArchFunc(arch string) string {
	if arch == "amd64" {
		return "x86_64"
	}
	if arch == "arm64" {
		return "aarch64"
	}
	return arch
}

func butaneGetOsFunc(os string) string {
	if os == "darwin" {
		return "apple-darwin"
	}
	if os == "linux" {
		return "unknown-linux-gnu"
	}
	if os == "windows" {
		return "pc-windows-gnu"
	}
	return os
}

func hadolintGetArchFunc(arch string) string {
	if arch == "amd64" {
		return "x86_64"
	}
	return arch
}
