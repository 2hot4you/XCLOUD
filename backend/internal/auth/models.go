package auth

// 请求结构体

// LoginRequest 登录请求
type LoginRequest struct {
    Username string `json:"username" binding:"required" example:"admin"`
    Password string `json:"password" binding:"required" example:"password123"`
}

// RefreshRequest 刷新令牌请求
type RefreshRequest struct {
    RefreshToken string `json:"refresh_token" binding:"required" example:"refresh_token_here"`
}

// 响应结构体

// BaseResponse 基础响应
type BaseResponse struct {
    Code    int    `json:"code" example:"200"`
    Message string `json:"message" example:"操作成功"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
    Code    int    `json:"code" example:"400"`
    Message string `json:"message" example:"请求参数错误"`
    Error   string `json:"error,omitempty" example:"具体错误信息"`
}

// TokenData 令牌数据
type TokenData struct {
    AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
    RefreshToken string `json:"refresh_token" example:"refresh_token_example"`
    ExpiresIn    int    `json:"expires_in" example:"86400"`
    TokenType    string `json:"token_type" example:"Bearer"`
}

// LoginResponse 登录响应
type LoginResponse struct {
    Code    int       `json:"code" example:"200"`
    Message string    `json:"message" example:"登录成功"`
    Data    TokenData `json:"data"`
}