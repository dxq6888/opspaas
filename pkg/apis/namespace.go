package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opspaas/pkg/service"
)

type namespaces struct {
	Name string `json:"namespaces"`
}

func GetNamespace(c *gin.Context) {
	namespaces, err := service.GetNamespace()
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
	}
	c.JSON(http.StatusOK,namespaces)
}

func CreateNamespace(c *gin.Context) {
	namespace, err := service.CreateNamespace(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"create namespace failed"})
	}
	c.JSON(http.StatusOK,namespace)
	//var ns namespaces
	//err := c.ShouldBindJSON(&ns)
	//if err != nil {
	//	log.Fatalln("params is wrong!")
	//	return
	//}
	//c.JSON(http.StatusOK,gin.H{
	//	"ns":ns.Name,
	//})
}

func DeleteNamespace(c *gin.Context) {
	err := service.DeleteNamespace(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"delete namespace failed"})
	}
	c.JSON(http.StatusOK,gin.H{"message":"delete namespace success"})
}