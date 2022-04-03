package router

import (
	"github.com/gin-gonic/gin"
	"opspaas/pkg/apis"
)

func InitRouter(r *gin.Engine) {
	r.GET("/namespaces",apis.GetNamespace)
	r.POST("/namespace",apis.CreateNamespace)
	r.DELETE("/namespace",apis.DeleteNamespace)
	r.GET("/deployments",apis.GetDeployment)
	r.GET("/allDeployments",apis.GetAllDeployment)
	r.POST("/createDeployment",apis.CreateDeployment)
	r.DELETE("/deleteDeployment",apis.DeleteDeployment)
	r.POST("/updateDeployment",apis.UpdateDeployment)
	r.POST("/createService",apis.CreateService)
	r.GET("/getNsService",apis.GetService)
	r.POST("/updateService",apis.UpdateService)
	r.GET("/getAllService",apis.GetAllService)
	r.DELETE("/deleteService",apis.DeleteService)
}
