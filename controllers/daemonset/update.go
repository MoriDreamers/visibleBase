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

func Update(r *gin.Context) {
	logs.Info(nil, "更新daemonset")
	returnData := config.NewReturnData() //初始化返回数据
	var daemonset appsv1.DaemonSet       //初始化daemonset
	clientset, basicInfo, err := controllers.Basicinit(r, &daemonset)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.AppsV1().DaemonSets(basicInfo.Namespace).Update(context.TODO(), &daemonset, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新daemonset失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新daemonset成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
