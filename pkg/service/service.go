package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"opspaas/pkg/client"
	"opspaas/pkg/tools"
	"opspaas/resource"
)

func GetService(c *gin.Context) (*v1.ServiceList,error) {
	logger := tools.InitLogger()
	var service resource.Service
	c.ShouldBindJSON(&service)
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("get namespace failed",zap.String("err",err.Error()))
		return nil,err
	}
	serviceList, err := clientSet.CoreV1().Services(service.Namespace).List(context.TODO(), v12.ListOptions{})
	if err != nil {
		logger.Info("get serviceList failed",zap.String("err",err.Error()))
		return nil,err
	}
	return serviceList,nil
}

func GetAllService() (svc []string , err error) {
	logger := tools.InitLogger()
	clientSet, err := client.GetK8sClientset()
	//var ns config.Namespaces
	//c.ShouldBindJSON(&ns)
	namespaces, err := GetNamespace()
	if err != nil {
		logger.Info("get namespaces failed",zap.String("err",err.Error()))
		return nil,err
	}
	for _, namespace := range namespaces {
		serviceList, err := clientSet.CoreV1().Services(namespace.Name).List(context.TODO(), v12.ListOptions{})
		if err != nil {
			logger.Info("get serviceList failed",zap.String("err",err.Error()))
			return nil, err
		}
		for _, service := range serviceList.Items {
			svc = append(svc,service.Name)
		}
	}
	return svc,nil
}

func CreateService(c *gin.Context) (*v1.Service, error) {
	logger := tools.InitLogger()
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("get client failed,", zap.String("err", err.Error()))
	}
	var service resource.Service
	c.ShouldBindJSON(&service)
	svc := &v1.Service{
		ObjectMeta: v12.ObjectMeta{
			Name:      service.Name,
			Namespace: service.Namespace,
			Labels:    service.Selector,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:     fmt.Sprintf(service.Name + "-" + fmt.Sprint(service.Ports[0])),
					Protocol: service.Protocol,
					Port:     service.Ports[0],
				},
			},
			Type:     service.Type,
			Selector: service.Selector,
		},
	}
	create, err := clientSet.CoreV1().Services(service.Namespace).Create(context.TODO(), svc, v12.CreateOptions{})
	if err != nil {
		logger.Info("create service failed,", zap.String("err", err.Error()))
		return nil, err
	}
	return create, nil
}

func DeleteService(c *gin.Context) error {
	logger := tools.InitLogger()
	clientSet, err := client.GetK8sClientset()
	var svc resource.Service
	c.ShouldBindJSON(&svc)
	err = clientSet.CoreV1().Services(svc.Namespace).Delete(context.TODO(), svc.Name, v12.DeleteOptions{})
	if err != nil {
		logger.Info("delete service failed",zap.String("err",err.Error()))
		return err
	}
	return nil
}

func UpdateService(c *gin.Context) (*v1.Service,error) {
	logger := tools.InitLogger()
	var service resource.Service
	c.ShouldBindJSON(&service)
	clientSet, err := client.GetK8sClientset()
	svc, err := clientSet.CoreV1().Services(service.Namespace).Get(context.TODO(), service.Name, v12.GetOptions{})
	if err != nil {
		logger.Info("get svc failed",zap.String("err",err.Error()))
		return nil,err
	}
	svc.Labels = service.Selector
	newSvc, err := clientSet.CoreV1().Services(service.Name).Update(context.TODO(), svc, v12.UpdateOptions{})
	if err != nil {
		logger.Info("update svc failed",zap.String("err",err.Error()))
		return nil,err
	}
	return newSvc,nil
}