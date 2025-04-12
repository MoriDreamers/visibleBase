package namespace

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(r *gin.Context) {
	logs.Info(nil, "获取namespace列表")
	returnData := config.NewReturnData()
	returnData.Data = make(map[string]interface{})
	clientset, _, err := controllers.Basicinit(r, nil)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	//获取列表
	namspaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		msg := "获取namespace列表失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
	} else {
		returnData.Status = 200
		returnData.Message = "获取namespace列表成功"
		/*
			这里可以优化一下 因为我们只需要返回namespace的名称 所以可以直接返回一个字符串数组
			var nsNames []string
			for _, namespace := range namspaceList.Items {
				nsNames = append(nsNames, namespace.Name)
			}
			returnData.Data["namespaceList"] = nsNames
		*/
		returnData.Data["items"] = namspaceList.Items
		r.JSON(200, returnData)
	}
}
