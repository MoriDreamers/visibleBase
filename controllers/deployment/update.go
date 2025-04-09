package deployment

import (
	"visibleBase/config"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(r *gin.Context) {
	logs.Info(nil, "更新pod")

	returnData := config.NewReturnData() //初始化返回数据
	returnData.Status = 200
	returnData.Message = "POD不支持更新！"
	r.JSON(200, returnData)
}
