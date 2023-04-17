package controllers

import (
	"xxxiklaus/email-verification/models"
	_ "xxxiklaus/email-verification/models"

	"github.com/gin-gonic/gin"
)

type errorString struct {
	string
}

// Ping godoc
//
//	@Summary	测试连通
//	@Schemes
//	@Tags		Health
//	@Produce	json
//
// @Success 200 object models.Resp 返回列表
// @Failure 500 object models.Resp 查询失败
//
//	@Router		/ping [get]
func Ping(c *gin.Context) {
	models.OkWithMessage("pong", c)
}
