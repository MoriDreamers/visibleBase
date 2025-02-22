// 项目总入口
package main

import (
	"JWT-TEST/config"
	_ "JWT-TEST/config"
	"JWT-TEST/middlewares"
	"JWT-TEST/routers"
	"JWT-TEST/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.RegisterRouters(r)
	logs.Info(nil, "开始加载程序配置")
	r.Use(middlewares.JWTAuth)
	/*
		//测试jwt生成token
		ss, err := jwtutil.GenToken("genshin")
		fmt.Println(ss)
		maps := make(map[string]interface{})
		maps["错误信息"] = err
		logs.Info(maps, "token生成错误日志")
		//验证解析token的方法
		_, err2 := jwtutil.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJhciIsImlzcyI6Ik1vcmkiLCJzdWIiOiJNb3JpIiwiZXhwIjoxNzQwMTk5MTYzLCJuYmYiOjE3NDAxOTE5NjMsImlhdCI6MTc0MDE5MTk2M30.SNRHKUoRsmQbHVwfYrate6RNRVDidw5K2meTi4nJPdY")
		if err2 != nil {
			fmt.Println("解析token失败", err2.Error())
		}
	*/
	r.Run(config.Port)
}
