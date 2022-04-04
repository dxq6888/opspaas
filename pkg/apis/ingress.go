package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"opspaas/pkg/service"
	"opspaas/pkg/tools"
)

func GetIngress(c *gin.Context)  {
	logger := tools.InitLogger()
	ingress, err := service.GetIngress(c)
	var ingresses []string
	if err != nil {
		logger.Info("get ingressList failed",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,gin.H{"msg":err})
		return
	}
	for _, ing := range ingress.Items {
		ingresses = append(ingresses,ing.Name)
	}
	c.JSON(http.StatusOK,gin.H{"msg":ingresses})
}

func GetAllIngress(c *gin.Context) {
	logger := tools.InitLogger()
	ingress, err := service.GetAllIngress()
	if err != nil {
		logger.Info("get ingress failed",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusOK,ingress)
}

func CreateIngress(c *gin.Context) {
	logger := tools.InitLogger()
	ingress, err := service.CreateIngress(c)
	if err != nil {
		logger.Info("create ingress failed",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,gin.H{"msg":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":ingress})
}


