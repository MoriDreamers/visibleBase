package middlewares

import (
	"visibleBase/config"
	"visibleBase/utils/jwtutil"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
)

func JWTAuth(r *gin.Context) {
	//除了登录登出之外 都要验证是否携带token 以及是否有效
	requestUrl := r.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "登入登出不需要验证token")
		r.Next() //r.Next()表示继续执行后续的中间件或路由,后面执行完了再返回到这里
		return   //退出中间件
	}

	// 验证token
	returnData := config.NewReturnData()
	tokenString := r.Request.Header.Get("Authorization")
	returnData.Status = 400
	returnData.Message = "请求未携带TOKEN，请登录后重试"
	if tokenString == "" {
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	// 解析token

	claims, err := jwtutil.ParseToken(tokenString)
	/*
		gin.H返回类型太麻烦 已弃用 修改为config.returnData结构体
		if err != nil {
			r.JSON(200, gin.H{
				"status":  401,
				"message": "Token验证不通过",
			})
		r.Abort() */
	returnData.Status = 400
	returnData.Message = "TOKEN验证不通过"
	if err != nil {
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	r.Set("claims", claims) //把一个变量放进gin的上下文中，后续的路由可以直接使用这个变量
	r.Next()                //验证成功后继续执行下一个中间件或者路由函数
}
