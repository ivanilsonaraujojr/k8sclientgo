package main

import (
	"k8sclientgo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/deployment", func(context *gin.Context) {
		deployments, err := controllers.GetAllDeployments()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, deployments)
	})

	r.GET("/namespace", func(context *gin.Context) {
		namespaces, err := controllers.GetAllNamespaces()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, namespaces)
	})

	r.GET("/pod", func(context *gin.Context) {
		pods, err := controllers.GetAllPods()

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, pods)
	})

	_ = r.Run()
}
