package ingress

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新ingress")
	returnData := config.NewReturnData() //初始化返回数据
	var ingress networkingv1.Ingress     //初始化ingress
	clientset, basicInfo, err := controllers.Basicinit(r, &ingress)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.NetworkingV1().Ingresses(basicInfo.Namespace).Update(context.TODO(), &ingress, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新ingress失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新ingress成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
