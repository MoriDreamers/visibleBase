package persistentvolume

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(r *gin.Context) {
	logs.Info(nil, "获取persistentvolume 列表")
	returnData := config.NewReturnData()
	returnData.Data = make(map[string]interface{})
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
	List, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		msg := "获取persistentvolume列表失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
	} else {
		returnData.Status = 200
		returnData.Message = "获取persistentvolume列表成功"
		/*
			这里可以优化一下 因为我们只需要返回名称 所以可以直接返回一个字符串数组 详见persistentvolume 中的函数注释
		*/
		returnData.Data["persistentvolume List"] = List.Items
		r.JSON(200, returnData)
	}
}
