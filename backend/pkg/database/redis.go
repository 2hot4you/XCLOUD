package database

import (
    "context"
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/spf13/viper"

    appLogger "xcloud-backend/pkg/logger"
)

var rdb *redis.Client

// InitRedis 初始化Redis连接
func InitRedis() (*redis.Client, error) {
    rdb = redis.NewClient(&redis.Options{
        Addr:         viper.GetString("redis.addr"),
        Password:     viper.GetString("redis.password"),
        DB:           viper.GetInt("redis.db"),
        MaxRetries:   viper.GetInt("redis.max_retries"),
        PoolSize:     viper.GetInt("redis.pool_size"),
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolTimeout:  30 * time.Second,
    })

    // 测试连接
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        return nil, err
    }

    appLogger.GetLogger().Info("Redis连接成功")
    return rdb, nil
}

// GetRedis 获取Redis客户端
func GetRedis() *redis.Client {
    return rdb
}

// CloseRedis 关闭Redis连接
func CloseRedis() error {
    if rdb != nil {
        return rdb.Close()
    }
    return nil
}