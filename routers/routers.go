// 路由信息管理
package routers

import (
	"visibleBase/routers/auth"
	clusterRouter "visibleBase/routers/cluster"
	cronjobRouter "visibleBase/routers/cronjob"
	deamomsetRouter "visibleBase/routers/deamonset"
	deploymentRouter "visibleBase/routers/deployment"
	namespaceRouter "visibleBase/routers/namespace"
	podRouter "visibleBase/routers/pod"
	statefulsetRouter "visibleBase/routers/statefulset"

	"github.com/gin-gonic/gin"
)

// 路由注册
func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
	clusterRouter.RegisterSubRouter(apiGroup)
	namespaceRouter.RegisterSubRouter(apiGroup)
	podRouter.RegisterSubRouter(apiGroup)
	deploymentRouter.RegisterSubRouter(apiGroup)
	statefulsetRouter.RegisterSubRouter(apiGroup)
	deamomsetRouter.RegisterSubRouter(apiGroup)
	cronjobRouter.RegisterSubRouter(apiGroup)
}
