package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 设置 kubeconfig 路径
	kubeconfig := "./kube/config"

	// 加载 kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("加载 kubeconfig 失败: %v", err)
	}

	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("创建 Kubernetes 客户端失败: %v", err)
	}

	// 获取所有 Pod
	clientset.AppsV1()
	clientset.NetworkingV1()
	clientset.StorageV1()
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("获取 Pod 失败: %v", err)
	}

	// 打印 Pod 名称
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	fmt.Printf("有%d个pod\n", len(pods.Items))
	//查询deployment列表
	deployments, _ := clientset.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
	deploymentsItems := deployments.Items
	for _, deploy := range deploymentsItems {
		fmt.Printf("name:%s,namespace:%s\n", deploy.Name, deploy.Namespace)
	}
	//GET可以查询单个
	detail, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx-5869d7778c-hq49g", metav1.GetOptions{})
	if err != nil {
		fmt.Println("查询失败!")
	} else {
		fmt.Printf("镜像详情：%s\n", detail.Spec.Containers[0].Image)
	}
	nsDetail, _ := clientset.CoreV1().Namespaces().Get(context.TODO(), "kube-system", metav1.GetOptions{})
	fmt.Printf("ns详情：%s\n", nsDetail)

	//获取deployment,并修改
	deployDetail, _ := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	fmt.Println("查询deployment名字：", deployDetail.Name)
	//修改本地获取到的deployment对象的变量值
	//label测试 修改的是引用类型
	labels := deployDetail.Labels
	labels["testlable"] = "testvalue"
	//Ps: 直接更改，如果操作的条目不存在会报空指针错误
	//修改副本数测试,它接受一个指针类型的变量
	repValue := int32(3)
	deployDetail.Spec.Replicas = &repValue
	deployDetail.Spec.Template.Spec.Containers[0].Image = "nginx:1.20.1"
	//将此变量作为参数传入update方法进行修改
	_, err2 := clientset.AppsV1().Deployments("default").Update(context.TODO(), deployDetail, metav1.UpdateOptions{})
	if err2 != nil {
		fmt.Println("修改失败!", err2.Error())
	} else {
		fmt.Println("修改后的label：", deployDetail.Labels)
		fmt.Println("修改后的replicas：", *deployDetail.Spec.Replicas)
	}

	//删除资源
	err3 := clientset.CoreV1().Pods("default").Delete(context.TODO(), "nginx-5869d7778c-7r6j8", metav1.DeleteOptions{})
	if err3 != nil {
		fmt.Println("删除POD失败!", err3.Error())
	} else {
		fmt.Println("删除POD成功!")
	}
	err4 := clientset.AppsV1().Deployments("default").Delete(context.TODO(), "nginx", metav1.DeleteOptions{})
	if err4 != nil {
		fmt.Println("删除DEPLOYMENT失败!", err4.Error())
	} else {
		fmt.Println("删除DEPLOYMENT成功!")
	}

}
