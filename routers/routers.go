// 路由信息管理
package routers

import (
	"JWT-TEST/routers/auth"

	"github.com/gin-gonic/gin"
)

// 路由注册
func RegisterRouters(r *gin.Engine) {
	//登陆的路由配置
	//登录 login
	//登出 logout
	//3. /api/auth/login
	//	/api/auth/logout
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
}
