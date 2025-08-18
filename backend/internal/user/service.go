package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Service 用户服务
type Service struct {
	db *gorm.DB
}

// NewService 创建用户服务
func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// AuthenticateUser 验证用户登录
func (s *Service) AuthenticateUser(username, password string) (*User, error) {
	var user User
	err := s.db.Where("username = ? AND is_active = true", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLoginAt = &now
	s.db.Model(&user).Update("last_login_at", now)

	return &user, nil
}

// GetUserByID 根据ID获取用户
func (s *Service) GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	err := s.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (s *Service) GetUserByUsername(username string) (*User, error) {
	var user User
	err := s.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建用户
func (s *Service) CreateUser(req UserCreateRequest, createdBy uuid.UUID) (*User, error) {
	// 检查用户名是否已存在
	var existingUser User
	err := s.db.Where("username = ?", req.Username).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("用户名已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 检查邮箱是否已存在
	err = s.db.Where("email = ?", req.Email).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("邮箱已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 验证角色
	if !req.Role.IsValid() {
		return nil, errors.New("无效的用户角色")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
		IsActive:     true,
		CreatedBy:    &createdBy,
		UpdatedBy:    &createdBy,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser 更新用户
func (s *Service) UpdateUser(id uuid.UUID, req UserUpdateRequest, updatedBy uuid.UUID) (*User, error) {
	var user User
	err := s.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 更新字段
	if req.Username != "" {
		// 检查用户名是否已被其他用户使用
		var existingUser User
		err := s.db.Where("username = ? AND id != ?", req.Username, id).First(&existingUser).Error
		if err == nil {
			return nil, errors.New("用户名已存在")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		user.Username = req.Username
	}

	if req.Email != "" {
		// 检查邮箱是否已被其他用户使用
		var existingUser User
		err := s.db.Where("email = ? AND id != ?", req.Email, id).First(&existingUser).Error
		if err == nil {
			return nil, errors.New("邮箱已存在")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		user.Email = req.Email
	}

	if req.Role != "" {
		if !req.Role.IsValid() {
			return nil, errors.New("无效的用户角色")
		}
		user.Role = req.Role
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	user.UpdatedBy = &updatedBy

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser 删除用户（软删除）
func (s *Service) DeleteUser(id uuid.UUID, deletedBy uuid.UUID) error {
	var user User
	err := s.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	user.UpdatedBy = &deletedBy
	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	return s.db.Delete(&user).Error
}

// ListUsers 获取用户列表
func (s *Service) ListUsers(page, pageSize int) ([]User, int64, error) {
	var users []User
	var total int64

	// 计算总数
	if err := s.db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := s.db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ChangePassword 修改密码
func (s *Service) ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error {
	var user User
	err := s.db.First(&user, "id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return errors.New("原密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	user.UpdatedBy = &userID

	return s.db.Save(&user).Error
}