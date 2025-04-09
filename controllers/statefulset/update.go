package statefulset

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
	logs.Info(nil, "更新statefulset")
	returnData := config.NewReturnData() //初始化返回数据
	var statefulset appsv1.StatefulSet   //初始化statefulset
	clientset, basicInfo, err := controllers.Basicinit(r, &statefulset)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.AppsV1().StatefulSets(basicInfo.Namespace).Update(context.TODO(), &statefulset, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新statefulset失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新statefulset成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
