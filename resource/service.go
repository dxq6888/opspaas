package resource

import v1 "k8s.io/api/core/v1"

type Service struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Ports []int32 `json:"ports"`
	Selector map[string]string `json:"selector"`
	Type v1.ServiceType `json:"type"`
	Protocol v1.Protocol `json:"protocol"`
}
