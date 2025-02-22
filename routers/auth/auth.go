package auth

import (
	"JWT-TEST/controllers/auth"

	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}
func RegisterSubRouter(g *gin.RouterGroup) {
	//配置登录功能的路由策略
	authGroup := g.Group("/auth")
	login(authGroup)
	logout(authGroup)
}
