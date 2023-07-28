package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteNoReadyPods(clientset *kubernetes.Clientset, force bool) error {
	var err error
	namespaceClient := clientset.CoreV1().Namespaces()
	namespaces, err := namespaceClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, namespace := range namespaces.Items {
		podClient := clientset.CoreV1().Pods(namespace.Name)
		pods, err := podClient.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		for _, pod := range pods.Items {
			if pod.Status.Phase != "Running" {
				fmt.Printf("Deleting %s pod %s in namespace %s\n", pod.Status.Phase, pod.Name, namespace.Name)
				deleteOptions := metav1.DeleteOptions{}
				if force {
					var zero int64 = 0
					deleteOptions = metav1.DeleteOptions{
						GracePeriodSeconds: &zero,
					}
				}
				err = podClient.Delete(context.TODO(), pod.Name, deleteOptions)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}
