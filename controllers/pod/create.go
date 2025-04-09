package pod

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
	logs.Info(nil, "pod")
	var pod corev1.Pod
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r, &pod)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	//clientset的Create会自动读取namespace.Name的值 然后创建一个namespace
	_, err = clientset.CoreV1().Pods(basicInfo.Namespace).Create(context.TODO(), &pod, metav1.CreateOptions{})

	if err != nil {
		msg := "创建pod失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建pod成功"
	r.JSON(200, returnData)
}
