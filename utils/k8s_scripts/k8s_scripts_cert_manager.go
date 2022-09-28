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
