package controllers

import (
	"errors"
	"visibleBase/config"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//大伙都能经常用到的玩意都在这

type Basicinfo struct {
	CluserId  string      `json:"clusterId"`
	Namespace string      `json:"namespace"`
	Name      string      `json:"name"`
	Item      interface{} `json:"item"` // 用于存储一些配置文件
}

// 这个函数是用来初始化clientset的，因为我们是多集群的，所以需要根据不同的集群创建不同的clientgo客户端工具
func Basicinint(r *gin.Context) (clientset *kubernetes.Clientset, basicInfo Basicinfo, err error) {
	basicInfo = Basicinfo{} //初始化基础信息

	if err = r.ShouldBindJSON(&basicInfo); err != nil {
		msg := "出错啦！请联系管理员" + err.Error()

		return nil, basicInfo, errors.New(msg) // 返回错误和基本信息，以便在调用方处理错误和基本信息
	}
	// 前端传一个id名 通过这个id名在全局变量中获取kubeconfig
	kubeconfig := config.CluserKubeConfig[basicInfo.CluserId]
	//转换一下kubeconfig 因为这个是string类型的 所以需要转换一下
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		msg := "获取kubeconfig失败" + err.Error()
		return nil, basicInfo, errors.New(msg)

	}
	//以此kubeconfig创建一个clientgo客户端工具，这是因为我们是多集群的，所以需要根据不同的集群创建不同的clientgo客户端工具
	clientset, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "建立clientgo客户端工具失败" + err.Error()
		return nil, basicInfo, errors.New(msg)

	}
	return clientset, basicInfo, nil
}
