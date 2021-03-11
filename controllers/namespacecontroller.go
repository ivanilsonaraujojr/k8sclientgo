package controllers

import (
	"k8sclientgo/models"
	"k8sclientgo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NamespaceController struct{}

//GetAllNamespaces Retorna todos namespaces do cluster
func (namespace *NamespaceController) GetAllNamespaces(context *gin.Context) {
	namespaces, err := services.GetAllNamespaces()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	namespacesObjs := make([]models.Namespace, 0)

	for _, value := range namespaces.Items {
		namespacesObjs = append(namespacesObjs, models.Namespace{Name: value.Name, CreatedAt: value.CreationTimestamp.Time})
	}

	context.JSON(http.StatusOK, namespacesObjs)
}
