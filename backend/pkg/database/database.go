package database

import (
    "fmt"
    "time"

    "github.com/spf13/viper"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

    appLogger "xcloud-backend/pkg/logger"
)

var db *gorm.DB

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
        viper.GetString("database.host"),
        viper.GetString("database.username"),
        viper.GetString("database.password"),
        viper.GetString("database.name"),
        viper.GetInt("database.port"),
    )

    // 配置GORM日志
    var gormLogger logger.Interface
    if viper.GetString("app.mode") == "debug" {
        gormLogger = logger.Default.LogMode(logger.Info)
    } else {
        gormLogger = logger.Default.LogMode(logger.Error)
    }

    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: gormLogger,
        NowFunc: func() time.Time {
            return time.Now().Local()
        },
    })

    if err != nil {
        return nil, fmt.Errorf("数据库连接失败: %w", err)
    }

    // 获取通用数据库对象 sql.DB，然后使用其提供的功能
    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("获取数据库实例失败: %w", err)
    }

    // 设置连接池参数
    sqlDB.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
    sqlDB.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
    sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("database.conn_max_lifetime")) * time.Second)

    // 测试连接
    if err := sqlDB.Ping(); err != nil {
        return nil, fmt.Errorf("数据库连接测试失败: %w", err)
    }

    appLogger.GetLogger().Info("数据库连接成功")
    return db, nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
    return db
}

// CloseDB 关闭数据库连接
func CloseDB() error {
    if db != nil {
        sqlDB, err := db.DB()
        if err != nil {
            return err
        }
        return sqlDB.Close()
    }
    return nil
}