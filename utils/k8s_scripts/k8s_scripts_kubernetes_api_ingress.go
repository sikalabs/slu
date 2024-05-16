package k8s_scripts

func ApplyKubernetesApiIngress(domain string, dry bool) {
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kubernetes-api
  namespace: default
  annotations:
    cert-manager.io/cluster-issuer: letsencrypt
    nginx.ingress.kubernetes.io/backend-protocol: HTTPS
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - `+domain+`
    secretName: `+domain+`-tls
  rules:
  - host: `+domain+`
    http:
      paths:
      - backend:
          service:
            name: kubernetes
            port:
              number: 443
        path: /
        pathType: Prefix
EOF`, dry)
}
