package namespace

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
	logs.Info(nil, "添加namespace")
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	var namespace corev1.Namespace
	//clientset的Create会自动读取namespace.Name的值 然后创建一个namespace
	namespace.Name = basicInfo.Name
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})

	if err != nil {
		msg := "创建namespace失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建namespace成功"
	r.JSON(200, returnData)
}
