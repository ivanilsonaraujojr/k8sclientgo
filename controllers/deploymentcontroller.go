package controllers

import (
	"k8sclientgo/models"
	"k8sclientgo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeploymentController struct{}

//GetAllDeployments Retorna todos deployments do cluster
func (deployment *DeploymentController) GetAllDeployments(context *gin.Context) {

	deployments, err := services.GetAllDeployments()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deploymentsObjs := make([]models.Deployment, 0)

	for _, value := range deployments.Items {
		deploymentsObjs = append(deploymentsObjs, models.Deployment{Name: value.Name, Labels: value.Labels, Replicas: *value.Spec.Replicas, Namespace: value.Namespace})
	}

	context.JSON(http.StatusOK, deploymentsObjs)
}
