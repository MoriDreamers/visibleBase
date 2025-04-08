package namespace

import (
	"context"
	"visibleBase/config"
	"visibleBase/controllers"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新namespace")
	returnData := config.NewReturnData() //初始化返回数据
	var ns corev1.Namespace              //初始化namespace
	clientset, _, err := controllers.Basicinit(r, &ns)
	if err != nil {
		msg := err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}

	/*
		解析前端传来的数据 并反序列化为clientset要求的namespace类型的结构体
		因为此逻辑太复杂，考虑到扩展性，已弃用
		// 将 item 转换为 JSON 再反序列化为 Namespace
		itemJSON, err := json.Marshal(basicInfo.Item)
		if err != nil {
			msg := "转换数据失败: " + err.Error()
			returnData.Status = 400
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}

		var namespace corev1.Namespace

		if err1 := json.Unmarshal(itemJSON, &namespace); err1 != nil {
			msg := "解析 Namespace 数据失败: " + err1.Error()
			returnData.Status = 400
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}
		_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), &namespace, metav1.UpdateOptions{})
		if err != nil {
			msg := "更新namespace失败" + err.Error()
			returnData.Status = 401
			returnData.Message = msg
			r.JSON(200, returnData)
			return
		}
	*/
	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新namespace失败" + err.Error()
		returnData.Status = 401
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	msg := "更新namespace成功"
	returnData.Status = 401
	returnData.Message = msg
	r.JSON(200, returnData)
	return
}
