package deployment

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新deployment")
	returnData := config.NewReturnData() //初始化返回数据
	var deployment appsv1.Deployment     //初始化deployment
	clientset, basicInfo, err := controllers.Basicinit(r, &deployment)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.AppsV1().Deployments(basicInfo.Namespace).Update(context.TODO(), &deployment, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新deployment失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新deployment成功"
	returnData.Status = 200
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
