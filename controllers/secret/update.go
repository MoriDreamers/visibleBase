package secret

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
	logs.Info(nil, "更新secret")
	returnData := config.NewReturnData() //初始化返回数据
	var secret corev1.Secret             //初始化secret
	clientset, basicInfo, err := controllers.Basicinit(r, &secret)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().Secrets(basicInfo.Namespace).Update(context.TODO(), &secret, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新secret失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新secret成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
