package contract

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// RegisterRoutes 注册合同管理相关路由
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
    handler := NewHandler(db)

    router.GET("", handler.GetContracts)
    router.GET("/:id", handler.GetContract)
    router.POST("", handler.CreateContract)
}