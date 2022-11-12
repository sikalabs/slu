module github.com/sikalabs/slu

go 1.16

require (
	github.com/aquasecurity/go-version v0.0.0-20210121072130-637058cfe492
	github.com/argoproj/argo-cd/v2 v2.4.14
	github.com/atotto/clipboard v0.1.4
	github.com/aws/aws-sdk-go v1.44.136
	github.com/cheggaaa/pb/v3 v3.1.0
	github.com/cloudflare/cloudflare-go v0.52.0
	github.com/denisenkom/go-mssqldb v0.12.2
	github.com/digitalocean/godo v1.83.0
	github.com/docker/docker v20.10.21+incompatible
	github.com/frankban/quicktest v1.14.3 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/vault/api v1.7.2
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/heroku/docker-registry-client v0.0.0-20211012143308-9463674c8930
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jhump/protoreflect v1.6.1 // indirect
	github.com/jpillora/go-tcp-proxy v1.0.2
	github.com/lib/pq v1.10.7
	github.com/matryer/is v1.4.0 // indirect
	github.com/mhale/smtpd v0.8.0
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.4 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/zerolog v1.28.0
	github.com/russross/blackfriday v1.6.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cobra v1.6.1
	github.com/xanzy/go-gitlab v0.73.1
	go.starlark.net v0.0.0-20200821142938-949cc6f4b097 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20220504211119-3d4a969bb56b
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/sqlite v1.4.2
	gorm.io/gorm v1.24.0
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
