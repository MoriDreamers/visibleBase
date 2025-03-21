package main

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"encoding/json"
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
	// //查询deployment列表
	// deployments, _ := clientset.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
	// deploymentsItems := deployments.Items
	// for _, deploy := range deploymentsItems {
	// 	fmt.Printf("name:%s,namespace:%s\n", deploy.Name, deploy.Namespace)
	// }
	// //GET可以查询单个
	// detail, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx-5869d7778c-hq49g", metav1.GetOptions{})
	// if err != nil {
	// 	fmt.Println("查询失败!")
	// } else {
	// 	fmt.Printf("镜像详情：%s\n", detail.Spec.Containers[0].Image)
	// }
	// nsDetail, _ := clientset.CoreV1().Namespaces().Get(context.TODO(), "kube-system", metav1.GetOptions{})
	// fmt.Printf("ns详情：%s\n", nsDetail)

	// //获取deployment,并修改
	// deployDetail, _ := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	// fmt.Println("查询deployment名字：", deployDetail.Name)
	// //修改本地获取到的deployment对象的变量值
	// //label测试 修改的是引用类型
	// labels := deployDetail.Labels
	// labels["testlable"] = "testvalue"
	// //Ps: 直接更改，如果操作的条目不存在会报空指针错误
	// //修改副本数测试,它接受一个指针类型的变量
	// repValue := int32(3)
	// deployDetail.Spec.Replicas = &repValue
	// deployDetail.Spec.Template.Spec.Containers[0].Image = "nginx:1.20.1"
	// //将此变量作为参数传入update方法进行修改
	// _, err2 := clientset.AppsV1().Deployments("default").Update(context.TODO(), deployDetail, metav1.UpdateOptions{})
	// if err2 != nil {
	// 	fmt.Println("修改失败!", err2.Error())
	// } else {
	// 	fmt.Println("修改后的label：", deployDetail.Labels)
	// 	fmt.Println("修改后的replicas：", *deployDetail.Spec.Replicas)
	// }

	// //删除资源
	// err3 := clientset.CoreV1().Pods("default").Delete(context.TODO(), "nginx-5869d7778c-7r6j8", metav1.DeleteOptions{})
	// if err3 != nil {
	// 	fmt.Println("删除POD失败!", err3.Error())
	// } else {
	// 	fmt.Println("删除POD成功!")
	// }
	// err4 := clientset.AppsV1().Deployments("default").Delete(context.TODO(), "nginx", metav1.DeleteOptions{})
	// if err4 != nil {
	// 	fmt.Println("删除DEPLOYMENT失败!", err4.Error())
	// } else {
	// 	fmt.Println("删除DEPLOYMENT成功!")
	// }

	//手动填入数值创建资源测试
	//创建ns测试
	var newnameSpace corev1.Namespace
	newnameSpace.Name = "test"
	//传递的参数是一个指针
	_, err5 := clientset.CoreV1().Namespaces().Create(context.TODO(), &newnameSpace, metav1.CreateOptions{})
	if err5 != nil {
		fmt.Println("创建ns失败", err5.Error())
	} else {
		fmt.Println("创建ns成功:", newnameSpace.Name)
	}
	//创建deployment测试
	//在集群里看一下需要的参数：（--dry-run=client -oyaml 不持久化到集群里）kubectl create deployment nginx --image=nginx --dry-run=client -oyaml
	/*
		apiVersion: apps/v1
		kind: Deployment
		metadata:
		creationTimestamp: null
		labels:
			app: nginx
		name: nginx
		spec:
		replicas: 1
		selector:
			matchLabels:
			app: nginx
		strategy: {}
		template:
			metadata:
			creationTimestamp: null
			labels:
				app: nginx
			spec:
			containers:
			- image: nginx
				name: nginx
				resources: {}
		status: {}
	*/
	var newDeployment appsv1.Deployment
	newDeployment.Name = "nginx"
	newDeployment.Namespace = "test"
	label := make(map[string]string)
	label["app"] = "nginx"
	label["version"] = "v1"

	//修改的是是deployment的label
	newDeployment.Labels = label
	//这里更改的是selector的label，matchlabels和template的label是一致的，同时前者不存在需要初始化一下来避免空指针错误
	//这里的初始化就是把包里的空的selector对象赋值给newdeployment的selector
	newDeployment.Spec.Selector = &metav1.LabelSelector{}
	newDeployment.Spec.Selector.MatchLabels = label
	//更改的是模板创建的pod的标签 objectmeta可以省略
	newDeployment.Spec.Template.ObjectMeta.Labels = label
	//创建容器,先声明一个容器再赋值给deployment的spec
	var Containers []corev1.Container
	var Container corev1.Container
	Container.Name = "redis"
	Container.Image = "redis"
	Containers = append(Containers, Container)
	Container.Name = "nginx"
	Container.Image = "nginx"
	Containers = append(Containers, Container)
	newDeployment.Spec.Template.Spec.Containers = Containers
	//正式进行创建
	_, err6 := clientset.AppsV1().Deployments("test").Create(context.TODO(), &newDeployment, metav1.CreateOptions{})
	if err6 != nil {
		fmt.Println("创建失败", err6.Error())
	} else {
		fmt.Println("创建deployment成功:", newDeployment.Name)
	}

	//使用json串来创建资源,查看所需要的json串:	kubectl create deploy redis --image=redis --dry-run=client -ojson
	deploy := `{
		"kind": "Deployment",
		"apiVersion": "apps/v1",
		"metadata": {
			"name": "redis",
			"creationTimestamp": null,
			"labels": {
				"app": "redis"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "redis"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "redis"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis",
							"resources": {}
						}
					]
				}
			},
			"strategy": {}
		},
		"status": {}
	}
	`
	var newDeployment2 appsv1.Deployment
	//进行json转换yaml 把[]byte(A)赋值给B,同时B使用指针类型
	err7 := json.Unmarshal([]byte(deploy), &newDeployment2)
	if err7 != nil {
		fmt.Println("json转换失败", err7.Error())
	}
	fmt.Println("转换后的yaml:", newDeployment2)
}
