package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/api/networking/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"opspaas/pkg/client"
	"opspaas/pkg/tools"
	"opspaas/resource"
)

func GetIngress(c *gin.Context) (*v1.IngressList, error) {
	logger := tools.InitLogger()
	var ing resource.Ingress
	clientSet, _ := client.GetK8sClientset()
	c.ShouldBindJSON(&ing)
	ingressList, err := clientSet.NetworkingV1().Ingresses(ing.NameSpace).List(context.TODO(), v12.ListOptions{})
	if err != nil {
		logger.Info("get ingressList failed", zap.String("err", err.Error()))
		return nil, err
	}
	return ingressList, nil
}

func GetAllIngress() (ingList []string, err error) {
	logger := tools.InitLogger()
	clientSet, _ := client.GetK8sClientset()
	namespaces, err := GetNamespace()
	if err != nil {
		logger.Info("get namespaces failed", zap.String("err", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces {
		ingressList, err := clientSet.NetworkingV1().Ingresses(namespace.Name).List(context.TODO(), v12.ListOptions{})
		if err != nil {
			logger.Info("get ingressList failed", zap.String("err", err.Error()))
			return nil, err
		}
		for _, i := range ingressList.Items {
			ingList = append(ingList, i.Name)
		}
	}
	return ingList, nil
}

func CreateIngress(c *gin.Context) (*v1.Ingress, error) {
	logger := tools.InitLogger()
	clientSet, err := client.GetK8sClientset()
	var ingress resource.Ingress
	c.ShouldBindJSON(&ingress)
	ing := &v1.Ingress{
		ObjectMeta: v12.ObjectMeta{
			Name:        ingress.Name,
			Namespace:   ingress.NameSpace,
			Labels:      ingress.Labels,
			Annotations: ingress.Annotations,
		},
		Spec: v1.IngressSpec{
			Rules: []v1.IngressRule{
				{
					Host: ingress.Host,
					IngressRuleValue: v1.IngressRuleValue{
						HTTP: &v1.HTTPIngressRuleValue{
							Paths: []v1.HTTPIngressPath{
								{
									Path:     ingress.Path,
									PathType: &ingress.PathType,
									Backend: v1.IngressBackend{
										Service: &v1.IngressServiceBackend{
											Name: ingress.Name,
											Port: v1.ServiceBackendPort{
												Name:   ingress.Name,
												Number: ingress.Port,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	NewIngress, err := clientSet.NetworkingV1().Ingresses(ingress.NameSpace).Create(context.TODO(), ing, v12.CreateOptions{})
	if err != nil {
		logger.Info("create ingress failed", zap.String("err", err.Error()))
		return nil, err
	}
	return NewIngress, nil
}
