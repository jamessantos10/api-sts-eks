package controllers

import (
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

// newClientset responsavel por gerar o .kubeconfig, com as credencias do STS
func newClientset(cluster *eks.Cluster) (*kubernetes.Clientset, error) {
	log.Printf("%+v", cluster)
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return nil, err
	}

	// Variável responsável por gerar token com o Cluster
	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(cluster.Name),
	}

	// Obter um token com informações fornecidas da variável opts
	tok, err := gen.GetWithOptions(opts)

	// Obeter o certificado do Cluster e o decode para base64
	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(cluster.CertificateAuthority.Data))
	if err != nil {
		return nil, err
	}

	// Gera um cliente com as informações as consultas feitas
	clientset, err := kubernetes.NewForConfig(
		&rest.Config{
			Host:        aws.StringValue(cluster.Endpoint),
			BearerToken: tok.Token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: ca,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
