package persistentvolumeclaim

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
	logs.Info(nil, "创建persistentvolumeclaim ")
	var persistentvolumeclaim corev1.PersistentVolumeClaim
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r, &persistentvolumeclaim)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().PersistentVolumeClaims(basicInfo.Namespace).Create(context.TODO(), &persistentvolumeclaim, metav1.CreateOptions{})

	if err != nil {
		msg := "创建persistentvolumeclaim 失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建persistentvolumeclaim 成功"
	r.JSON(200, returnData)
}
