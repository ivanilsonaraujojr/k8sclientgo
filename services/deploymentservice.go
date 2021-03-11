package services

import (
	"errors"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAllDeployments() (*v1.DeploymentList, error) {
	client, err := getClient()

	if err != nil {
		return nil, err
	}

	deployments, err := client.AppsV1().Deployments("").List(metav1.ListOptions{})

	if err != nil {
		return nil, errors.New("Erro ao listar os Deployments: " + err.Error())
	}

	return deployments, nil
}

func GetDeploymentAnnotations(deploymentName, namespace string) (map[string]string, error) {
	client, err := getClient()

	if err != nil {
		return nil, err
	}

	deployment, err := client.AppsV1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, errors.New("Erro ao listar anotações do deployment " + deploymentName + ": " + err.Error())
	}

	return deployment.ObjectMeta.Annotations, nil
}
