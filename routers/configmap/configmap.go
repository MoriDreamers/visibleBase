package configmap

import (
	"visibleBase/controllers/configmap"

	"github.com/gin-gonic/gin"
)

func add(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/add", configmap.Create)
}

func delete(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/delete", configmap.Delete)
}

func update(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/update", configmap.Update)
}

func get(configmapGroup *gin.RouterGroup) {
	configmapGroup.GET("/get", configmap.Get)
}

func list(configmapGroup *gin.RouterGroup) {
	configmapGroup.GET("/list", configmap.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	configmapGroup := g.Group("/configmap")
	add(configmapGroup)
	delete(configmapGroup)
	update(configmapGroup)
	get(configmapGroup)
	list(configmapGroup)
}
