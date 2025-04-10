// 路由信息管理
package routers

import (
	"visibleBase/routers/auth"
	clusterRouter "visibleBase/routers/cluster"
	configmapRouter "visibleBase/routers/configmap"
	cronjobRouter "visibleBase/routers/cronjob"
	deamomsetRouter "visibleBase/routers/deamonset"
	deploymentRouter "visibleBase/routers/deployment"
	ingressRouter "visibleBase/routers/ingress"
	namespaceRouter "visibleBase/routers/namespace"
	nodeRouter "visibleBase/routers/node"
	podRouter "visibleBase/routers/pod"
	replicasetRouter "visibleBase/routers/replicaset"
	secretRouter "visibleBase/routers/secret"
	serviceRouter "visibleBase/routers/service"
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
	replicasetRouter.RegisterSubRouter(apiGroup)
	nodeRouter.RegisterSubRouter(apiGroup)
	serviceRouter.RegisterSubRouter(apiGroup)
	ingressRouter.RegisterSubRouter(apiGroup)
	configmapRouter.RegisterSubRouter(apiGroup)
	secretRouter.RegisterSubRouter(apiGroup)
}
