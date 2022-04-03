package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"opspaas/pkg/service"
	"opspaas/pkg/tools"
)

func GetService(c *gin.Context) {
	logger := tools.InitLogger()
	serviceList, err := service.GetService(c)
	var services []string
	if err != nil {
		logger.Info("get serviceList failed",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,gin.H{"msg":err})
		return
	}
	for _, svc := range serviceList.Items {
		services = append(services,svc.Name)
	}
	c.JSON(http.StatusOK,gin.H{"msg":services})
}

func GetAllService(c *gin.Context) {
	logger := tools.InitLogger()
	allService, err := service.GetAllService()
	if err != nil {
		logger.Info("get allService failed",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,zap.String("err",err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg":allService})
}

func CreateService(c *gin.Context) {
	logger := tools.InitLogger()
	createService, err := service.CreateService(c)
	if err != nil {
		logger.Info("create service failed,",zap.String("err",err.Error()))
		c.JSON(http.StatusOK,gin.H{"msg":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":createService})
}

func UpdateService(c *gin.Context) {
	service, err := service.UpdateService(c)
	logger := tools.InitLogger()
	if err != nil {
		logger.Info("update service failed",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,gin.H{"msg":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":service})
}

func DeleteService(c *gin.Context) {
	logger := tools.InitLogger()
	err := service.DeleteService(c)
	if err != nil {
		logger.Info("delete service failed",zap.String("err",err.Error()))
		c.JSON(http.StatusOK,gin.H{"msg":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":"success"})
}