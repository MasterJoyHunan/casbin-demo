package route

import (
	"casbindemo/config"
	"casbindemo/controller"
	"casbindemo/pkg/casbin"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(config.ApplicationConfig.Env)
	route := gin.New()

	// 假设登录,一般都用 JWT ,这里默认设置
	route.Use(func(ctx *gin.Context) {
		ctx.Set("role", "admin")
		ctx.Next()
	})

	route.Use(casbin.Authentication)

	// 查询列表
	route.GET("/news", controller.GetNewsList)

	// 查询单条
	route.GET("/news/:id", controller.GetNewsDetail)

	// 新增
	route.POST("/news", controller.AddNews)

	// 修改
	route.PUT("/news/:id", controller.EditNews)

	// 删除
	route.DELETE("/news/:id", controller.DelNews)
	return route
}

