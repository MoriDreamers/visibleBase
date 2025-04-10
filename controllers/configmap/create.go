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

func Create(r *gin.Context) {
	logs.Info(nil, "创建configmap")
	var configmap corev1.ConfigMap
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r, &configmap)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().ConfigMaps(basicInfo.Namespace).Create(context.TODO(), &configmap, metav1.CreateOptions{})

	if err != nil {
		msg := "创建configmap失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建configmap成功"
	r.JSON(200, returnData)
}
