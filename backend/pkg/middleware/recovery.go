package middleware

import (
    "net"
    "net/http"
    "net/http/httputil"
    "os"
    "runtime/debug"
    "strings"

    "github.com/gin-gonic/gin"
    appLogger "xcloud-backend/pkg/logger"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
    logger := appLogger.GetLogger()

    return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
        // 检查连接是否断开
        if err, ok := recovered.(error); ok {
            if isBrokenPipeError(err) {
                logger.Error(c.Request.URL.Path, " 连接断开: ", err)
                c.Error(err.(error))
                c.Abort()
                return
            }
        }

        // 记录panic信息
        httpRequest, _ := httputil.DumpRequest(c.Request, false)
        logger.WithFields(map[string]interface{}{
            "error":   recovered,
            "request": string(httpRequest),
            "stack":   string(debug.Stack()),
        }).Error("服务器panic恢复")

        c.JSON(http.StatusInternalServerError, gin.H{
            "code":    500,
            "message": "服务器内部错误",
            "error":   "Internal Server Error",
        })
    })
}

// isBrokenPipeError 检查是否是连接断开错误
func isBrokenPipeError(err error) bool {
    if netErr, ok := err.(*net.OpError); ok {
        if sysErr, ok := netErr.Err.(*os.SyscallError); ok {
            if strings.Contains(strings.ToLower(sysErr.Error()), "broken pipe") ||
                strings.Contains(strings.ToLower(sysErr.Error()), "connection reset by peer") {
                return true
            }
        }
    }
    return false
}