package daemonset

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(r *gin.Context) {
	logs.Info(nil, "创建daemonset")
	var daemonset appsv1.DaemonSet
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r, &daemonset)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.AppsV1().DaemonSets(basicInfo.Namespace).Create(context.TODO(), &daemonset, metav1.CreateOptions{})

	if err != nil {
		msg := "创建daemonset失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建daemonset成功"
	r.JSON(200, returnData)
}
