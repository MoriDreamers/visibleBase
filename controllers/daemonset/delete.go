package daemonset

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(r *gin.Context) {
	logs.Info(nil, "删除daemonset")

	//returnData := config.NewReturnData() //初始化返回数据
	returnData := config.NewReturnData() //初始化返回数据
	clientset, basicInfo, err := controllers.Basicinit(r, nil)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//初始化

	err = clientset.AppsV1().DaemonSets(basicInfo.Namespace).Delete(context.TODO(), basicInfo.Name, metav1.DeleteOptions{})

	if err != nil {
		msg := "daemonset删除失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "删除daemonset成功"
	r.JSON(200, returnData)
}
