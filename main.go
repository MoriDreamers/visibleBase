// 项目总入口
package main

import (
	"JWT-TEST/config"
	_ "JWT-TEST/config"
	"JWT-TEST/middlewares"
	"JWT-TEST/routers"
	"JWT-TEST/utils/logs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	logs.Info(nil, "开始加载程序配置")
	//本地调试时取消跨域想限制
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1200 * time.Hour,
	}))

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
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
