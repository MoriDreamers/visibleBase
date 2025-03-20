package auth

import (
	"visibleBase/config"
	"visibleBase/utils/jwtutil"
	"visibleBase/utils/logs"

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
	returnData := config.NewReturnData()
	if err := r.ShouldBindBodyWithJSON(&userInfo); err != nil {
		returnData.Message = err.Error()
		r.JSON(200, returnData)
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
			returnData.Message = "token生成失败"
			returnData.Status = 401
			r.JSON(200, returnData)
			return
		}
		//token生成成功的处理
		logs.Info(map[string]interface{}{
			"用户名": userInfo.Username}, "登陆成功")
		returnData.Message = "登陆成功"
		returnData.Status = 200
		returnData.Data["token"] = ss
		r.JSON(200, returnData)
		return
	} else {
		returnData.Message = "用户名或密码错误"
		returnData.Status = 401
		r.JSON(200, returnData)
	}

}
func Logout(r *gin.Context) {
	//登出逻辑
	returnData := config.NewReturnData()
	returnData.Message = "登出成功"
	returnData.Status = 200
	r.JSON(200, returnData)
	logs.Debug(nil, "登出成功")
}
