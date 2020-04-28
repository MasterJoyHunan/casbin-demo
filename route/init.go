package route

import (
	"casbindemo/config"
	"casbindemo/controller"
	"casbindemo/pkg/casbin"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(config.ApplicationConfig.Env)
	route := gin.Default()

	// 假设登录,一般都用 JWT ,这里假设从 JWT 中解析了角色（用户）
	route.Use(func(ctx *gin.Context) {
		role := ctx.Request.Header.Get("Authority")
		// role 就是 sub 的抽象，在 RBAC 中可以理解 role 为角色，在 ACL 中可以理解 role 为用户
		ctx.Set("role", role)
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
