package auth

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// RegisterRoutes 注册认证相关路由
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
    handler := NewHandler(db)

    router.POST("/login", handler.Login)
    router.POST("/refresh", handler.Refresh)
    router.POST("/logout", handler.Logout)
}