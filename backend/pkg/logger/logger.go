package logger

import (
    "os"

    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

var logger *logrus.Logger

// Init 初始化日志
func Init() {
    logger = logrus.New()

    // 设置日志级别
    level := viper.GetString("log.level")
    switch level {
    case "trace":
        logger.SetLevel(logrus.TraceLevel)
    case "debug":
        logger.SetLevel(logrus.DebugLevel)
    case "info":
        logger.SetLevel(logrus.InfoLevel)
    case "warn":
        logger.SetLevel(logrus.WarnLevel)
    case "error":
        logger.SetLevel(logrus.ErrorLevel)
    case "fatal":
        logger.SetLevel(logrus.FatalLevel)
    case "panic":
        logger.SetLevel(logrus.PanicLevel)
    default:
        logger.SetLevel(logrus.InfoLevel)
    }

    // 设置日志格式
    format := viper.GetString("log.format")
    if format == "json" {
        logger.SetFormatter(&logrus.JSONFormatter{
            TimestampFormat: "2006-01-02 15:04:05",
        })
    } else {
        logger.SetFormatter(&logrus.TextFormatter{
            FullTimestamp:   true,
            TimestampFormat: "2006-01-02 15:04:05",
        })
    }

    // 设置输出
    output := viper.GetString("log.output")
    if output == "file" {
        file, err := os.OpenFile("xcloud.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            logger.Fatal("无法打开日志文件:", err)
        }
        logger.SetOutput(file)
    } else {
        logger.SetOutput(os.Stdout)
    }
}

// GetLogger 获取日志实例
func GetLogger() *logrus.Logger {
    if logger == nil {
        Init()
    }
    return logger
}