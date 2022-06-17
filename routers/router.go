package routers

import (
	"entry-task-web/middleware/jwt"
	"entry-task-web/middleware/request"
	"github.com/gin-gonic/gin" //nolint:goimports

	"entry-task-web/routers/api"
	"entry-task-web/routers/api/v1"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(request.GenRequestID())
	r.POST("/register", api.Register)
	r.POST("/login", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// 获取用户基本信息
		apiV1.GET("/profile", v1.GetProfile)
		// 获取头像信息
		apiV1.GET("/profile/image", v1.GetHeadImage)
		// 编辑用户信息
		apiV1.POST("/profile", v1.EditProfile)
	}

	return r
}
