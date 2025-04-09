package replicaset

import (
	"visibleBase/controllers/replicaset"

	"github.com/gin-gonic/gin"
)

func get(replicasetGroup *gin.RouterGroup) {
	replicasetGroup.GET("/get", replicaset.Get)
}

func list(replicasetGroup *gin.RouterGroup) {
	replicasetGroup.GET("/list", replicaset.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	replicasetGroup := g.Group("/replicaset")
	get(replicasetGroup)
	list(replicasetGroup)
}
