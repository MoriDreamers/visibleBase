package cluster

import (
	"context"
	"visibleBase/config"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Add(r *gin.Context) {
	logs.Info(nil, "添加集群")
	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnData() //初始化返回数据

	if err := r.ShouldBindJSON(&clusterConfig); err != nil {
		msg := "集群数据绑定失败，请检查输入的数据是否完整" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	logs.Info(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id}, "执行集群添加操作")
	returnData.Status = 200
	returnData.Message = "集群添加成功"
	r.JSON(200, returnData)
	//保存集群相关信息到k8s中，详见思维图中*用于存储的工具*
	var clusterConfigSecret corev1.Secret
	clusterConfigSecret.Name = clusterConfig.Id
	//默认是空的 所以实例化一下
	clusterConfigSecret.Labels = make(map[string]string)
	clusterConfigSecret.Labels["k8s.moridreamers.com/cluster.metadata"] = "true"
	//默认是空的 所以实例化一下 这里可以优化一下结构 闲了再说 先用着
	clusterConfigSecret.Annotations = make(map[string]string)
	clusterConfigSecret.Annotations["city"] = clusterConfig.City
	clusterConfigSecret.Annotations["District"] = clusterConfig.District
	clusterConfigSecret.Annotations["displayName"] = clusterConfig.DisplayName
	//保存kuconfig文件内容
	clusterConfigSecret.StringData = make(map[string]string)
	clusterConfigSecret.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	//正式创建
	_, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
	if err != nil {
		msg := "添加集群失败" + err.Error()
		logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "msg=": err.Error()}, "添加集群失败")
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
	}
	logs.Info(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id}, "集群添加成功")
	returnData.Status = 200
	returnData.Message = "集群添加成功"
	r.JSON(200, returnData)
	return
}

func Update(r *gin.Context) {
	logs.Info(nil, "更新集群")
	return
}

func Delete(r *gin.Context) {
	logs.Info(nil, "删除集群")
	return
}
func Get(r *gin.Context) {
	logs.Info(nil, "获取集群")
	return
}
func List(r *gin.Context) {
	logs.Info(nil, "获取集群列表")
	return
}
