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

func CreateDeployment(c *gin.Context) {
	logger := tools.InitLogger()
	deployment, err := service.CreateDeployment(c)
	if err != nil {
		logger.Info("create deployment failed,",zap.String("err",err.Error()))
		c.JSON(http.StatusOK,gin.H{"message":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":deployment})
}

func DeleteDeployment(c *gin.Context) {
	logger := tools.InitLogger()
	err := service.DeleteDeployment(c)
	if err != nil {
		logger.Info("delete deployment failed,",zap.String("err",err.Error()))
		c.JSON(http.StatusInternalServerError,gin.H{"msg":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":"success"})
}

func UpdateDeployment()  {
	
}