package configmap

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
	logs.Info(nil, "更新configmap")
	returnData := config.NewReturnData() //初始化返回数据
	var configmap corev1.ConfigMap       //初始化configmap
	clientset, basicInfo, err := controllers.Basicinit(r, &configmap)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().ConfigMaps(basicInfo.Namespace).Update(context.TODO(), &configmap, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新configmap失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新configmap成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
