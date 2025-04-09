package daemonset

import (
	"visibleBase/controllers/daemonset"

	"github.com/gin-gonic/gin"
)

func add(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/add", daemonset.Create)
}

func delete(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/delete", daemonset.Delete)
}

func update(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/update", daemonset.Update)
}

func get(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.GET("/get", daemonset.Get)
}

func list(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.GET("/list", daemonset.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	daemonsetGroup := g.Group("/daemonset")
	add(daemonsetGroup)
	delete(daemonsetGroup)
	update(daemonsetGroup)
	get(daemonsetGroup)
	list(daemonsetGroup)
}
