package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"opspaas/pkg/client"
	"opspaas/pkg/config"
	"opspaas/pkg/tools"
)

func GetNamespace() ([]v12.Namespace,error) {
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		log.Fatalln("can not get clientSet")
		return nil,err
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalln("can not get namespaceList")
		return nil, err
	}
	return namespaceList.Items,nil
}


func CreateNamespace(c *gin.Context) (name string,err error) {
	logger := tools.InitLogger()
	var ns config.Namespaces
	err = c.ShouldBindJSON(&ns)
	if err != nil {
		logger.Info("params wrong")
		return "", err
	}
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("can not get clientSet")
		return "", err
	}
	namespaces := &v12.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: ns.Name,
		},
		Status: v12.NamespaceStatus{
			Phase: v12.NamespaceActive,
		},
	}
	result, err := clientSet.CoreV1().Namespaces().Create(context.TODO(), namespaces, v1.CreateOptions{})
	if err != nil {
		logger.Info("create namespace error: ",zap.String("err: ",err.Error()))
		//logger.Info("create namespace error: ")
		fmt.Println(err)
		return "", err
	}
	name = result.Name
	return name,nil
}

func DeleteNamespace(c *gin.Context) error {
	logger := tools.InitLogger()
	var ns config.Namespaces
	err := c.ShouldBindJSON(&ns)
	if err != nil {
		log.Println("params is wrong!")
		return err
	}
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		log.Fatalln("can not get clientSet")
		return err
	}
	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), ns.Name, v1.DeleteOptions{})
	if err != nil {
		logger.Info("delete namespace failed",zap.String("err: ",err.Error()))
		return err
	}
	return nil
}