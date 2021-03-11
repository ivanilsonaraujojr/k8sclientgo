package services

import (
	"errors"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAllPods() (*v1.PodList, error) {
	client, err := getClient()

	if err != nil {
		return nil, err
	}

	podsList, err := client.CoreV1().Pods("").List(metav1.ListOptions{})

	if err != nil {
		return nil, errors.New("Erro ao listar os Pods do cluster: " + err.Error())
	}

	return podsList, nil

}
