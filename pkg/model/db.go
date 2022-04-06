package model

import (
	"fmt"
	"gorm.io/gorm"
	"opspaas/pkg/config"
)

func InitDb(opts ...string) {
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetString(config.MyUser),
		config.GetString(config.MyPasswd),
		config.GetString(config.MyHost),
		config.GetString(config.MyPort),
		config.GetString(config.MyDb),
	)
	gorm.Open()
}
