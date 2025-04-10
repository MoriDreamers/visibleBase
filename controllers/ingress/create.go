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

func Create(r *gin.Context) {
	logs.Info(nil, "创建ingress")
	var ingress networkingv1.Ingress
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r, &ingress)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.NetworkingV1().Ingresses(basicInfo.Namespace).Create(context.TODO(), &ingress, metav1.CreateOptions{})

	if err != nil {
		msg := "创建ingress失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建ingress成功"
	r.JSON(200, returnData)
}
