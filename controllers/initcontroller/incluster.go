package initcontroller

import (
	"context"
	"fmt"
	"log"
	"visibleBase/config"
	"visibleBase/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func metadataInit() {
	logs.Debug(nil, "初始化元数据命名空间")
	// 设置 kubeconfig 路径 因为mian.go调用所以基准是根目录
	kubeconfig := "./kube/config"

	// 加载 kubeconfig
	setconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("加载 kubeconfig 失败: %v", err)
		panic(err.Error())
	}

	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(setconfig)
	if err != nil {
		log.Fatalf("创建 Kubernetes 客户端失败: %v", err)
		panic(err.Error())
	}
	config.InClusterClinetSet = clientset
	inClusterVersion, _ := clientset.Discovery().ServerVersion()
	_, err1 := clientset.CoreV1().Namespaces().Get(context.TODO(), config.MetaDataNameSpace, metav1.GetOptions{})
	if err1 != nil {
		logs.Debug(nil, "元数据命名空间不存在，创建元数据命名空间")
		var createNamespace corev1.Namespace
		createNamespace.Name = config.MetaDataNameSpace
		_, err2 := clientset.CoreV1().Namespaces().Create(context.TODO(), &createNamespace, metav1.CreateOptions{})
		if err2 != nil {
			logs.Error(nil, "创建元数据命名空间失败")
			panic(err2.Error())
		}
		logs.Debug(map[string]interface{}{"namespace": config.MetaDataNameSpace, "clusterVersion": inClusterVersion}, "创建元数据命名空间成功,集群版本已打印")
	} else {
		//已存在namespace
		logs.Info(map[string]interface{}{"clusterVersion": inClusterVersion}, "元数据命名空间已存在,集群版本已打印")
	}
	// 初始化clusterkubeconfig
	config.CluserKubeConfig = make(map[string]string)
	//这是一个解耦的筛选器 用于在下面进行筛选我们的元数据
	listOptions := metav1.ListOptions{
		LabelSelector: "k8s.moridreamers.com/cluster.metadata=true",
	}
	// 获取我们存放kubeconfig的Secret 用于后续的获取kubecofig
	secretsList, err := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).List(context.TODO(), listOptions)
	for _, secret := range secretsList.Items {
		clusterId := secret.Name
		kubeconfig := secret.Data["kubeconfig"]
		config.CluserKubeConfig[clusterId] = string(kubeconfig)

	}
	//调试用 fmt.Print("clusterIDtEST:", config.CluserKubeConfig)
	// 获取所有 Pod
	clientset.AppsV1()
	clientset.NetworkingV1()
	clientset.StorageV1()
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("获取 Pod 失败: %v", err)
	}
	fmt.Printf("有%d个pod\n", len(pods.Items))
}
