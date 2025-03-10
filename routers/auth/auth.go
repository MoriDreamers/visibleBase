package auth

import (
	"JWT-TEST/controllers/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}

// 本地调试用户列表 请删除
func userlist(authGroup *gin.RouterGroup) {
	authGroup.GET("/userlist", func(c *gin.Context) {
		// 返回一些写死的 JSON 数据
		//这里睡眠一秒，模拟网络延迟，我弄加载动画玩
		time.Sleep(1 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"items": []gin.H{
					{"id": 1, "username": "暮羽", "qq": "1111111", "address": "北京市朝阳区"},
					{"id": 2, "username": "美冬", "qq": "2222222", "address": "北京市海淀区"},
					{"id": 3, "username": "初音", "qq": "3333333", "address": "北京市昌平区"},
					{"id": 4, "username": "圆香", "qq": "4444444", "address": "上海市虹桥区"},
					{"id": 5, "username": "真幌", "qq": "5555555", "address": "深圳市福田区"},
				},
			},
		})
	})
}

func RegisterSubRouter(g *gin.RouterGroup) {
	//配置登录功能的路由策略
	authGroup := g.Group("/auth")
	testGroup := g.Group("/test")
	login(authGroup)
	logout(authGroup)
	userlist(testGroup)
}
