package argocd_utils

import (
	"context"
	"log"

	argocdclient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	applicationpkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	argoappv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/util/errors"
	argoio "github.com/argoproj/argo-cd/v2/util/io"

	sessionpkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/session"
	"github.com/argoproj/argo-cd/v2/util/io"

	"github.com/sikalabs/slu/utils/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ArgoCDGetToken(
	ctx context.Context,
	serverAddr string,
	insecure bool,
	username string,
	password string,
) string {
	opts := argocdclient.ClientOptions{
		ServerAddr: serverAddr,
		Insecure:   insecure,
	}

	acdClient, err := argocdclient.NewClient(&opts)
	errors.CheckError(err)
	sessConn, sessionIf := acdClient.NewSessionClientOrDie()
	defer io.Close(sessConn)

	sessionRequest := sessionpkg.SessionCreateRequest{
		Username: username,
		Password: password,
	}
	createdSession, err := sessionIf.Create(ctx, &sessionRequest)
	errors.CheckError(err)
	return createdSession.Token
}

func ArgoCDRefresh(
	ctx context.Context,
	serverAddr string,
	insecure bool,
	authToken string,
	appName string,
) {
	opts := argocdclient.ClientOptions{
		ServerAddr: serverAddr,
		Insecure:   insecure,
		AuthToken:  authToken,
	}

	acdClient, err := argocdclient.NewClient(&opts)
	errors.CheckError(err)
	conn, appIf := acdClient.NewApplicationClientOrDie()
	defer argoio.Close(conn)

	refreshType := string(argoappv1.RefreshTypeNormal)
	_, err = appIf.Get(ctx, &applicationpkg.ApplicationQuery{
		Name:    &appName,
		Refresh: &refreshType,
	})
	errors.CheckError(err)
}

func ArgoCDGetInitialPassword(namespace string) string {
	clientset, _, _ := k8s.KubernetesClient()

	secretClient := clientset.CoreV1().Secrets(namespace)

	secret, err := secretClient.Get(context.TODO(), "argocd-initial-admin-secret", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return string(secret.Data["password"])
}
