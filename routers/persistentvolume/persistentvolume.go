package persistentvolume

import (
	"visibleBase/controllers/persistentvolume"

	"github.com/gin-gonic/gin"
)

func add(persistentvolumeGroup *gin.RouterGroup) {
	persistentvolumeGroup.POST("/add", persistentvolume.Create)
}

func delete(persistentvolumeGroup *gin.RouterGroup) {
	persistentvolumeGroup.POST("/delete", persistentvolume.Delete)
}

func update(persistentvolumeGroup *gin.RouterGroup) {
	persistentvolumeGroup.POST("/update", persistentvolume.Update)
}

func get(persistentvolumeGroup *gin.RouterGroup) {
	persistentvolumeGroup.GET("/get", persistentvolume.Get)
}

func list(persistentvolumeGroup *gin.RouterGroup) {
	persistentvolumeGroup.GET("/list", persistentvolume.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	persistentvolumeGroup := g.Group("/persistentvolume")
	add(persistentvolumeGroup)
	delete(persistentvolumeGroup)
	update(persistentvolumeGroup)
	get(persistentvolumeGroup)
	list(persistentvolumeGroup)
}
