package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Role       int           `json:"role"`
	Deployment []*Deployment `json:"deployment" gorm:"many2many:user_deployment"`
}
