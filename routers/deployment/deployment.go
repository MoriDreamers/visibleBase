package deployment

import (
	"visibleBase/controllers/deployment"

	"github.com/gin-gonic/gin"
)

func add(deployGroup *gin.RouterGroup) {
	deployGroup.POST("/add", deployment.Create)
}

func delete(deployGroup *gin.RouterGroup) {
	deployGroup.POST("/delete", deployment.Delete)
}

func update(deployGroup *gin.RouterGroup) {
	deployGroup.POST("/update", deployment.Update)
}

func get(deployGroup *gin.RouterGroup) {
	deployGroup.GET("/get", deployment.Get)
}

func list(deployGroup *gin.RouterGroup) {
	deployGroup.GET("/list", deployment.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	deployGroup := g.Group("/deployment")
	add(deployGroup)
	delete(deployGroup)
	update(deployGroup)
	get(deployGroup)
	list(deployGroup)
}
