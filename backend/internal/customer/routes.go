package customer

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// RegisterRoutes 注册客户管理相关路由
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
    handler := NewHandler(db)

    router.GET("", handler.GetCustomers)
    router.GET("/:id", handler.GetCustomer)
    router.POST("", handler.CreateCustomer)
    router.PUT("/:id", handler.UpdateCustomer)
    router.DELETE("/:id", handler.DeleteCustomer)
}