package cluster

import (
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Context) {
	logs.Info(nil, "添加集群")
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
