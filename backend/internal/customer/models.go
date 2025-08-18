package customer

// 请求结构体

// CreateCustomerRequest 创建客户请求
type CreateCustomerRequest struct {
    CustomerCode string  `json:"customer_code" binding:"required" example:"CUST001"`
    CompanyName  string  `json:"company_name" binding:"required" example:"示例科技公司"`
    ContactName  string  `json:"contact_name" binding:"required" example:"张三"`
    ContactPhone string  `json:"contact_phone" example:"13800138000"`
    ContactEmail string  `json:"contact_email" example:"contact@example.com"`
    Address      string  `json:"address" example:"北京市朝阳区示例大厦"`
    ParentID     *string `json:"parent_id,omitempty" example:"parent-uuid"`
    Level        int     `json:"level" example:"1"`
    BusinessType string  `json:"business_type" example:"互联网"`
    CreditLimit  float64 `json:"credit_limit" example:"1000000.00"`
}

// UpdateCustomerRequest 更新客户请求
type UpdateCustomerRequest struct {
    CompanyName  string  `json:"company_name,omitempty" example:"示例科技公司"`
    ContactName  string  `json:"contact_name,omitempty" example:"张三"`
    ContactPhone string  `json:"contact_phone,omitempty" example:"13800138000"`
    ContactEmail string  `json:"contact_email,omitempty" example:"contact@example.com"`
    Address      string  `json:"address,omitempty" example:"北京市朝阳区示例大厦"`
    Status       string  `json:"status,omitempty" example:"active"`
    BusinessType string  `json:"business_type,omitempty" example:"互联网"`
    CreditLimit  float64 `json:"credit_limit,omitempty" example:"1000000.00"`
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

// CustomerData 客户数据
type CustomerData struct {
    ID           string  `json:"id" example:"uuid-string"`
    CustomerCode string  `json:"customer_code" example:"CUST001"`
    CompanyName  string  `json:"company_name" example:"示例科技公司"`
    ContactName  string  `json:"contact_name" example:"张三"`
    ContactPhone string  `json:"contact_phone" example:"13800138000"`
    ContactEmail string  `json:"contact_email" example:"contact@example.com"`
    Address      string  `json:"address" example:"北京市朝阳区示例大厦"`
    Status       string  `json:"status" example:"active"`
    ParentID     *string `json:"parent_id,omitempty" example:"parent-uuid"`
    Level        int     `json:"level" example:"1"`
    BusinessType string  `json:"business_type" example:"互联网"`
    CreditLimit  float64 `json:"credit_limit" example:"1000000.00"`
    CreatedAt    string  `json:"created_at" example:"2024-01-01T00:00:00Z"`
    UpdatedAt    string  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// CustomerResponse 客户响应
type CustomerResponse struct {
    Code    int          `json:"code" example:"200"`
    Message string       `json:"message" example:"操作成功"`
    Data    CustomerData `json:"data"`
}

// CustomerListData 客户列表数据
type CustomerListData struct {
    Customers []CustomerData `json:"customers"`
    Total     int64          `json:"total" example:"100"`
    Page      int            `json:"page" example:"1"`
    PageSize  int            `json:"page_size" example:"20"`
}

// CustomerListResponse 客户列表响应
type CustomerListResponse struct {
    Code    int              `json:"code" example:"200"`
    Message string           `json:"message" example:"获取成功"`
    Data    CustomerListData `json:"data"`
}