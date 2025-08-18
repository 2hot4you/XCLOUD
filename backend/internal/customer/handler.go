package customer

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Handler struct {
    db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
    return &Handler{db: db}
}

// GetCustomers 获取客户列表
// @Summary 获取客户列表
// @Description 分页获取客户列表
// @Tags 客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Param search query string false "搜索关键词"
// @Success 200 {object} CustomerListResponse "获取成功"
// @Failure 401 {object} ErrorResponse "未认证"
// @Router /customers [get]
func (h *Handler) GetCustomers(c *gin.Context) {
    // TODO: 实现获取客户列表逻辑
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "获取客户列表成功",
        "data": gin.H{
            "customers": []gin.H{},
            "total":     0,
            "page":      1,
            "page_size": 20,
        },
    })
}

// GetCustomer 获取客户详情
// @Summary 获取客户详情
// @Description 根据ID获取客户详细信息
// @Tags 客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户ID"
// @Success 200 {object} CustomerResponse "获取成功"
// @Failure 404 {object} ErrorResponse "客户不存在"
// @Router /customers/{id} [get]
func (h *Handler) GetCustomer(c *gin.Context) {
    id := c.Param("id")
    // TODO: 实现获取客户详情逻辑
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "获取客户详情成功",
        "data": gin.H{
            "id":           id,
            "company_name": "示例公司",
            "status":       "active",
        },
    })
}

// CreateCustomer 创建客户
// @Summary 创建客户
// @Description 创建新的客户记录
// @Tags 客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body CreateCustomerRequest true "客户信息"
// @Success 201 {object} CustomerResponse "创建成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Router /customers [post]
func (h *Handler) CreateCustomer(c *gin.Context) {
    var req CreateCustomerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "请求参数错误",
            "error":   err.Error(),
        })
        return
    }

    // TODO: 实现创建客户逻辑
    c.JSON(http.StatusCreated, gin.H{
        "code":    201,
        "message": "客户创建成功",
        "data": gin.H{
            "id":           "new-customer-id",
            "company_name": req.CompanyName,
            "status":       "active",
        },
    })
}

// UpdateCustomer 更新客户
// @Summary 更新客户信息
// @Description 根据ID更新客户信息
// @Tags 客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户ID"
// @Param body body UpdateCustomerRequest true "客户信息"
// @Success 200 {object} CustomerResponse "更新成功"
// @Failure 404 {object} ErrorResponse "客户不存在"
// @Router /customers/{id} [put]
func (h *Handler) UpdateCustomer(c *gin.Context) {
    id := c.Param("id")
    var req UpdateCustomerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "请求参数错误",
            "error":   err.Error(),
        })
        return
    }

    // TODO: 实现更新客户逻辑
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "客户更新成功",
        "data": gin.H{
            "id":           id,
            "company_name": req.CompanyName,
            "status":       "active",
        },
    })
}

// DeleteCustomer 删除客户
// @Summary 删除客户
// @Description 根据ID删除客户（软删除）
// @Tags 客户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "客户ID"
// @Success 200 {object} BaseResponse "删除成功"
// @Failure 404 {object} ErrorResponse "客户不存在"
// @Router /customers/{id} [delete]
func (h *Handler) DeleteCustomer(c *gin.Context) {
    id := c.Param("id")
    // TODO: 实现删除客户逻辑
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "客户删除成功",
        "data": gin.H{
            "id": id,
        },
    })
}