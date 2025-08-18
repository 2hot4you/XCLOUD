package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"xcloud-backend/pkg/logger"
)

type Handler struct {
	userSvc *Service
	logger  *logrus.Logger
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		userSvc: NewService(db),
		logger:  logger.GetLogger(),
	}
}

// GetProfile 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取当前登录用户的详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} UserProfileResponse "用户信息"
// @Failure 401 {object} ErrorResponse "未认证"
// @Router /users/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未认证",
		})
		return
	}

	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		h.logger.Error("用户ID格式错误:", userID, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	user, err := h.userSvc.GetUserByID(uid)
	if err != nil {
		h.logger.Error("获取用户信息失败:", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, UserProfileResponse{
		Code:    200,
		Message: "获取成功",
		Data:    user.ToResponse(),
	})
}

// ListUsers 获取用户列表
// @Summary 获取用户列表
// @Description 分页获取用户列表（需要管理员权限）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页大小" default(10)
// @Success 200 {object} UserListResponse "用户列表"
// @Failure 401 {object} ErrorResponse "未认证"
// @Failure 403 {object} ErrorResponse "权限不足"
// @Router /users [get]
func (h *Handler) ListUsers(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	users, total, err := h.userSvc.ListUsers(page, pageSize)
	if err != nil {
		h.logger.Error("获取用户列表失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户列表失败",
		})
		return
	}

	// 转换为响应格式
	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	c.JSON(http.StatusOK, UserListResponse{
		Code:    200,
		Message: "获取成功",
		Data: UserListData{
			Users: userResponses,
			Pagination: PaginationInfo{
				Page:      page,
				PageSize:  pageSize,
				Total:     total,
				TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
			},
		},
	})
}

// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建新用户（需要管理员权限）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body UserCreateRequest true "用户信息"
// @Success 201 {object} UserProfileResponse "用户创建成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未认证"
// @Failure 403 {object} ErrorResponse "权限不足"
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("创建用户请求参数错误:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 获取创建者ID
	createdBy, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未认证",
		})
		return
	}

	createdByUUID, err := uuid.Parse(createdBy.(string))
	if err != nil {
		h.logger.Error("创建者ID格式错误:", createdBy, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	user, err := h.userSvc.CreateUser(req, createdByUUID)
	if err != nil {
		h.logger.Error("创建用户失败:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	h.logger.Info("用户创建成功:", user.Username)
	c.JSON(http.StatusCreated, UserProfileResponse{
		Code:    201,
		Message: "用户创建成功",
		Data:    user.ToResponse(),
	})
}

// UpdateUser 更新用户
// @Summary 更新用户信息
// @Description 更新指定用户的信息（需要管理员权限）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Param body body UserUpdateRequest true "更新信息"
// @Success 200 {object} UserProfileResponse "更新成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未认证"
// @Failure 403 {object} ErrorResponse "权限不足"
// @Failure 404 {object} ErrorResponse "用户不存在"
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("用户ID格式错误:", idStr, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	var req UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("更新用户请求参数错误:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 获取更新者ID
	updatedBy, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未认证",
		})
		return
	}

	updatedByUUID, err := uuid.Parse(updatedBy.(string))
	if err != nil {
		h.logger.Error("更新者ID格式错误:", updatedBy, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	user, err := h.userSvc.UpdateUser(userID, req, updatedByUUID)
	if err != nil {
		h.logger.Error("更新用户失败:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	h.logger.Info("用户更新成功:", user.Username)
	c.JSON(http.StatusOK, UserProfileResponse{
		Code:    200,
		Message: "更新成功",
		Data:    user.ToResponse(),
	})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除指定用户（软删除，需要管理员权限）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Success 200 {object} BaseResponse "删除成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未认证"
// @Failure 403 {object} ErrorResponse "权限不足"
// @Failure 404 {object} ErrorResponse "用户不存在"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("用户ID格式错误:", idStr, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	// 获取删除者ID
	deletedBy, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未认证",
		})
		return
	}

	deletedByUUID, err := uuid.Parse(deletedBy.(string))
	if err != nil {
		h.logger.Error("删除者ID格式错误:", deletedBy, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	// 不允许删除自己
	if userID == deletedByUUID {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "不能删除自己的账户",
		})
		return
	}

	err = h.userSvc.DeleteUser(userID, deletedByUUID)
	if err != nil {
		h.logger.Error("删除用户失败:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	h.logger.Info("用户删除成功:", userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 修改当前用户的密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body ChangePasswordRequest true "密码信息"
// @Success 200 {object} BaseResponse "修改成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未认证"
// @Router /users/change-password [post]
func (h *Handler) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("修改密码请求参数错误:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未认证",
		})
		return
	}

	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		h.logger.Error("用户ID格式错误:", userID, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
		})
		return
	}

	err = h.userSvc.ChangePassword(uid, req.OldPassword, req.NewPassword)
	if err != nil {
		h.logger.Error("修改密码失败:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	username, _ := c.Get("username")
	h.logger.Info("密码修改成功:", username)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "密码修改成功",
	})
}

// 响应结构体
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type UserProfileResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type UserListResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    UserListData `json:"data"`
}

type UserListData struct {
	Users      []UserResponse `json:"users"`
	Pagination PaginationInfo `json:"pagination"`
}

type PaginationInfo struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" example:"old123"`
	NewPassword string `json:"new_password" binding:"required,min=6" example:"new123"`
}