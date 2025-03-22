package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ClusterInfo struct {
	Id          string `json:"id"`          //集群ID
	DisplayName string `json:"displayName"` //别名
	City        string `json:"city"`        //城市
	District    string `json:"district"`    //区域
}

type ClusterStatus struct {
	ClusterInfo
	ClusterStatus  string `json:"clusterStatus"`  //集群状态
	ClusterVersion string `json:"clusterVersion"` //集群版本
}

type ClusterConfig struct {
	ClusterInfo        //继承一些基础信息，在一些场景下不需要发送kubeconfig
	Kubeconfig  string `json:"kubeconfig"` //kubeconfig文件内容
}

// 检查集群状态 方法耦合在上面的结构体中
func (c *ClusterConfig) checkClusterStatus() (ClusterStatus, error) {
	clusterstatus := ClusterStatus{}
	clusterstatus.ClusterInfo = c.ClusterInfo
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		return clusterstatus, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return clusterstatus, err
	}
	severVersion, err := clientset.ServerVersion()
	if err != nil {
		return clusterstatus, err
	}
	clusterVersion := severVersion.String()
	clusterstatus.ClusterStatus = "Active"
	clusterstatus.ClusterVersion = clusterVersion
	return clusterstatus, nil
}
