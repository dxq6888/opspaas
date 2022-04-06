package model

import "gorm.io/gorm"

type Ingress struct {
	gorm.Model
	Name string `json:"name"`
}
