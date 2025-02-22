package auth

import (
	"JWT-TEST/config"
	"JWT-TEST/utils/jwtutil"
	"JWT-TEST/utils/logs"

	"github.com/gin-gonic/gin"
)

// 登陆逻辑
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	//1. 获取前端传递的用户名和密码
	userInfo := UserInfo{}
	if err := r.ShouldBindBodyWithJSON(&userInfo); err != nil {
		r.JSON(200, gin.H{
			"msg":    err.Error(),
			"status": 401,
		})
		return
	}
	logs.Debug(map[string]interface{}{
		"用户名": userInfo.Username,
		"密码":  userInfo.Password,
	}, "用户尝试登陆")
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		//认证成功
		//生成jwttoken
		ss, err := jwtutil.GenToken(userInfo.Username)
		//token生成错误的处理
		if err != nil {
			logs.Error(map[string]interface{}{
				"用户名":  userInfo.Username,
				"错误日志": err.Error()},
				"信息正确但是生成失败！")
			r.JSON(200, gin.H{
				"message": "token生成失败",
				"status":  401,
			})
			return
		}
		//token生成成功的处理
		logs.Info(map[string]interface{}{
			"用户名": userInfo.Username}, "登陆成功")
		data := make(map[string]interface{})
		data["token"] = ss
		r.JSON(200, gin.H{
			"status":  200,
			"message": "登陆成功",
			"data":    data,
		})
		return
	} else {
		r.JSON(200, gin.H{
			"message": "用户名或密码错误",
		})
	}

}
func Logout(r *gin.Context) {
	//登出逻辑
	r.JSON(200, gin.H{
		"message": "登出成功",
		"status":  200,
	})
	logs.Debug(nil, "登出成功")
}
