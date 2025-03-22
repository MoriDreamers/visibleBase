// 路由信息管理
package routers

import (
	"visibleBase/routers/auth"
	clusterRouter "visibleBase/routers/cluster"

	"github.com/gin-gonic/gin"
)

// 路由注册
func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
	clusterRouter.RegisterSubRouter(apiGroup)
}
