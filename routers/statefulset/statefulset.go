package statefulset

import (
	"visibleBase/controllers/statefulset"

	"github.com/gin-gonic/gin"
)

func add(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/add", statefulset.Create)
}

func delete(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/delete", statefulset.Delete)
}

func update(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/update", statefulset.Update)
}

func get(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.GET("/get", statefulset.Get)
}

func list(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.GET("/list", statefulset.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	statefulsetGroup := g.Group("/statefulset")
	add(statefulsetGroup)
	delete(statefulsetGroup)
	update(statefulsetGroup)
	get(statefulsetGroup)
	list(statefulsetGroup)
}
