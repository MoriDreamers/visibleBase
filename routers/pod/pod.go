package pod

import (
	"visibleBase/controllers/pod"

	"github.com/gin-gonic/gin"
)

// 先写路由方法，调用中间件，再注册到路由组里
func update(podGroup *gin.RouterGroup) {
	podGroup.POST("/update", pod.Update)
}
func add(podGroup *gin.RouterGroup) {
	podGroup.POST("/add", pod.Create)
}
func delete(podGroup *gin.RouterGroup) {
	podGroup.POST("/delete", pod.Delete)
}

func get(podGroup *gin.RouterGroup) {
	podGroup.GET("/get", pod.Get)
}
func list(podGroup *gin.RouterGroup) {
	podGroup.GET("/list", pod.List)
}
func RegisterSubRouter(g *gin.RouterGroup) {
	podGroup := g.Group("/pod")
	add(podGroup)
	update(podGroup)
	delete(podGroup)
	get(podGroup)
	list(podGroup)
}
