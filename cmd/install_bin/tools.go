package install_bin

import (
	"os"

	"github.com/sikalabs/slu/internal/k3s_utils"
	"github.com/sikalabs/slu/utils/github_utils"
	"github.com/sikalabs/slu/utils/http_utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		RunBeforeInstall: func(name string, version string, os_ string, arch string, binDir string) error {
			if binDir == "/usr/local/bin" && k3s_utils.CheckIfKubectlIsLinkOfK3s() {
				err := os.Remove("/usr/local/bin/kubectl")
				if err != nil {
					return err
				}
			}
			return nil
		},
	},
	{
		Name:           "minikube",
		GetVersionFunc: func() string { return "latest" },
		UrlTemplate:    "https://storage.googleapis.com/minikube/releases/{{.Version}}/minikube-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "helm",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("helm", "helm") },
		SourcePath:     "{{.Os}}-{{.Arch}}/helm",
		UrlTemplate:    "https://get.helm.sh/helm-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:       "docker",
		SourcePath: "docker/docker",
		GetVersionFunc: func() string {
			// the version is like docker-v29.0.2 (checkout slu gh latest-release moby/moby)
			return github_utils.GetLatestRelease("moby", "moby")
		},
		UrlTemplate: "https://download.docker.com/{{.OsDocker}}/static/stable/{{.ArchDocker}}/{{.Version|removev}}.tgz",
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
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("hashicorp", "vault") },
		UrlTemplate:    hashicorpUrlTemplate("vault"),
	},
	{
		Name:           "packer",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("hashicorp", "packer") },
		UrlTemplate:    hashicorpUrlTemplate("packer"),
	},
	{
		Name:           "consul",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("hashicorp", "consul") },
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
		UrlTemplate:    `https://github.com/grafana/k6/releases/download/{{.Version}}/k6-{{.Version}}-{{.OsK6}}-{{.Arch}}.{{ if eq .Os "linux" }}tar.gz{{ else }}zip{{ end }}`,
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
		Name:           "openshift-install-okd",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("okd-project", "okd") },
		UrlTemplate:    "https://github.com/okd-project/okd/releases/download/{{.Version}}/openshift-install-{{.Os}}{{.Arch}}-{{.Version}}.tar.gz",
		GetOsFunc:      openshiftGetOs,
		GetArchFunc:    openshiftGetArch,
		SourcePath:     "openshift-install",
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
	{
		Name:           "oauth2-proxy",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("oauth2-proxy", "oauth2-proxy") },
		UrlTemplate:    "https://github.com/oauth2-proxy/oauth2-proxy/releases/download/{{.Version}}/oauth2-proxy-{{.Version}}.{{.Os}}-{{.Arch}}.tar.gz",
		SourcePath:     "oauth2-proxy-{{.Version}}.{{.Os}}-{{.Arch}}/oauth2-proxy",
		GetArchFunc:    func(_ string) string { return "amd64" },
	},
	{
		Name:           "goexpandenv",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "goexpandenv") },
		UrlTemplate:    "https://github.com/sikalabs/goexpandenv/releases/download/{{.Version}}/goexpandenv_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "cloudflared",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("cloudflare", "cloudflared") },
		UrlTemplate:    "https://github.com/cloudflare/cloudflared/releases/download/{{.Version}}/cloudflared-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "clicksecret-cli",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "clicksecret-cli") },
		UrlTemplate:    "https://github.com/sikalabs/clicksecret-cli/releases/download/{{.Version}}/clicksecret-cli_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "flog",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("mingrammer", "flog") },
		UrlTemplate:    "https://github.com/mingrammer/flog/releases/download/{{.Version}}/flog_{{.Version|removev}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "slc",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "slc") },
		UrlTemplate:    "https://github.com/sikalabs/slc/releases/download/{{.Version}}/slc_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "kubelogin",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("Azure", "kubelogin") },
		UrlTemplate:    "https://github.com/Azure/kubelogin/releases/download/{{.Version}}/kubelogin-{{.Os}}-{{.Arch}}.zip",
		SourcePath:     "bin/{{.Os}}_{{.Arch}}/kubelogin",
	},
	{
		Name:           "mon",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "mon") },
		UrlTemplate:    "https://github.com/sikalabs/mon/releases/download/{{.Version}}/mon_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "slr",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "slr") },
		UrlTemplate:    "https://github.com/sikalabs/slr/releases/download/{{.Version}}/slr_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "kubeseal",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("bitnami-labs", "sealed-secrets") },
		UrlTemplate:    "https://github.com/bitnami-labs/sealed-secrets/releases/download/{{.Version}}/kubeseal-{{.Version|removev}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:           "goreleaser",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("goreleaser", "goreleaser") },
		UrlTemplate:    "https://github.com/goreleaser/goreleaser/releases/download/{{.Version}}/goreleaser_{{.Os}}_{{.Arch}}.tar.gz",
		GetOsFunc:      func(x string) string { return cases.Title(language.Und).String(x) },
		GetArchFunc:    craneGetArch,
	},
	{
		Name:           "loki",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("grafana", "loki") },
		UrlTemplate:    "https://github.com/grafana/loki/releases/download/{{.Version}}/loki-{{.Os}}-{{.Arch}}.zip",
		SourcePath:     "loki-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "promtail",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("grafana", "loki") },
		UrlTemplate:    "https://github.com/grafana/loki/releases/download/{{.Version}}/promtail-{{.Os}}-{{.Arch}}.zip",
		SourcePath:     "promtail-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "logcli",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("grafana", "loki") },
		UrlTemplate:    "https://github.com/grafana/loki/releases/download/{{.Version}}/logcli-{{.Os}}-{{.Arch}}.zip",
		SourcePath:     "logcli-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "alloy",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("grafana", "alloy") },
		UrlTemplate:    "https://github.com/grafana/alloy/releases/download/{{.Version}}/alloy-{{.Os}}-{{.Arch}}.zip",
		SourcePath:     "alloy-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "coredns",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("coredns", "coredns") },
		UrlTemplate:    "https://github.com/coredns/coredns/releases/download/{{.Version}}/coredns_{{.Version|removev}}_{{.Os}}_{{.Arch}}.tgz",
	},
	{
		Name:           "rclone",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("rclone", "rclone") },
		UrlTemplate:    "https://github.com/rclone/rclone/releases/download/{{.Version}}/rclone-{{.Version}}-{{.Os}}-{{.Arch}}.zip",
		SourcePath:     "rclone-{{.Version}}-{{.Os}}-{{.Arch}}/rclone",
		GetOsFunc:      rcloneGetOsFunc,
	},
	{
		Name:           "dogsay",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabs", "dogsay") },
		UrlTemplate:    "https://github.com/sikalabs/dogsay/releases/download/{{.Version}}/dogsay_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "counter",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("ondrejsika", "counter") },
		UrlTemplate:    "https://github.com/ondrejsika/counter/releases/download/{{.Version}}/counter_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "master",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabsx", "master") },
		UrlTemplate:    "https://github.com/sikalabsx/master/releases/download/{{.Version}}/master_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "master_slu",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabsx", "master") },
		UrlTemplate:    "https://github.com/sikalabsx/master/releases/download/{{.Version}}/master_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "master_slr",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabsx", "master") },
		UrlTemplate:    "https://github.com/sikalabsx/master/releases/download/{{.Version}}/master_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "master_tergum",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("sikalabsx", "master") },
		UrlTemplate:    "https://github.com/sikalabsx/master/releases/download/{{.Version}}/master_{{.Version}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "k3d",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("k3d-io", "k3d") },
		UrlTemplate:    "https://github.com/k3d-io/k3d/releases/download/{{.Version}}/k3d-{{.Os}}-{{.Arch}}",
	},
	{
		Name:           "terragrunt",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("gruntwork-io", "terragrunt") },
		UrlTemplate:    "https://github.com/gruntwork-io/terragrunt/releases/download/{{.Version}}/terragrunt_{{.Os}}_{{.Arch}}",
	},
	{
		Name:           "asdf",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("asdf-vm", "asdf") },
		UrlTemplate:    "https://github.com/asdf-vm/asdf/releases/download/{{.Version}}/asdf-{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
	},
	{
		Name:           "kube-score",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("zegl", "kube-score") },
		UrlTemplate:    "https://github.com/zegl/kube-score/releases/download/{{.Version}}/kube-score_{{.Version|removev}}_{{.Os}}_{{.Arch}}.tar.gz",
	},
	{
		Name:           "migrate",
		GetVersionFunc: func() string { return github_utils.GetLatestRelease("golang-migrate", "migrate") },
		UrlTemplate:    "https://github.com/golang-migrate/migrate/releases/download/{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
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

func rcloneGetOsFunc(os string) string {
	if os == "darwin" {
		return "osx"
	}
	return os
}
