package resource

import v1 "k8s.io/api/core/v1"

type Deployment struct {
	Name             string            `json:"deployment_name"`
	NameSpace        string            `json:"namespace"`
	DeploymentLabels map[string]string `json:"deployment_labels"`
	Annotations      map[string]string `json:"annotations"`
	Replicas         *int32            `json:"replicas"`
	Selector         map[string]string `json:"selector"`
	Image            string            `json:"image"`
	Ports            int32             `json:"ports"`
	Protocol         v1.Protocol       `json:"protocol"`
}

