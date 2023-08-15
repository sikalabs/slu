package k8s_scripts

func CreateClusterAdmin(suffix string, dry bool) {
	sh(`cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cluster-admin-`+suffix+`
  namespace: kube-system
secrets:
  - name: cluster-admin-`+suffix+`
---
apiVersion: v1
kind: Secret
metadata:
  name: cluster-admin-`+suffix+`
  namespace: kube-system
  annotations:
    kubernetes.io/service-account.name: cluster-admin-`+suffix+`
type: kubernetes.io/service-account-token
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cluster-admin-`+suffix+`
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: cluster-admin-`+suffix+`
    namespace: kube-system
EOF`, dry)
}
