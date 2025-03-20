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
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("获取 Pod 失败: %v", err)
	}

	// 打印 Pod 名称
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	fmt.Printf("有%d个pod\n", len(pods.Items))
}
