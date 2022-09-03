module github.com/sikalabs/slu

go 1.16

require (
	github.com/aquasecurity/go-version v0.0.0-20210121072130-637058cfe492
	github.com/argoproj/argo-cd/v2 v2.4.11
	github.com/atotto/clipboard v0.1.4
	github.com/aws/aws-sdk-go v1.44.91
	github.com/cheggaaa/pb/v3 v3.1.0
	github.com/cloudflare/cloudflare-go v0.46.0
	github.com/digitalocean/godo v1.83.0
	github.com/docker/docker v20.10.17+incompatible
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/hashicorp/vault/api v1.7.2
	github.com/heroku/docker-registry-client v0.0.0-20211012143308-9463674c8930
	github.com/jpillora/go-tcp-proxy v1.0.2
	github.com/lib/pq v1.10.6
	github.com/mhale/smtpd v0.8.0
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/zerolog v1.28.0
	github.com/spf13/cobra v1.5.0
	github.com/xanzy/go-gitlab v0.68.2
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20220504211119-3d4a969bb56b
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.3.5
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
	gotest.tools/v3 v3.3.0 // indirect
	k8s.io/api v0.24.2
	k8s.io/apimachinery v0.24.3
	k8s.io/client-go v0.24.2
	software.sslmate.com/src/go-pkcs12 v0.2.0
)

replace (
	github.com/go-check/check => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
	k8s.io/api => k8s.io/api v0.23.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.3
	k8s.io/apiserver => k8s.io/apiserver v0.23.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.3
	k8s.io/client-go => k8s.io/client-go v0.23.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.3
	k8s.io/code-generator => k8s.io/code-generator v0.23.3
	k8s.io/component-base => k8s.io/component-base v0.23.3
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.3
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.3
	k8s.io/cri-api => k8s.io/cri-api v0.23.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.3
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.3
	k8s.io/kubectl => k8s.io/kubectl v0.23.3
	k8s.io/kubelet => k8s.io/kubelet v0.23.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.3
	k8s.io/metrics => k8s.io/metrics v0.23.3
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.3
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.3
)
