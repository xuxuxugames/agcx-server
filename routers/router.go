package routers

import (
	"github.com/agcx_server/config"
	"github.com/agcx_server/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router 设置路由和公共中间件，返回 Gin Engine 对象
func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = []string{"*"}
	r.Use(cors.New(corsConfig))

	gin.SetMode(config.App.RunMode)

	r.GET("/", controllers.Home)

	// User
	// 用户认证
	r.POST("/user/auth", controllers.UserAuth)
	// 用户注册
	r.POST("/user", controllers.UserCreate)
	// 修改密码
	r.PUT("/user/:user_id/password", controllers.UserPassword)

	return r
}
