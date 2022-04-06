package model

import "gorm.io/gorm"

type Deployment struct {
	gorm.Model
	Name string `json:"name"`
	User []User `json:"user" gorm:"many2many:user_deployment"`
}
