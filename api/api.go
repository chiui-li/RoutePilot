package api

import (
	publicApi "RoutePilot/api/public"
	userApi "RoutePilot/api/user"

	"github.com/gin-gonic/gin"
)

func ApiServer() {
	// 初始化 Gin 框架
	r := gin.Default()

	// 注册路由
	publicRouter := r.Group("/api/public")
	publicRouter.Any("/ping", publicApi.PingHandler)
	publicRouter.Any("/login", publicApi.LoginHandler)

	userRouter := r.Group("/api/user", AuthMiddleware())
	userRouter.Any("/whoami", userApi.WhoamiHandler)

	r.Run(":8080")
}
