package casbin

import (
	"casbindemo/logger"
	"casbindemo/model"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	"github.com/gin-gonic/gin"
	"os"
)

// 鉴权
func Authentication(ctx *gin.Context) {
	data, _ := ctx.Get("role")
	role := data.(string)
	e, err := initCasbin()
	if err != nil {
		logger.Logger.Panic("初始化 Casbin 出现错误：", err)
	}
	ok, err := e.Enforce(role, ctx.FullPath(), ctx.Request.Method)
	logger.Logger.Debugf("sub:[%s] obj:[%s] act:[%s] res:[%t]", role, ctx.FullPath(), ctx.Request.Method, ok)
	if err != nil {
		logger.Logger.Panic("执行 e.Enforce() 出错：", err)
	}
	if ok {
		ctx.Next()
	} else {
		ctx.JSON(403, gin.H{
			"code": 403,
			"msg":  "权限不足",
		})
		ctx.Abort()
	}
}

// 初始化 casbin
func initCasbin() (e *casbin.Enforcer, err error) {
	adapter, err := gormadapter.NewAdapterByDB(model.Db)
	if err != nil {
		return
	}

	wd, _ := os.Getwd()
	e, err = casbin.NewEnforcer(wd+"/rbac_model.conf", adapter)
	if err != nil {
		return
	}

	if err = e.LoadPolicy(); err != nil {
		return
	}
	return
}
