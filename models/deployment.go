package models

type Deployment struct {
	Name      string            `json:"name"`
	Labels    map[string]string `json:"labels"`
	Replicas  int32             `json:"replicas"`
	Namespace string            `json:"namespace"`
}
