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

func Create(r *gin.Context) {
	logs.Info(nil, "创建storageclass ")
	var storageclass storagev1.StorageClass
	returnData := config.NewReturnData()
	clientset, _, err := controllers.Basicinit(r, &storageclass)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.StorageV1().StorageClasses().Create(context.TODO(), &storageclass, metav1.CreateOptions{})

	if err != nil {
		msg := "创建storageclass 失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建storageclass 成功"
	r.JSON(200, returnData)
}
