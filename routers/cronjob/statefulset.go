package cronjob

import (
	"visibleBase/controllers/cronjob"

	"github.com/gin-gonic/gin"
)

func add(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/add", cronjob.Create)
}

func delete(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/delete", cronjob.Delete)
}

func update(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/update", cronjob.Update)
}

func get(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.GET("/get", cronjob.Get)
}

func list(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.GET("/list", cronjob.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	cronjobGroup := g.Group("/cronjob")
	add(cronjobGroup)
	delete(cronjobGroup)
	update(cronjobGroup)
	get(cronjobGroup)
	list(cronjobGroup)
}
