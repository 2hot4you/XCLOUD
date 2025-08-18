package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"xcloud-backend/pkg/middleware"
)

// RegisterRoutes 注册用户相关路由
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	// 当前用户信息
	router.GET("/profile", handler.GetProfile)
	router.POST("/change-password", handler.ChangePassword)

	// 用户管理（需要管理员权限）
	adminRoutes := router.Group("")
	adminRoutes.Use(middleware.RequireRole("admin"))
	{
		adminRoutes.GET("", handler.ListUsers)
		adminRoutes.POST("", handler.CreateUser)
		adminRoutes.PUT("/:id", handler.UpdateUser)
		adminRoutes.DELETE("/:id", handler.DeleteUser)
	}
}