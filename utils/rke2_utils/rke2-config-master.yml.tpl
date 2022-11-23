{{ if not .IsFirstMaster }}
server: https://{{.ServerDomain}}:9345
{{ end -}}
{{ if .TlsSans }}
tls-san:
{{ end }}
{{ range $san := .TlsSans }}
- {{ $san }}
{{ end }}
token: {{.Token}}
node-taint:
  - "CriticalAddonsOnly=true:NoExecute"
disable:
  - rke2-ingress-nginx
