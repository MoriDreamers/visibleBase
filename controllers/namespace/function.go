package namespace

import (
	"context"
	"visibleBase/config"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Add(r *gin.Context) {
	UpdateAndAdd(r, "add")
}

func Update(r *gin.Context) {
	UpdateAndAdd(r, "update")
}

func Delete(r *gin.Context) {
	logs.Info(nil, "删除namespace")
	//接受参数
	namespaceId := r.Query("namespaceId")
	returnData := config.NewReturnData()
	//删除
	err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Delete(context.TODO(), namespaceId, metav1.DeleteOptions{})
	if err != nil {
		msg := "删除namespace失败" + err.Error()

		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		logs.Error(map[string]interface{}{"namespaceID": namespaceId, "msg=": err.Error()}, "删除namespace失败")
		return
	}
	logs.Error(map[string]interface{}{"namespaceID": namespaceId}, "删除namespace成功")
	returnData.Status = 200
	returnData.Message = "删除namespace成功"
	r.JSON(200, returnData)
	return
}
func Get(r *gin.Context) {
	logs.Info(nil, "获取namespace配置信息")
	cluserId := r.Query("namespaceId")
	returnData := config.NewReturnData()
	namespaceSecret, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Get(context.TODO(), cluserId, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"namespaceID": cluserId, "msg=": err.Error()}, "获取namespace配置信息失败")
		returnData.Status = 401
		returnData.Message = "获取namespace配置信息失败"
	} else {
		logs.Error(map[string]interface{}{"namespaceID": cluserId}, "获取namespace配置信息成功")
		returnData.Status = 200
		returnData.Message = "获取namespace配置信息成功"
		returnData.Data = make(map[string]interface{})
		namespaceInfoDetail := namespaceSecret.Annotations
		namespaceInfoDetail["kubeconfig"] = string(namespaceSecret.Data["kubeconfig"])
		returnData.Data["item"] = namespaceInfoDetail
	}
	r.JSON(200, returnData)
	return
}
func List(r *gin.Context) {
	logs.Info(nil, "获取namespace列表")
	//根据之前打的标签进行一下筛选 避免把其他东西也返回进来
	listOptions := metav1.ListOptions{
		LabelSelector: "k8s.moridreamers.com/namespace.metadata=true",
	}
	returnData := config.NewReturnData()
	newlist, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).List(context.TODO(), listOptions)
	if err != nil {
		//拉取列表失败
		msg := "拉取列表失败" + err.Error()
		logs.Error(map[string]interface{}{"msg:": err.Error()}, "获取namespace列表失败")
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	logs.Error(map[string]interface{}{}, "获取namespace列表成功")
	returnData.Data = make(map[string]interface{})
	returnData.Status = 200
	returnData.Message = "获取namespace列表成功"
	//返回的数据太多，so 这里只返回部分信息
	var namespaceList []map[string]string
	for _, item := range newlist.Items {
		annos := item.Annotations
		namespaceList = append(namespaceList, annos)
	}
	//写BUG了 这样便利出来的不是数组 只保留最后一个的数据 但是可以实现定向查询 所以留着
	// for _, item := range newlist.Items {
	// 	namespaceList["displayName"] = item.Annotations["displayName"]
	// 	namespaceList["city"] = item.Annotations["city"]
	// 	namespaceList["namespaceStatus"] = item.Annotations["namespaceStatus"]
	// 	namespaceList["namespaceVersion"] = item.Annotations["namespaceVersion"]
	// 	namespaceList["district"] = item.Annotations["district"]
	// 	namespaceList["id"] = item.Annotations["id"]
	// }
	returnData.Data["items"] = namespaceList
	r.JSON(200, returnData)
	return
}
