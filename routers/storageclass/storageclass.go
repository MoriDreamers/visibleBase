package storageclass

import (
	"visibleBase/controllers/storageclass"

	"github.com/gin-gonic/gin"
)

func add(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.POST("/add", storageclass.Create)
}

func delete(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.POST("/delete", storageclass.Delete)
}

func update(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.POST("/update", storageclass.Update)
}

func get(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.GET("/get", storageclass.Get)
}

func list(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.GET("/list", storageclass.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	storageclassGroup := g.Group("/storageclass")
	add(storageclassGroup)
	delete(storageclassGroup)
	update(storageclassGroup)
	get(storageclassGroup)
	list(storageclassGroup)
}
