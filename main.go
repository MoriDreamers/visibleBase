// 项目总入口
package main

import (
	"JWT-TEST/config"
	_ "JWT-TEST/config"
	"JWT-TEST/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logs.Info(nil, "开始加载程序配置")
	r.Run(config.Port)
}
