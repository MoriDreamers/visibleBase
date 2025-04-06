package namespace

import (
	"visibleBase/controllers/namespace"

	"github.com/gin-gonic/gin"
)

// 先写路由方法，调用中间件，再注册到路由组里
func update(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/update", namespace.Update)
}
func add(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/add", namespace.Create)
}
func delete(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/delete", namespace.Delete)
}

func get(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/get", namespace.Get)
}
func list(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/list", namespace.List)
}
func RegisterSubRouter(g *gin.RouterGroup) {
	namespaceGroup := g.Group("/namespace")
	add(namespaceGroup)
	update(namespaceGroup)
	delete(namespaceGroup)
	get(namespaceGroup)
	list(namespaceGroup)
}
