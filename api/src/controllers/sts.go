package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Assumerolests vai gerar STS com a conta de destino
func assumerolests(Idaccount string) {
	role := fmt.Sprintf("arn:aws:iam::%s:role/ROLESTS", Idaccount)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	}))

	// Criar credencias com a conta de destino, usando o metodo AssumeRole
	creds := stscreds.NewCredentials(sess, "", func(p *stscreds.AssumeRoleProvider) {
		p.RoleARN = *aws.String(role)
		p.RoleSessionName = *aws.String("600")

	})

	// Criar uma sessão cliente com o EKS
	eksSvc := eks.New(sess, &aws.Config{Credentials: creds})

	// Com os crendencias obtidas com o AssumeRole, liste as informações do Cluster
	input := &eks.DescribeClusterInput{
		Name: aws.String(os.Getenv("NAME")),
	}
	result, err := eksSvc.DescribeCluster(input)
	if err != nil {
		log.Fatalf("Error calling DescribeCluster: %v", err)
	}

	// Com as credencias locais do STS, geramos o .kubeconfig com a funçãao newClientset
	clientset, err := newClientset(result.Cluster)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}

	// Criar o BatchV1 cliente
	batchClient := clientset.BatchV1().Jobs("NAMESPACE")

	namejob := "kube-"
	// Criar um novo job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: namejob,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					ServiceAccountName: "kube-downscaler",
					Containers: []v1.Container{
						{
							Name:  "downscaler",
							Image: "bitnami/kubectl:1.27",
							Command: []string{
								"sh",
								"-c",
								"kubectl patch deployment kube-downscaler -n teste-system --type='json'  -p='[{\"op\": \"replace\", \"path\": \"/spec/template/spec/containers/0/env\", \"value\": [{\"name\": \"DEFAULT_UPTIME\",\"value\": \"Mon-Fri 08:00-21:00 America/Sao_Paulo\"}]}]' ",
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}

	// Criar o job
	create, err := batchClient.Create(context.Background(), job, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created job %q.\n", create.GetObjectMeta().GetName())
}
