package secret

import (
	"visibleBase/controllers/secret"

	"github.com/gin-gonic/gin"
)

func add(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/add", secret.Create)
}

func delete(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/delete", secret.Delete)
}

func update(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/update", secret.Update)
}

func get(secretGroup *gin.RouterGroup) {
	secretGroup.GET("/get", secret.Get)
}

func list(secretGroup *gin.RouterGroup) {
	secretGroup.GET("/list", secret.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	secretGroup := g.Group("/secret")
	add(secretGroup)
	delete(secretGroup)
	update(secretGroup)
	get(secretGroup)
	list(secretGroup)
}
