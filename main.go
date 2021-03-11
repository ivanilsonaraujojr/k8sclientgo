package main

import (
	"k8sclientgo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		deploymentController := new(controllers.DeploymentController)
		namespaceController := new(controllers.NamespaceController)
		podController := new(controllers.PodController)

		// Deployments
		v1.GET("/deployment", deploymentController.GetAllDeployments)

		// Namespaces
		v1.GET("/namespace", namespaceController.GetAllNamespaces)

		//Pods
		v1.GET("/pod", podController.GetAllPods)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	r.Run(":8080")
}
