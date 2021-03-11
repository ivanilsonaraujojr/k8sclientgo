package services

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
)

func GetAllNamespaces() (*corev1.NamespaceList, error) {
	client, err := getClient()

	if err != nil {
		return nil, err
	}

	namespacesRequest := client.AppsV1().RESTClient().Get().RequestURI("api/v/namespaces").Do()
	names, err := namespacesRequest.Get()

	if err != nil {
		return nil, errors.New("Erro ao listar os Namespaces do cluster: " + err.Error())
	}

	return names.(*corev1.NamespaceList), nil
}
