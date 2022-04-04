package resource

import v1 "k8s.io/api/networking/v1"

type Ingress struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Host string `json:"host"`
	Path string `json:"path"`
	PathType v1.PathType `json:"path_type"`
	Port int32 `json:"port"`
}
