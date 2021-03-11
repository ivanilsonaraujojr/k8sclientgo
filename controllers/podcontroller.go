package controllers

import (
	"k8sclientgo/models"
	"k8sclientgo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PodController struct{}

func (pod *PodController) GetAllPods(context *gin.Context) {
	pods, err := services.GetAllPods()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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

	context.JSON(http.StatusOK, podsObjs)
}
