// 项目总入口
package main

import (
	"JWT-TEST/config"
	_ "JWT-TEST/config"
	"JWT-TEST/utils/jwtutil"
	"JWT-TEST/utils/logs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logs.Info(nil, "开始加载程序配置")
	//测试jwt生成token
	ss, err := jwtutil.GenToken("genshin")
	fmt.Println(ss)
	maps := make(map[string]interface{})
	maps["错误信息"] = err
	logs.Info(maps, "token生成错误日志")
	r.Run(config.Port)
}
