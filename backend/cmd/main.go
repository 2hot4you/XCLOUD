package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "gorm.io/gorm"

    "xcloud-backend/internal/auth"
    "xcloud-backend/internal/customer"
    "xcloud-backend/internal/contract"
    "xcloud-backend/internal/user"
    "xcloud-backend/pkg/database"
    "xcloud-backend/pkg/logger"
    "xcloud-backend/pkg/middleware"
    "xcloud-backend/docs"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title XCloud多云对账平台API
// @version 1.0
// @description XCloud多云对账平台的后端API服务
// @termsOfService https://github.com/xcloud/terms

// @contact.name XCloud开发团队
// @contact.email dev@xcloud.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Token, 格式: Bearer <token>

func main() {
    // 初始化配置
    initConfig()

    // 初始化日志
    logger.Init()
    log := logger.GetLogger()

    // 初始化数据库
    db, err := database.InitDB()
    if err != nil {
        log.Fatal("数据库初始化失败:", err)
    }

    // 初始化Redis
    rdb, err := database.InitRedis()
    if err != nil {
        log.Fatal("Redis初始化失败:", err)
    }

    // 初始化基础数据
    if err := database.InitializeData(db); err != nil {
        log.Fatal("基础数据初始化失败:", err)
    }

    // 设置Gin模式
    if viper.GetString("app.mode") == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    // 创建Gin路由
    router := setupRouter(db, rdb)

    // 配置服务器
    srv := &http.Server{
        Addr:    ":" + viper.GetString("server.port"),
        Handler: router,
    }

    // 启动服务器
    go func() {
        log.Info("服务器启动在端口:", viper.GetString("server.port"))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal("服务器启动失败:", err)
        }
    }()

    // 优雅关闭
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Info("正在关闭服务器...")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("服务器强制关闭:", err)
    }

    log.Info("服务器已关闭")
}

func initConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./configs")
    viper.AddConfigPath(".")

    // 设置默认值
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("app.mode", "debug")
    viper.SetDefault("database.host", "localhost")
    viper.SetDefault("database.port", 5432)
    viper.SetDefault("database.name", "xcloud")
    viper.SetDefault("database.username", "xcloud")
    viper.SetDefault("database.password", "xcloud123")
    viper.SetDefault("redis.addr", "localhost:6379")
    viper.SetDefault("redis.password", "")
    viper.SetDefault("redis.db", 0)

    if err := viper.ReadInConfig(); err != nil {
        fmt.Printf("配置文件读取失败，使用默认配置: %v\n", err)
    }
}

func setupRouter(db *gorm.DB, rdb interface{}) *gin.Engine {
    router := gin.New()

    // 中间件
    router.Use(middleware.Logger())
    router.Use(middleware.Recovery())
    router.Use(middleware.CORS())

    // Swagger文档
    docs.SwaggerInfo.BasePath = "/api/v1"
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // 健康检查
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "ok",
            "version": "1.0.0",
            "time":    time.Now().Format(time.RFC3339),
        })
    })

    // API路由组
    v1 := router.Group("/api/v1")
    {
        // 认证路由（不需要JWT认证）
        authGroup := v1.Group("/auth")
        auth.RegisterRoutes(authGroup, db)

        // 需要认证的路由
        authenticated := v1.Group("/")
        authenticated.Use(middleware.JWTAuth())
        {
            // 用户管理路由
            userGroup := authenticated.Group("/users")
            user.RegisterRoutes(userGroup, db)

            // 客户管理路由
            customerGroup := authenticated.Group("/customers")
            customer.RegisterRoutes(customerGroup, db)

            // 合同管理路由
            contractGroup := authenticated.Group("/contracts")
            contract.RegisterRoutes(contractGroup, db)
        }
    }

    return router
}