package k8s_scripts

func InstallCertManager(dry bool) {
	sh(`helm upgrade --install \
cert-manager cert-manager \
--repo https://charts.jetstack.io \
--create-namespace \
--namespace cert-manager \
--set installCRDs=true \
--wait`, dry)
}

func InstallClusterIssuer(email string, dry bool) {
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt
spec:
  acme:
    email: `+email+`
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-issuer-account-key
    solvers:
      - http01:
          ingress:
            class: nginx
EOF`, dry)
}

func InstallClusterIssuerCloudflare(
	email string,
	cloudflareToken string,
	dry bool,
) {
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt
spec:
  acme:
    email: `+email+`
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-issuer-account-key
    solvers:
      - dns01:
          cloudflare:
            apiTokenSecretRef:
              name: cloudflare-api-token
              key: api-token
EOF`, dry)
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: cloudflare-api-token
  namespace: cert-manager
stringData:
  api-token: `+cloudflareToken+`
EOF`, dry)
}

func InstallClusterIssuerZeroSSL(
	email string,
	keyId string,
	keySecret string,
	dry bool,
) {
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: zerossl
spec:
  acme:
    server: https://acme.zerossl.com/v2/DV90
    email: `+email+`
    privateKeySecretRef:
      name: zerossl
    externalAccountBinding:
      keyID: `+keyId+`
      keySecretRef:
        name: zerossl-eab
        key: secret
    solvers:
      - http01:
          ingress:
            class: nginx
EOF`, dry)
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: zerossl-eab
  namespace: cert-manager
stringData:
  secret: `+keySecret+`
EOF`, dry)
}
