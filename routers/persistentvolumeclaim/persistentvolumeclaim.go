package persistentvolumeclaim

import (
	"visibleBase/controllers/persistentvolumeclaim"

	"github.com/gin-gonic/gin"
)

func add(persistentvolumeclaimGroup *gin.RouterGroup) {
	persistentvolumeclaimGroup.POST("/add", persistentvolumeclaim.Create)
}

func delete(persistentvolumeclaimGroup *gin.RouterGroup) {
	persistentvolumeclaimGroup.POST("/delete", persistentvolumeclaim.Delete)
}

func update(persistentvolumeclaimGroup *gin.RouterGroup) {
	persistentvolumeclaimGroup.POST("/update", persistentvolumeclaim.Update)
}

func get(persistentvolumeclaimGroup *gin.RouterGroup) {
	persistentvolumeclaimGroup.GET("/get", persistentvolumeclaim.Get)
}

func list(persistentvolumeclaimGroup *gin.RouterGroup) {
	persistentvolumeclaimGroup.GET("/list", persistentvolumeclaim.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	persistentvolumeclaimGroup := g.Group("/persistentvolumeclaim")
	add(persistentvolumeclaimGroup)
	delete(persistentvolumeclaimGroup)
	update(persistentvolumeclaimGroup)
	get(persistentvolumeclaimGroup)
	list(persistentvolumeclaimGroup)
}
