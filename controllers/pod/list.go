package pod

import (
	"visibleBase/config"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
)

func List(r *gin.Context) {
	logs.Info(nil, "获取namespace列表")
	returnData := config.NewReturnData()
	returnData.Data = make(map[string]interface{})

}
