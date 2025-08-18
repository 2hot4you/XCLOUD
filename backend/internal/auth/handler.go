package auth

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/sirupsen/logrus"
    "gorm.io/gorm"

    "xcloud-backend/internal/user"
    "xcloud-backend/pkg/jwt"
    "xcloud-backend/pkg/logger"
)

type Handler struct {
    db         *gorm.DB
    userSvc    *user.Service
    jwtManager *jwt.JWTManager
    logger     *logrus.Logger
}

func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        db:         db,
        userSvc:    user.NewService(db),
        jwtManager: jwt.NewJWTManager(),
        logger:     logger.GetLogger(),
    }
}

// Login 用户登录
// @Summary 用户登录
// @Description 使用用户名和密码登录系统
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body LoginRequest true "登录信息"
// @Success 200 {object} LoginResponse "登录成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "认证失败"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        h.logger.Error("登录请求参数错误:", err)
        c.JSON(http.StatusBadRequest, ErrorResponse{
            Code:    400,
            Message: "请求参数错误",
            Error:   err.Error(),
        })
        return
    }

    // 验证用户身份
    user, err := h.userSvc.AuthenticateUser(req.Username, req.Password)
    if err != nil {
        h.logger.Warn("用户登录失败:", req.Username, err)
        c.JSON(http.StatusUnauthorized, ErrorResponse{
            Code:    401,
            Message: "用户名或密码错误",
            Error:   err.Error(),
        })
        return
    }

    // 生成JWT令牌
    accessToken, refreshToken, err := h.jwtManager.GenerateTokens(
        user.ID.String(),
        user.Username,
        string(user.Role),
    )
    if err != nil {
        h.logger.Error("生成JWT令牌失败:", err)
        c.JSON(http.StatusInternalServerError, ErrorResponse{
            Code:    500,
            Message: "登录失败",
            Error:   "Failed to generate tokens",
        })
        return
    }

    h.logger.Info("用户登录成功:", user.Username)
    c.JSON(http.StatusOK, LoginResponse{
        Code:    200,
        Message: "登录成功",
        Data: TokenData{
            AccessToken:  accessToken,
            RefreshToken: refreshToken,
            ExpiresIn:    3600, // 1小时
            TokenType:    "Bearer",
        },
    })
}

// Refresh 刷新令牌
// @Summary 刷新访问令牌
// @Description 使用刷新令牌获取新的访问令牌
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body RefreshRequest true "刷新令牌"
// @Success 200 {object} LoginResponse "刷新成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "令牌无效"
// @Router /auth/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
    var req RefreshRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        h.logger.Error("刷新令牌请求参数错误:", err)
        c.JSON(http.StatusBadRequest, ErrorResponse{
            Code:    400,
            Message: "请求参数错误",
            Error:   err.Error(),
        })
        return
    }

    // 解析刷新令牌
    claims, err := h.jwtManager.ParseToken(req.RefreshToken)
    if err != nil {
        h.logger.Warn("无效的刷新令牌:", err)
        c.JSON(http.StatusUnauthorized, ErrorResponse{
            Code:    401,
            Message: "刷新令牌无效",
            Error:   err.Error(),
        })
        return
    }

    // 检查令牌是否过期
    if time.Now().Unix() > claims.ExpiresAt.Unix() {
        h.logger.Warn("刷新令牌已过期:", claims.Subject)
        c.JSON(http.StatusUnauthorized, ErrorResponse{
            Code:    401,
            Message: "刷新令牌已过期",
            Error:   "Refresh token expired",
        })
        return
    }

    // 验证用户是否仍然存在且活跃
    userID, err := uuid.Parse(claims.UserID)
    if err != nil {
        h.logger.Error("用户ID格式错误:", claims.UserID, err)
        c.JSON(http.StatusUnauthorized, ErrorResponse{
            Code:    401,
            Message: "令牌无效",
            Error:   "Invalid user ID in token",
        })
        return
    }

    user, err := h.userSvc.GetUserByID(userID)
    if err != nil {
        h.logger.Warn("用户不存在或无效:", userID, err)
        c.JSON(http.StatusUnauthorized, ErrorResponse{
            Code:    401,
            Message: "用户无效",
            Error:   err.Error(),
        })
        return
    }

    if !user.IsActive {
        h.logger.Warn("用户已被禁用:", user.Username)
        c.JSON(http.StatusUnauthorized, ErrorResponse{
            Code:    401,
            Message: "用户已被禁用",
            Error:   "User is inactive",
        })
        return
    }

    // 生成新的令牌
    accessToken, refreshToken, err := h.jwtManager.GenerateTokens(
        user.ID.String(),
        user.Username,
        string(user.Role),
    )
    if err != nil {
        h.logger.Error("生成新令牌失败:", err)
        c.JSON(http.StatusInternalServerError, ErrorResponse{
            Code:    500,
            Message: "刷新失败",
            Error:   "Failed to generate new tokens",
        })
        return
    }

    h.logger.Info("令牌刷新成功:", user.Username)
    c.JSON(http.StatusOK, LoginResponse{
        Code:    200,
        Message: "令牌刷新成功",
        Data: TokenData{
            AccessToken:  accessToken,
            RefreshToken: refreshToken,
            ExpiresIn:    3600, // 1小时
            TokenType:    "Bearer",
        },
    })
}

// Logout 用户登出
// @Summary 用户登出
// @Description 退出登录并销毁令牌
// @Tags 认证
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} BaseResponse "登出成功"
// @Failure 401 {object} ErrorResponse "未认证"
// @Router /auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
    // 从上下文获取用户信息
    username, exists := c.Get("username")
    if !exists {
        h.logger.Warn("登出时无法获取用户信息")
    } else {
        h.logger.Info("用户登出:", username)
    }

    // TODO: 实现令牌黑名单机制
    // 在实际生产环境中，应该将令牌加入黑名单（Redis）
    // 这里简单返回成功响应

    c.JSON(http.StatusOK, BaseResponse{
        Code:    200,
        Message: "登出成功",
    })
}