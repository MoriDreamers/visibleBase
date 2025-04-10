package persistentvolume

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
	logs.Info(nil, "更新persistentvolume ")
	returnData := config.NewReturnData()         //初始化返回数据
	var persistentvolume corev1.PersistentVolume //初始化persistentvolume
	clientset, _, err := controllers.Basicinit(r, &persistentvolume)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().PersistentVolumes().Update(context.TODO(), &persistentvolume, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新persistentvolume 失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新persistentvolume 成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
