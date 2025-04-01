package namespace

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type namespaceInfo struct {
	Id          string `json:"id"`          //namespaceID
	DisplayName string `json:"displayName"` //别名
	City        string `json:"city"`        //城市
	District    string `json:"district"`    //区域
}

type namespaceStatus struct {
	namespaceInfo
	NamespaceStatus  string `json:"namespaceStatus"`  //namespace状态
	NamespaceVersion string `json:"namespaceVersion"` //namespace版本
}

type namespaceConfig struct {
	namespaceInfo        //继承一些基础信息，在一些场景下不需要发送kubeconfig
	Kubeconfig    string `json:"kubeconfig"` //kubeconfig文件内容
}

// 检查namespace状态 方法耦合在上面的结构体中
func (c *namespaceConfig) checknamespaceStatus() (namespaceStatus, error) {
	namespacestatus := namespaceStatus{}
	namespacestatus.namespaceInfo = c.namespaceInfo
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		return namespacestatus, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return namespacestatus, err
	}
	severVersion, err := clientset.ServerVersion()
	if err != nil {
		return namespacestatus, err
	}
	namespaceVersion := severVersion.String()
	namespacestatus.NamespaceStatus = "Active"
	namespacestatus.NamespaceVersion = namespaceVersion
	return namespacestatus, nil
}
