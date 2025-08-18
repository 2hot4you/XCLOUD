package database

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"xcloud-backend/internal/user"
	"xcloud-backend/pkg/logger"
)

// InitializeData 初始化基础数据
func InitializeData(db *gorm.DB) error {
	log := logger.GetLogger()
	
	// 检查是否已有管理员用户
	var count int64
	if err := db.Model(&user.User{}).Where("role = ?", user.RoleAdmin).Count(&count).Error; err != nil {
		log.Error("检查管理员用户失败:", err)
		return err
	}

	if count > 0 {
		log.Info("管理员用户已存在，跳过初始化")
		return nil
	}

	// 创建默认管理员用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Error("密码加密失败:", err)
		return err
	}

	adminUser := user.User{
		ID:           uuid.New(),
		Username:     "admin",
		Email:        "admin@xcloud.com",
		PasswordHash: string(hashedPassword),
		Role:         user.RoleAdmin,
		IsActive:     true,
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Error("创建默认管理员失败:", err)
		return err
	}

	log.Info("默认管理员用户创建成功 - 用户名: admin, 密码: admin123")
	return nil
}