package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opspaas/pkg/service"
)


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
		return
	}
	//c.JSON(http.StatusOK,namespace)
	c.JSON(http.StatusOK,gin.H{"message":"create namespace success","namespace":namespace})
}

func DeleteNamespace(c *gin.Context) {
	err := service.DeleteNamespace(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"delete namespace success"})
}