package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole 用户角色枚举
type UserRole string

const (
	RoleAdmin    UserRole = "admin"    // 管理员
	RoleManager  UserRole = "manager"  // 经理
	RoleEmployee UserRole = "employee" // 员工
	RoleViewer   UserRole = "viewer"   // 查看者
)

// User 用户模型
type User struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username     string         `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email        string         `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"type:varchar(255);not null"`
	Role         UserRole       `json:"role" gorm:"type:varchar(20);not null;default:'viewer'"`
	IsActive     bool           `json:"is_active" gorm:"not null;default:true"`
	LastLoginAt  *time.Time     `json:"last_login_at,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	CreatedBy    *uuid.UUID     `json:"created_by,omitempty"`
	UpdatedBy    *uuid.UUID     `json:"updated_by,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}

// IsValidRole 检查角色是否有效
func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleManager, RoleEmployee, RoleViewer:
		return true
	default:
		return false
	}
}

// HasPermission 检查用户是否有指定权限
func (u *User) HasPermission(requiredRole UserRole) bool {
	switch u.Role {
	case RoleAdmin:
		return true // 管理员拥有所有权限
	case RoleManager:
		return requiredRole == RoleManager || requiredRole == RoleEmployee || requiredRole == RoleViewer
	case RoleEmployee:
		return requiredRole == RoleEmployee || requiredRole == RoleViewer
	case RoleViewer:
		return requiredRole == RoleViewer
	default:
		return false
	}
}

// UserCreateRequest 用户创建请求
type UserCreateRequest struct {
	Username string   `json:"username" binding:"required,min=3,max=50" example:"johndoe"`
	Email    string   `json:"email" binding:"required,email,max=100" example:"john@example.com"`
	Password string   `json:"password" binding:"required,min=6,max=50" example:"password123"`
	Role     UserRole `json:"role" binding:"required" example:"employee"`
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	Username string   `json:"username,omitempty" binding:"omitempty,min=3,max=50" example:"johndoe"`
	Email    string   `json:"email,omitempty" binding:"omitempty,email,max=100" example:"john@example.com"`
	Role     UserRole `json:"role,omitempty" binding:"omitempty" example:"employee"`
	IsActive *bool    `json:"is_active,omitempty" example:"true"`
}

// UserResponse 用户响应
type UserResponse struct {
	ID          uuid.UUID  `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Role        UserRole   `json:"role"`
	IsActive    bool       `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ToResponse 转换为响应格式
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		Role:        u.Role,
		IsActive:    u.IsActive,
		LastLoginAt: u.LastLoginAt,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}