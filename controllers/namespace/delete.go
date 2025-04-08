package namespace

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(r *gin.Context) {
	logs.Info(nil, "删除namespace")

	returnData := config.NewReturnData() //初始化返回数据
	clientset, basicInfo, err := controllers.Basicinit(r, nil)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	} //初始化
	protectedNamespace := []string{"kube-system", "kube-public", "kube-node-lease"} //保护的namespace
	for _, namespace := range protectedNamespace {
		if basicInfo.Name == namespace {
			msg := "保护的namespace不能删除"
			returnData.Status = 401
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}
	}

	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.Name, metav1.DeleteOptions{})
	if err != nil {
		msg := "删除namespace失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//返回数据
	returnData.Status = 200
	returnData.Message = "删除namespace成功"
	r.JSON(200, returnData)
}
