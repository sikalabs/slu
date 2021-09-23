package k8s_lock_utils

import (
	"context"

	"github.com/sikalabs/slu/utils/k8s"
	"github.com/sikalabs/slu/version"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func getConfigMapClient(namespace string) corev1.ConfigMapInterface {
	clientset, defaultNamespace, _ := k8s.KubernetesClient()
	if namespace == "" {
		namespace = defaultNamespace
	}
	configmapClient := clientset.CoreV1().ConfigMaps(namespace)
	return configmapClient
}

func CheckLock(name, namespace string) bool {
	cmc := getConfigMapClient(namespace)
	configMaps, _ := cmc.List(context.TODO(), metav1.ListOptions{
		LabelSelector: "lock=true",
	})
	for _, cm := range configMaps.Items {
		if cm.Name == "lock-"+name {
			return true
		}
	}
	return false
}

func createLock(name, namespace string) error {
	cmc := getConfigMapClient(namespace)
	cmc.Create(
		context.TODO(),
		&v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: "lock-" + name,
				Labels: map[string]string{
					"lock": "true",
				},
				Annotations: map[string]string{
					"slu.version": version.Version,
					"slu.lock":    "true",
				},
			},
			Data: map[string]string{
				"locked": "true",
			},
		},
		metav1.CreateOptions{})
	return nil
}

func deleteLock(name, namespace string) error {
	cmc := getConfigMapClient(namespace)
	return cmc.Delete(
		context.TODO(),
		"lock-"+name,
		metav1.DeleteOptions{})
}

func Lock(lockName, namespace string) (bool, error) {
	isLocked := CheckLock(lockName, namespace)
	if isLocked {
		return false, nil
	} else {
		createLock(lockName, namespace)
		return true, nil
	}
}

func Unlock(lockName, namespace string) (bool, error) {
	isLocked := CheckLock(lockName, namespace)
	if isLocked {
		deleteLock(lockName, namespace)
		return false, nil
	} else {
		return false, nil
	}
}
