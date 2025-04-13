package service

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新service")
	returnData := config.NewReturnData() //初始化返回数据
	var service corev1.Service           //初始化service
	clientset, basicInfo, err := controllers.Basicinit(r, &service)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().Services(basicInfo.Namespace).Update(context.TODO(), &service, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新service失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新service成功"
	returnData.Status = 200
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
