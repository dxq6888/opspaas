package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"opspaas/pkg/client"
	"opspaas/pkg/config"
	"opspaas/pkg/tools"
	"opspaas/resource"
)

func GetDeployment(c *gin.Context) (deploymentList []v12.Deployment, err error) {
	logger := tools.InitLogger()
	var ns config.Namespaces
	c.ShouldBindJSON(&ns)
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("get clientSet failed", zap.String("err: ", err.Error()))
		return nil, err
	}
	list, err := clientSet.AppsV1().Deployments(ns.Name).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		logger.Info("get deploymentList failed", zap.String("err: ", err.Error()))
		return nil, err
	}
	items := list.Items
	return items, nil
}

func GetAllDeployment(c *gin.Context) (allDeploy []string, err error) {
	logger := tools.InitLogger()
	var ns config.Namespaces
	nameSpaces, err := GetNamespace()
	c.ShouldBindJSON(&ns)
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("get all deployment failed: ", zap.String("err:", err.Error()))
		return nil, err
	}
	for _, nsList := range nameSpaces {
		dpList, err := clientSet.AppsV1().Deployments(nsList.Name).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			logger.Info("get dpList failed", zap.String("err: ", err.Error()))
			return nil, err
		}
		for _, dp := range dpList.Items {
			allDeploy = append(allDeploy, dp.Name)
		}
	}
	return allDeploy, nil
}

func CreateDeployment(c *gin.Context) (dep *v12.Deployment,err error) {
	logger := tools.InitLogger()
	clientSet, err := client.GetK8sClientset()
	var deployment resource.Deployment
	c.ShouldBindJSON(&deployment)
	if err != nil {
		logger.Info("get clientSet failed,", zap.String("err:", err.Error()))
		return
	}

	deploy := &v12.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:        deployment.Name,
			Namespace:   deployment.NameSpace,
			Labels:      deployment.DeploymentLabels,
			Annotations: deployment.Annotations,
		},
		Spec: v12.DeploymentSpec{
			Replicas: deployment.Replicas,
			Selector: &v1.LabelSelector{
				MatchLabels: deployment.Selector,
			},
			Template: v13.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:   deployment.Name,
					Labels: deployment.Selector,
				},
				Spec: v13.PodSpec{
					Containers: []v13.Container{
						{
							Name:  deployment.Name,
							Image: deployment.Image,
							Ports: []v13.ContainerPort{
								{
									Name:          deployment.Name,
									Protocol:      deployment.Protocol,
									ContainerPort: deployment.Ports,
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println(deployment)
	dep, err = clientSet.AppsV1().Deployments(deployment.NameSpace).Create(context.TODO(), deploy, v1.CreateOptions{})
	if err != nil {
		logger.Info("create deployment failed:",zap.String("err",err.Error()))
		return nil,err
	}
	return dep,nil
}

func DeleteDeployment(c *gin.Context) (err error) {
	logger := tools.InitLogger()
	clientSet, err := client.GetK8sClientset()
	var deployment resource.Deployment
	c.ShouldBindJSON(&deployment)
	if err != nil {
		logger.Info("get clientSet failed,",zap.String("err:",err.Error()))
		return err
	}
	err = clientSet.AppsV1().Deployments(deployment.NameSpace).Delete(context.TODO(), deployment.Name, v1.DeleteOptions{})
	if err != nil {
		logger.Info("delete deployment failed,",zap.String("err:",err.Error()))
		return err
	}
	return nil
}

func UpdateDeployment(c *gin.Context) {
	logger := tools.InitLogger()
	var deployment resource.Deployment
	c.ShouldBindJSON(&deployment)
	clientSet, err := client.GetK8sClientset()
	if err != nil {
		logger.Info("get clientSet failed,",zap.String("err",err.Error()))
		return
	}
	clientSet.AppsV1().Deployments(deployment.NameSpace).Update(context.TODO(),)
}