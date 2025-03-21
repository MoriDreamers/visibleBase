package initcontroller

import (
	"visibleBase/utils/logs"
)

func init() {
	logs.Debug(nil, "开始程序初始化")
	metadataInit()
	//创建clientgo相关服务
	//判断namespace存在
}
