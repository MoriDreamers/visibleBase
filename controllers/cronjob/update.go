package cronjob

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	batch "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新cronjob")
	returnData := config.NewReturnData() //初始化返回数据
	var cronjob batch.CronJob            //初始化cronjob
	clientset, basicInfo, err := controllers.Basicinit(r, &cronjob)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.BatchV1().CronJobs(basicInfo.Namespace).Update(context.TODO(), &cronjob, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新cronjob失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新cronjob成功"
	returnData.Status = 200
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
