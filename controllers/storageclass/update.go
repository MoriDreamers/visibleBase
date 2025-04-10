package storageclass

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新storageclass ")
	returnData := config.NewReturnData()    //初始化返回数据
	var storageclass storagev1.StorageClass //初始化storageclass
	clientset, _, err := controllers.Basicinit(r, &storageclass)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.StorageV1().StorageClasses().Update(context.TODO(), &storageclass, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新storageclass 失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新storageclass 成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
