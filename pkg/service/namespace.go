package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"opspaas/pkg/client"
)

type Namespaces struct {
	Name string `json:"name"`
}

func GetNamespace() ([]v12.Namespace,error) {
	clientset, err := client.GetK8sClientset()
	if err != nil {
		log.Fatalln("can not get clientset")
		return nil,err
	}
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalln("can not get namespacelist")
		return nil, err
	}
	return namespaceList.Items,nil
}

func CreateNamespace(c *gin.Context) (name string,err error) {
	var ns Namespaces
	err = c.ShouldBindJSON(&ns)
	fmt.Println("----->",ns)
	if err != nil {
		log.Println("params is wrong!")
		return "", err
	}
	clientset, err := client.GetK8sClientset()
	if err != nil {
		log.Fatalln("can not get clientset")
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
	result, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespaces, v1.CreateOptions{})
	if err != nil {
		log.Fatalln("create namespace error")
		return "", err
	}
	name = result.Name
	return name,nil
}

func DeleteNamespace(c *gin.Context) error {
	var ns Namespaces
	err := c.ShouldBindJSON(&ns)
	if err != nil {
		log.Println("params is wrong!")
		return err
	}
	clientset, err := client.GetK8sClientset()
	if err != nil {
		log.Fatalln("can not get clientset")
		return err
	}
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), ns.Name, v1.DeleteOptions{})
	if err != nil {
		log.Fatalln("delete namespace failed")
		return err
	}
	return nil
}