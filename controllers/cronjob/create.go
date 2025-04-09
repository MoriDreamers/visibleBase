package cronjob

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(r *gin.Context) {
	logs.Info(nil, "创建cronjob")
	var cronjob batchv1.CronJob
	returnData := config.NewReturnData()
	clientset, basicInfo, err := controllers.Basicinit(r, &cronjob)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.BatchV1().CronJobs(basicInfo.Namespace).Create(context.TODO(), &cronjob, metav1.CreateOptions{})

	if err != nil {
		msg := "创建cronjob失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "创建cronjob成功"
	r.JSON(200, returnData)
}
