package controller

import (
	"casbindemo/logger"
	"github.com/gin-gonic/gin"
)

// 获取所有
func GetNewsList(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	roleStr := role.(string)
	logger.Logger.Debugf("%s 操作了 GetNewsList", roleStr)
}

// 获取单条
func GetNewsDetail(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	roleStr := role.(string)
	logger.Logger.Debugf("%s 操作了 GetNewsDetail", roleStr)
}

// 添加
func AddNews(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	roleStr := role.(string)
	logger.Logger.Debugf("%s 操作了 AddNews", roleStr)
}

// 编辑
func EditNews(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	roleStr := role.(string)
	logger.Logger.Debugf("%s 操作了 EditNews", roleStr)
}

// 删除
func DelNews(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	roleStr := role.(string)
	logger.Logger.Debugf("%s 操作了 DelNews", roleStr)
}