package models

type Pod struct {
	Name       string            `json:"name"`
	Labels     map[string]string `json:"labels"`
	Namespace  string            `json:"namespace"`
	Containers []Container       `json:"containers"`
}
