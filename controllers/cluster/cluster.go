package cluster

type ClusterInfo struct {
	Id          string `json:"id"`          //集群ID
	DisplayName string `json:"displayName"` //别名
	City        string `json:"city"`        //城市
	District    string `json:"district"`    //区域
}

type ClusterConfig struct {
	ClusterInfo        //继承一些基础信息，在一些场景下不需要发送kubeconfig
	Kubeconfig  string `json:"kubeconfig"` //kubeconfig文件内容
}
