package cluster

import (
	"context"
	"visibleBase/config"
	"visibleBase/utils"
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
	//检测集群状态
	clusterStatus, err := clusterConfig.checkClusterStatus()
	if err != nil {
		msg := "集群状态检查失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		logs.Error(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id, "ERR:": err.Error()}, "检查集群信息失败")
		return

	}

	//保存集群相关信息到k8s中，详见思维图中*用于存储的工具*
	var clusterConfigSecret corev1.Secret
	clusterConfigSecret.Name = clusterConfig.Id
	//默认是空的 所以实例化一下
	clusterConfigSecret.Labels = make(map[string]string)
	clusterConfigSecret.Labels["k8s.moridreamers.com/cluster.metadata"] = "true"
	// //默认是空的 所以实例化一下 这里可以优化一下结构 闲了再说 先用着
	// clusterConfigSecret.Annotations["city"] = clusterConfig.City
	// clusterConfigSecret.Annotations["District"] = clusterConfig.District
	// clusterConfigSecret.Annotations["displayName"] = clusterConfig.DisplayName
	clusterConfigSecret.Annotations = make(map[string]string)
	m := utils.StructToMap(clusterStatus)
	clusterConfigSecret.Annotations = m
	//保存kuconfig文件内容
	clusterConfigSecret.StringData = make(map[string]string)
	clusterConfigSecret.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	//正式创建
	_, err1 := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
	if err1 != nil {
		msg := "添加集群失败" + err1.Error()
		logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "msg=": err1.Error()}, "添加集群失败")
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
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
	//接受参数
	clusterId := r.Query("clusterId")
	returnData := config.NewReturnData()
	//删除
	err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Delete(context.TODO(), clusterId, metav1.DeleteOptions{})
	if err != nil {
		msg := "删除集群失败" + err.Error()

		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		logs.Error(map[string]interface{}{"集群ID": clusterId, "msg=": err.Error()}, "删除集群失败")
		return
	}
	logs.Error(map[string]interface{}{"集群ID": clusterId}, "删除集群成功")
	returnData.Status = 200
	returnData.Message = "删除集群成功"
	r.JSON(200, returnData)
	return
}
func Get(r *gin.Context) {
	logs.Info(nil, "获取集群")

}
func List(r *gin.Context) {
	logs.Info(nil, "获取集群列表")
	//根据之前打的标签进行一下筛选 避免把其他东西也返回进来
	listOptions := metav1.ListOptions{
		LabelSelector: "k8s.moridreamers.com/cluster.metadata=true",
	}
	returnData := config.NewReturnData()
	newlist, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).List(context.TODO(), listOptions)
	if err != nil {
		//拉取列表失败
		msg := "拉取列表失败" + err.Error()
		logs.Error(map[string]interface{}{"msg:": err.Error()}, "获取集群列表失败")
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	logs.Error(map[string]interface{}{}, "获取集群列表成功")
	returnData.Data = make(map[string]interface{})
	returnData.Status = 200
	returnData.Message = "获取集群列表成功"
	returnData.Data["items"] = newlist.Items
	r.JSON(200, returnData)
	return
}
