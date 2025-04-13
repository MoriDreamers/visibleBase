package node

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
	logs.Info(nil, "更新node")
	returnData := config.NewReturnData() //初始化返回数据
	var node corev1.Node                 //初始化node
	clientset, _, err := controllers.Basicinit(r, &node)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().Nodes().Update(context.TODO(), &node, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新node失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新node成功"
	returnData.Status = 200
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
