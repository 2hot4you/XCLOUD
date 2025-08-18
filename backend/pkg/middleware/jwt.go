package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    
    jwtPkg "xcloud-backend/pkg/jwt"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
    jwtManager := jwtPkg.NewJWTManager()
    
    return func(c *gin.Context) {
        // 获取Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code":    401,
                "message": "请提供认证令牌",
                "error":   "Authorization header required",
            })
            c.Abort()
            return
        }

        // 检查Bearer前缀
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code":    401,
                "message": "认证令牌格式错误",
                "error":   "Authorization header format must be Bearer {token}",
            })
            c.Abort()
            return
        }

        tokenString := parts[1]

        // 验证JWT令牌
        claims, err := jwtManager.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code":    401,
                "message": "认证令牌无效",
                "error":   err.Error(),
            })
            c.Abort()
            return
        }

        // 将用户信息存储到上下文
        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)
        c.Set("user_role", claims.Role)
        c.Next()
    }
}

// RequireRole 角色权限检查中间件
func RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("user_role")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{
                "code":    403,
                "message": "无法获取用户角色信息",
                "error":   "User role not found",
            })
            c.Abort()
            return
        }

        role := userRole.(string)
        for _, requiredRole := range roles {
            if role == requiredRole {
                c.Next()
                return
            }
        }

        c.JSON(http.StatusForbidden, gin.H{
            "code":    403,
            "message": "权限不足",
            "error":   "Insufficient permissions",
        })
        c.Abort()
    }
}