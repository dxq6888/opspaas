package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func GetK8sClientset() (clientset *kubernetes.Clientset, err error) {
	config, err := GetConfig()
	if err != nil {
		log.Fatalln("can not get config from GetConfig")
		return nil, err
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("can not get clientset")
		return nil, err
	}
	return clientset,nil
}

func GetConfig() (config *rest.Config, err error) {
	config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		inClusterConfig, err := rest.InClusterConfig() //集群内部创建config对象
		if err != nil {
			log.Fatalln("can not get config")
			return nil, err
		}
		config = inClusterConfig
	}

	return config,nil
}