package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"opspaas/pkg/config"
	"opspaas/pkg/router"
)

func main() {
	engine := gin.Default()
	router.InitRouter(engine)
	err := engine.Run(fmt.Sprintf("%s:%d", config.GetString(config.ServerHost), config.GetInt(config.ServerPort)))
	if err != nil {
		log.Fatalln("start failed")
		return
	}
}
