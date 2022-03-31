package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"opspaas/pkg/client"
	"opspaas/pkg/config"
	"opspaas/pkg/tools"
)

func GetDeployment(c *gin.Context) (deploymentList []v12.Deployment,err error) {
	logger := tools.InitLogger()
	var ns config.Namespaces
	c.ShouldBindJSON(&ns)
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("get clientSet failed",zap.String("err: ",err.Error()))
		return nil,err
	}
	list, err := clientSet.AppsV1().Deployments(ns.Name).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		logger.Info("get deploymentList failed",zap.String("err: ",err.Error()))
		return nil,err
	}
	items := list.Items
	return items,nil
}