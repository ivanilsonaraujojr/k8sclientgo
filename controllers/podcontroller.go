package controllers

import (
	"k8sclientgo/models"
	"k8sclientgo/services"
)

func GetAllPods() ([]models.Pod, error) {
	pods, err := services.GetAllPods()

	if err != nil {
		return nil, err
	}

	podsObjs := make([]models.Pod, 0)

	for _, value := range pods.Items {
		containers := make([]models.Container, 0)
		for _, container := range value.Spec.Containers {
			containers = append(containers, models.Container{Name: container.Name, Image: container.Image})
		}

		podsObjs = append(podsObjs, models.Pod{Name: value.Name, Labels: value.Labels,
			Namespace: value.Namespace, Containers: containers})
	}

	return podsObjs, nil
}
