package middleware

import (
    "github.com/gin-gonic/gin"
    appLogger "xcloud-backend/pkg/logger"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
    logger := appLogger.GetLogger()

    return gin.LoggerWithConfig(gin.LoggerConfig{
        Formatter: func(param gin.LogFormatterParams) string {
            logger.WithFields(map[string]interface{}{
                "status":     param.StatusCode,
                "method":     param.Method,
                "path":       param.Path,
                "ip":         param.ClientIP,
                "user-agent": param.Request.UserAgent(),
                "latency":    param.Latency.String(),
                "error":      param.ErrorMessage,
            }).Info("API请求")
            return ""
        },
        Output: logger.Out,
    })
}