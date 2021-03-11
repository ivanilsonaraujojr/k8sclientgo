package controllers

import (
	"k8sclientgo/models"
	"k8sclientgo/services"
)

//GetAllDeployments Retorna todos deployments do cluster
func GetAllDeployments() ([]models.Deployment, error) {

	deployments, err := services.GetAllDeployments()

	if err != nil {
		return nil, err
	}

	deploymentsObjs := make([]models.Deployment, 0)

	for _, value := range deployments.Items {
		deploymentsObjs = append(deploymentsObjs, models.Deployment{Name: value.Name, Labels: value.Labels, Replicas: *value.Spec.Replicas, Namespace: value.Namespace})
	}

	return deploymentsObjs, nil
}
