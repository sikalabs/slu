package eck_utils

import (
	"context"
	"log"

	"github.com/sikalabs/slu/utils/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetElasticPassword(namespace string, name string) string {
	clientset, _, _ := k8s.KubernetesClient()

	secretClient := clientset.CoreV1().Secrets(namespace)

	secret, err := secretClient.Get(context.TODO(), name+"-es-elastic-user", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return string(secret.Data["elastic"])
}
