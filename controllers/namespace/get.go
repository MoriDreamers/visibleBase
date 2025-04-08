package namespace

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Info(nil, "获取namespace列表")
	returnData := config.NewReturnData()
	returnData.Data = make(map[string]interface{})
	clientset, basicInfo, err := controllers.Basicinit(r) //初始化
	//获取列表
	var namespace corev1.Namespace
	namespace.Name = basicInfo.Name
	namespaceInfo, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace.Name, metav1.GetOptions{})
	if err != nil {
		msg := "获取namespace详情失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
	} else {
		returnData.Status = 200
		returnData.Message = "获取namespace详情成功"
		returnData.Data["item"] = namespaceInfo
		r.JSON(200, returnData)
	}
}
