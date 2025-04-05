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

func UpdateAndAdd(r *gin.Context, method string) {
	var msgValue string
	if method == "update" {
		msgValue = "更新"
	} else {
		msgValue = "添加"
	}

	//首先接收参数，绑定到clusterConfig结构体中，接着使用内嵌方法检测是否可用，如果可用那么返回一个clusteStatus其中包含anntions的必要字段，
	//将其转成json格式放置到clusterConfigSecrt中，然后通过slientgo客户端工具更新集群中的secret
	logs.Info(nil, msgValue+"集群")
	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnData() //初始化返回数据

	if err := r.ShouldBindJSON(&clusterConfig); err != nil {
		msg := "集群数据绑定失败，请检查输入的数据是否完整" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	logs.Info(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id}, "执行集群"+msgValue+"操作")
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
	if method == "add" {
		_, err1 := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
		if err1 != nil {
			msg := msgValue + "集群失败" + err1.Error()
			logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "msg=": err1.Error()}, msgValue+"集群失败")
			returnData.Status = 401
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}
	} else {
		_, err1 := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Update(context.TODO(), &clusterConfigSecret, metav1.UpdateOptions{})
		if err1 != nil {
			msg := msgValue + "集群失败" + err1.Error()
			logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "msg=": err1.Error()}, msgValue+"集群失败")
			returnData.Status = 401
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}
	}
	config.CluserKubeConfigPath[clusterConfig.Id] = clusterConfig.Kubeconfig
	logs.Info(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id}, "集群"+msgValue+"成功")
	returnData.Status = 200
	returnData.Message = "集群" + msgValue + "成功"
	r.JSON(200, returnData)
	return
}
