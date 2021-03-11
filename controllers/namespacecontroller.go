package controllers

import (
	"k8sclientgo/models"
	"k8sclientgo/services"
)

//GetAllNamespaces Retorna todos namespaces do cluster
func GetAllNamespaces() ([]models.Namespace, error) {
	namespaces, err := services.GetAllNamespaces()

	if err != nil {
		return nil, err
	}

	namespacesObjs := make([]models.Namespace, 0)

	for _, value := range namespaces.Items {
		namespacesObjs = append(namespacesObjs, models.Namespace{Name: value.Name, CreatedAt: value.CreationTimestamp.Time})
	}

	return namespacesObjs, nil
}
