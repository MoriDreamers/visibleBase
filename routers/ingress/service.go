package ingress

import (
	"visibleBase/controllers/ingress"

	"github.com/gin-gonic/gin"
)

func add(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/add", ingress.Create)
}

func delete(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/delete", ingress.Delete)
}

func update(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/update", ingress.Update)
}

func get(ingressGroup *gin.RouterGroup) {
	ingressGroup.GET("/get", ingress.Get)
}

func list(ingressGroup *gin.RouterGroup) {
	ingressGroup.GET("/list", ingress.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	ingressGroup := g.Group("/ingress")
	add(ingressGroup)
	delete(ingressGroup)
	update(ingressGroup)
	get(ingressGroup)
	list(ingressGroup)
}
