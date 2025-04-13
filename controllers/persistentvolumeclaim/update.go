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

func Update(r *gin.Context) {
	logs.Info(nil, "更新persistentvolumeclaim ")
	returnData := config.NewReturnData()                   //初始化返回数据
	var persistentvolumeclaim corev1.PersistentVolumeClaim //初始化persistentvolumeclaim
	clientset, basicInfo, err := controllers.Basicinit(r, &persistentvolumeclaim)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().PersistentVolumeClaims(basicInfo.Namespace).Update(context.TODO(), &persistentvolumeclaim, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新persistentvolumeclaim 失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新persistentvolumeclaim 成功"
	returnData.Status = 200
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
