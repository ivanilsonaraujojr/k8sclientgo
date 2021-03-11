package models

import "time"

type Namespace struct {
	Name      string    `json:"namespace"`
	CreatedAt time.Time `json:"created_at"`
}
