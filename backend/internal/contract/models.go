package contract

// 请求结构体

// CreateContractRequest 创建合同请求
type CreateContractRequest struct {
    ContractNo       string  `json:"contract_no" binding:"required" example:"CON202401001"`
    CustomerID       string  `json:"customer_id" binding:"required" example:"customer-uuid"`
    Title            string  `json:"title" binding:"required" example:"云服务代理合同"`
    StartDate        string  `json:"start_date" binding:"required" example:"2024-01-01"`
    EndDate          string  `json:"end_date" binding:"required" example:"2024-12-31"`
    SettlementCycle  int     `json:"settlement_cycle" example:"1"`
    PaymentTerms     string  `json:"payment_terms" example:"月结30天"`
    ContractAmount   float64 `json:"contract_amount" example:"1000000.00"`
    DiscountRate     float64 `json:"discount_rate" example:"0.85"`
}

// UpdateContractRequest 更新合同请求
type UpdateContractRequest struct {
    Title           string  `json:"title,omitempty" example:"云服务代理合同"`
    Status          string  `json:"status,omitempty" example:"active"`
    EndDate         string  `json:"end_date,omitempty" example:"2024-12-31"`
    SettlementCycle int     `json:"settlement_cycle,omitempty" example:"1"`
    PaymentTerms    string  `json:"payment_terms,omitempty" example:"月结30天"`
    ContractAmount  float64 `json:"contract_amount,omitempty" example:"1000000.00"`
    DiscountRate    float64 `json:"discount_rate,omitempty" example:"0.85"`
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

// ContractData 合同数据
type ContractData struct {
    ID              string  `json:"id" example:"uuid-string"`
    ContractNo      string  `json:"contract_no" example:"CON202401001"`
    CustomerID      string  `json:"customer_id" example:"customer-uuid"`
    Title           string  `json:"title" example:"云服务代理合同"`
    Status          string  `json:"status" example:"active"`
    StartDate       string  `json:"start_date" example:"2024-01-01"`
    EndDate         string  `json:"end_date" example:"2024-12-31"`
    SettlementCycle int     `json:"settlement_cycle" example:"1"`
    PaymentTerms    string  `json:"payment_terms" example:"月结30天"`
    ContractAmount  float64 `json:"contract_amount" example:"1000000.00"`
    DiscountRate    float64 `json:"discount_rate" example:"0.85"`
    CreatedAt       string  `json:"created_at" example:"2024-01-01T00:00:00Z"`
    UpdatedAt       string  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ContractResponse 合同响应
type ContractResponse struct {
    Code    int          `json:"code" example:"200"`
    Message string       `json:"message" example:"操作成功"`
    Data    ContractData `json:"data"`
}

// ContractListData 合同列表数据
type ContractListData struct {
    Contracts []ContractData `json:"contracts"`
    Total     int64          `json:"total" example:"100"`
    Page      int            `json:"page" example:"1"`
    PageSize  int            `json:"page_size" example:"20"`
}

// ContractListResponse 合同列表响应
type ContractListResponse struct {
    Code    int              `json:"code" example:"200"`
    Message string           `json:"message" example:"获取成功"`
    Data    ContractListData `json:"data"`
}