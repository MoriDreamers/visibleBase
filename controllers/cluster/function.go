package cluster

import (
	"context"
	"sync"
	"time"
	"visibleBase/config"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Add(r *gin.Context) {
	UpdateAndAdd(r, "add")
}

func Update(r *gin.Context) {
	UpdateAndAdd(r, "update")
}

// 检查集群是否可用
func checkClusterHealth(kubeconfigData []byte) bool {
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfigData)
	if err != nil {
		return false
	}

	//设置 Timeout
	restConfig.Timeout = 500 * time.Millisecond

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return false
	}

	//serverVersion() 方法用于检查集群是否可用
	_, err = clientset.Discovery().ServerVersion()
	if err != nil {
		return false
	}

	return true
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
	logs.Info(map[string]interface{}{"集群ID": clusterId}, "删除集群成功")
	returnData.Status = 200
	returnData.Message = "删除集群成功"
	delete(config.CluserKubeConfig, clusterId)
	//调试用 fmt.Println("deldeteTest", config.CluserKubeConfigPath)
	r.JSON(200, returnData)
	return
}
func Get(r *gin.Context) {
	logs.Info(nil, "获取集群配置信息")
	cluserId := r.Query("clusterId")
	returnData := config.NewReturnData()
	ClusterSecret, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Get(context.TODO(), cluserId, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"集群ID": cluserId, "msg=": err.Error()}, "获取集群配置信息失败")
		returnData.Status = 401
		returnData.Message = "获取集群配置信息失败"
	} else {
		logs.Error(map[string]interface{}{"集群ID": cluserId}, "获取集群配置信息成功")
		returnData.Status = 200
		returnData.Message = "获取集群配置信息成功"
		returnData.Data = make(map[string]interface{})
		clusterInfoDetail := ClusterSecret.Annotations
		clusterInfoDetail["kubeconfig"] = string(ClusterSecret.Data["kubeconfig"])
		returnData.Data["item"] = clusterInfoDetail
	}
	r.JSON(200, returnData)
	return
}
func List(r *gin.Context) {
	logs.Info(nil, "获取集群列表")

	listOptions := metav1.ListOptions{
		LabelSelector: "k8s.moridreamers.com/cluster.metadata=true",
	}

	returnData := config.NewReturnData()
	newlist, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).List(context.TODO(), listOptions)
	if err != nil {
		msg := "拉取列表失败" + err.Error()
		logs.Error(map[string]interface{}{"msg:": err.Error()}, "获取集群列表失败")
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	returnData.Data = make(map[string]interface{})
	returnData.Status = 200
	returnData.Message = "获取集群列表成功"

	var clusterList []map[string]string
	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, item := range newlist.Items {
		wg.Add(1)
		itemCopy := item // 防止闭包问题

		go func() {
			defer wg.Done()

			annos := itemCopy.Annotations
			kubeconfigData := itemCopy.Data["kubeconfig"]

			// 探活
			if !checkClusterHealth(kubeconfigData) {
				annos["clusterStatus"] = "inactive"
			}

			// 加锁写入
			lock.Lock()
			clusterList = append(clusterList, annos)
			lock.Unlock()
		}()
	}

	wg.Wait()

	returnData.Data["items"] = clusterList
	logs.Info(nil, "获取集群列表成功")
	r.JSON(200, returnData)
}
