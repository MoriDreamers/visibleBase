package namespace

import (
	"github.com/gin-gonic/gin"
)

func UpdateAndAdd(r *gin.Context, method string) {
	/*
		var msgValue string
		if method == "update" {
			msgValue = "更新"
		} else {
			msgValue = "添加"
		}

		//首先接收参数，绑定到namespaceConfig结构体中，接着使用内嵌方法检测是否可用，如果可用那么返回一个clusteStatus其中包含anntions的必要字段，
		//将其转成json格式放置到namespaceConfigSecrt中，然后通过slientgo客户端工具更新namespace中的secret
		logs.Info(nil, msgValue+"namespace")
		namespaceConfig := namespaceConfig{}
		returnData := config.NewReturnData() //初始化返回数据

		if err := r.ShouldBindJSON(&namespaceConfig); err != nil {
			msg := "namespace数据绑定失败，请检查输入的数据是否完整" + err.Error()
			returnData.Status = 401
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}
		logs.Info(map[string]interface{}{"namespace名称": namespaceConfig.DisplayName, "namespaceID": namespaceConfig.Id}, "执行namespace"+msgValue+"操作")
		//检测namespace状态
		namespaceStatus, err := namespaceConfig.checknamespaceStatus()
		if err != nil {
			msg := "namespace状态检查失败" + err.Error()
			returnData.Status = 401
			returnData.Message = msg
			r.JSON(200, returnData)
			logs.Error(map[string]interface{}{"namespace名称": namespaceConfig.DisplayName, "namespaceID": namespaceConfig.Id, "ERR:": err.Error()}, "检查namespace信息失败")
			return

		}

		//保存namespace相关信息到k8s中，详见思维图中*用于存储的工具*
		var namespaceConfigSecret corev1.Secret
		namespaceConfigSecret.Name = namespaceConfig.Id
		//默认是空的 所以实例化一下
		namespaceConfigSecret.Labels = make(map[string]string)
		namespaceConfigSecret.Labels["k8s.moridreamers.com/namespace.metadata"] = "true"
		// //默认是空的 所以实例化一下 这里可以优化一下结构 闲了再说 先用着
		// namespaceConfigSecret.Annotations["city"] = namespaceConfig.City
		// namespaceConfigSecret.Annotations["District"] = namespaceConfig.District
		// namespaceConfigSecret.Annotations["displayName"] = namespaceConfig.DisplayName
		namespaceConfigSecret.Annotations = make(map[string]string)
		m := utils.StructToMap(namespaceStatus)
		namespaceConfigSecret.Annotations = m
		//保存kuconfig文件内容
		namespaceConfigSecret.StringData = make(map[string]string)
		namespaceConfigSecret.StringData["kubeconfig"] = namespaceConfig.Kubeconfig
		//正式创建
		if method == "add" {
			_, err1 := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Create(context.TODO(), &namespaceConfigSecret, metav1.CreateOptions{})
			if err1 != nil {
				msg := msgValue + "namespace失败" + err1.Error()
				logs.Error(map[string]interface{}{"namespaceID": namespaceConfig.Id, "msg=": err1.Error()}, msgValue+"namespace失败")
				returnData.Status = 401
				returnData.Message = msg
				r.JSON(200, returnData)
				return
			}
		} else {
			_, err1 := config.InClusterClinetSet.CoreV1().Secrets(config.MetaDataNameSpace).Update(context.TODO(), &namespaceConfigSecret, metav1.UpdateOptions{})
			if err1 != nil {
				msg := msgValue + "namespace失败" + err1.Error()
				logs.Error(map[string]interface{}{"namespaceID": namespaceConfig.Id, "msg=": err1.Error()}, msgValue+"namespace失败")
				returnData.Status = 401
				returnData.Message = msg
				r.JSON(200, returnData)
				return
			}
		}
		logs.Info(map[string]interface{}{"namespace名称": namespaceConfig.DisplayName, "namespaceID": namespaceConfig.Id}, "namespace"+msgValue+"成功")
		returnData.Status = 200
		returnData.Message = "namespace" + msgValue + "成功"
		r.JSON(200, returnData)
		return

	*/
}
