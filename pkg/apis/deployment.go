package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"opspaas/pkg/service"
	"opspaas/pkg/tools"
)

func GetDeployment(c *gin.Context) {
	logger := tools.InitLogger()
	deploymentList, err := service.GetDeployment(c)
	if err != nil {
		logger.Info("get deployment failed: ", zap.String("err: ", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":deploymentList})
}

func GetAllDeployment(c *gin.Context)  {
	logger := tools.InitLogger()
	allDeployment, err := service.GetAllDeployment(c)
	if err != nil {
		logger.Info("get AllDeployment failed,",zap.String("err: ",err.Error()))
		c.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":allDeployment})
}