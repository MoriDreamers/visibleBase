package deployment

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Info(nil, "获取deployment列表")
	returnData := config.NewReturnData()
	returnData.Data = make(map[string]interface{})
	clientset, basicInfo, err := controllers.Basicinit(r, nil)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//获取列表
	deploymentInfo, err := clientset.AppsV1().Deployments(basicInfo.Namespace).Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		msg := "获取deployment详情失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
	} else {
		returnData.Status = 200
		returnData.Message = "获取deployment详情成功"
		returnData.Data["item"] = deploymentInfo
		r.JSON(200, returnData)
	}
}
