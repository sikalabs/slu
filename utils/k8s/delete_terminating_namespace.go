package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

func DeleteTerminatingNamespace(
	clientset *kubernetes.Clientset,
	namespace string,
) error {
	var err error
	client := clientset.CoreV1().RESTClient()
	err = client.Put().
		Resource("namespaces").
		Name(namespace).
		SubResource("finalize").
		VersionedParams(&metav1.GetOptions{}, scheme.ParameterCodec).
		Body(&v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
			Spec: v1.NamespaceSpec{
				Finalizers: nil,
			},
		}).
		Do(context.TODO()).
		Into(&v1.Namespace{})
	return err
}
