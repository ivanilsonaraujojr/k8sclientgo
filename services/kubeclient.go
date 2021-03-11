package services

import (
	"errors"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getClient() (*kubernetes.Clientset, error) {
	kubeconfig := filepath.Join(".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		return nil, errors.New("Erro ao gerar kubeconfig: " + err.Error())
	}

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, errors.New("Erro ao gerar config: " + err.Error())
	}

	return client, nil
}
