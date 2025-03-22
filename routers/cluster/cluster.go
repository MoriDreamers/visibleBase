package cluster

import (
	"visibleBase/controllers/cluster"

	"github.com/gin-gonic/gin"
)

// 先写路由方法，调用中间件，再注册到路由组里
func update(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/update", cluster.Update)
}
func add(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/add", cluster.Add)
}
func delete(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/delete", cluster.Delete)
}

func get(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", cluster.Get)
}
func list(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", cluster.List)
}
func RegisterSubRouter(g *gin.RouterGroup) {
	clusterGroup := g.Group("/cluster")
	add(clusterGroup)
	update(clusterGroup)
	delete(clusterGroup)
	get(clusterGroup)
	list(clusterGroup)
}
