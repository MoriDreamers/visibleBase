package namespace

import (
	"context"
	"fmt"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Create(r *gin.Context) {
	logs.Info(nil, "添加namespace")

	basicInfo := controllers.Basicinfo{} //初始化基础信息
	returnData := config.NewReturnData() //初始化返回数据

	if err := r.ShouldBindJSON(&basicInfo); err != nil {
		msg := "出错啦！请联系管理员" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	// 前端传一个id名 通过这个id名在全局变量中获取kubeconfig
	kubeconfig := config.CluserKubeConfig[basicInfo.CluserId]
	fmt.Println(config.CluserKubeConfig)
	//转换一下kubeconfig 因为这个是string类型的 所以需要转换一下
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		msg := "获取kubeconfig失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//以此kubeconfig创建一个clientgo客户端工具，这是因为我们是多集群的，所以需要根据不同的集群创建不同的clientgo客户端工具
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "建立clientgo客户端工具失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	var namespace corev1.Namespace
	//clientset的Create会自动读取namespace.Name的值 然后创建一个namespace
	namespace.Name = basicInfo.Name
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		msg := "创建namespace失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建namespace成功"
	r.JSON(200, returnData)
}
