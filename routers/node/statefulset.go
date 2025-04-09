package node

import (
	"visibleBase/controllers/node"

	"github.com/gin-gonic/gin"
)

func update(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/update", node.Update)
}

func get(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/get", node.Get)
}

func list(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/list", node.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	nodeGroup := g.Group("/node")
	update(nodeGroup)
	get(nodeGroup)
	list(nodeGroup)
}
