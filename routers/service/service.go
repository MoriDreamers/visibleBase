package service

import (
	"visibleBase/controllers/service"

	"github.com/gin-gonic/gin"
)

func add(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/add", service.Create)
}

func delete(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/delete", service.Delete)
}

func update(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/update", service.Update)
}

func get(serviceGroup *gin.RouterGroup) {
	serviceGroup.GET("/get", service.Get)
}

func list(serviceGroup *gin.RouterGroup) {
	serviceGroup.GET("/list", service.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	serviceGroup := g.Group("/service")
	add(serviceGroup)
	delete(serviceGroup)
	update(serviceGroup)
	get(serviceGroup)
	list(serviceGroup)
}
