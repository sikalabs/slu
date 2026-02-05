package argocd_utils

import (
	"context"
	"log"

	argocdclient "github.com/argoproj/argo-cd/v3/pkg/apiclient"
	applicationpkg "github.com/argoproj/argo-cd/v3/pkg/apiclient/application"
	argoappv1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v3/util/errors"
	argoio "github.com/argoproj/argo-cd/v3/util/io"

	sessionpkg "github.com/argoproj/argo-cd/v3/pkg/apiclient/session"
	"github.com/argoproj/argo-cd/v3/util/io"

	"github.com/sikalabs/slu/lib/vault_cfa_service_token"
	"github.com/sikalabs/slu/utils/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ArgoCDGetToken(
	ctx context.Context,
	serverAddr string,
	insecure bool,
	username string,
	password string,
	cloudflareAccessTokenName string,
) string {
	headers := []string{}
	grpcWeb := false

	if cloudflareAccessTokenName != "" {
		clientID, clientSecret, err := vault_cfa_service_token.Get(cloudflareAccessTokenName)
		if err != nil {
			log.Fatal(err)
		}
		headers = cloudflareAccessHeaders(clientID, clientSecret)
		grpcWeb = true
	}

	opts := argocdclient.ClientOptions{
		ServerAddr: serverAddr,
		Insecure:   insecure,
		Headers:    headers,
		GRPCWeb:    grpcWeb,
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
	cloudflareAccessTokenName string,
) {
	headers := []string{}
	grpcWeb := false
	if cloudflareAccessTokenName != "" {
		clientID, clientSecret, err := vault_cfa_service_token.Get(cloudflareAccessTokenName)
		if err != nil {
			log.Fatal(err)
		}
		headers = cloudflareAccessHeaders(clientID, clientSecret)
		grpcWeb = true
	}

	opts := argocdclient.ClientOptions{
		ServerAddr: serverAddr,
		Insecure:   insecure,
		AuthToken:  authToken,
		Headers:    headers,
		GRPCWeb:    grpcWeb,
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

	// Try argocd secret first
	secret, err := secretClient.Get(context.TODO(), "argocd-initial-admin-secret", metav1.GetOptions{})
	if err == nil {
		return string(secret.Data["password"])
	}

	// If argocd namespace fails and we're using default namespace, try openshift-gitops
	if namespace == "argocd" {
		secretClient = clientset.CoreV1().Secrets("openshift-gitops")
		secret, err = secretClient.Get(context.TODO(), "openshift-gitops-cluster", metav1.GetOptions{})
		if err == nil {
			return string(secret.Data["admin.password"])
		}
	}

	log.Fatalf("could not find initial password in namespace '%s' or 'openshift-gitops'", namespace)
	return ""
}

func ArgoCDGetDomain(namespace string) (string, error) {
	var err error

	clientset, _, err := k8s.KubernetesClient()
	if err != nil {
		return "", err
	}

	ingressClient := clientset.NetworkingV1().Ingresses(namespace)

	ingress, err := ingressClient.Get(context.TODO(), "argocd-server", metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	rule := ingress.Spec.Rules[0]
	return rule.Host, nil
}

func ArgoCDGetDomainOrDie(namespace string) string {
	domain, err := ArgoCDGetDomain(namespace)
	if err != nil {
		log.Fatal(err)
	}
	return domain
}

func cloudflareAccessHeaders(clientID, clientSecret string) []string {
	return []string{
		"CF-Access-Client-Id:" + clientID,
		"CF-Access-Client-Secret:" + clientSecret,
	}
}
