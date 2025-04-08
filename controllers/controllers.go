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
	CluserId  string      `json:"clusterId" ` //form : "clusterId" 这个是用来接收前端传过来的get参数的
	Namespace string      `json:"namespace" `
	Name      string      `json:"name" `
	Item      interface{} `json:"item"` // 用于存储一些配置文件
}

// 这个函数是用来初始化clientset的，因为我们是多集群的，所以需要根据不同的集群创建不同的clientgo客户端工具
func Basicinit(r *gin.Context, item interface{}) (clientset *kubernetes.Clientset, basicInfo Basicinfo, err error) {
	basicInfo = Basicinfo{} //初始化基础信息
	basicInfo.Item = item   //初始化配置文件
	if r.Request.Method == "GET" {
		//如果是GET请求，就从查询参数中获取基本信息，否则就从请求体中获取基本信息
		basicInfo.CluserId = r.Query("clusterId")
		basicInfo.Namespace = r.Query("namespace")
		basicInfo.Name = r.Query("name")
	} else {
		//如果是POST请求，就从请求体中获取基本信息，否则就从查询参数中获取基本信息
		err = r.ShouldBindJSON(&basicInfo)
	}
	if err != nil {
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
