-- XCloud多云对账平台数据库表结构设计
-- 创建时间: 2024年
-- 作者: XCloud开发团队

-- 设置时区
SET timezone = 'Asia/Shanghai';

-- 创建枚举类型
CREATE TYPE user_role AS ENUM ('admin', 'operator', 'viewer');
CREATE TYPE customer_status AS ENUM ('active', 'inactive', 'suspended');
CREATE TYPE contract_status AS ENUM ('draft', 'pending', 'active', 'expired', 'terminated');
CREATE TYPE cloud_provider AS ENUM ('tencent', 'alibaba', 'huawei', 'aws');
CREATE TYPE commission_status AS ENUM ('pending', 'calculated', 'paid');

-- 1. 用户表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role user_role NOT NULL DEFAULT 'viewer',
    is_active BOOLEAN NOT NULL DEFAULT true,
    last_login_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP
);

-- 2. 客户表
CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_code VARCHAR(20) UNIQUE NOT NULL, -- 客户编码
    company_name VARCHAR(200) NOT NULL,
    contact_name VARCHAR(50) NOT NULL,
    contact_phone VARCHAR(20),
    contact_email VARCHAR(100),
    address TEXT,
    status customer_status NOT NULL DEFAULT 'active',
    parent_id UUID REFERENCES customers(id), -- 支持客户层级关系
    level INTEGER NOT NULL DEFAULT 1, -- 客户级别（1=直客，2=二级代理等）
    business_type VARCHAR(50), -- 业务类型
    credit_limit DECIMAL(15,2), -- 信用额度
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP
);

-- 3. 云服务商配置表
CREATE TABLE cloud_providers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    provider cloud_provider NOT NULL,
    api_endpoint VARCHAR(200) NOT NULL,
    api_version VARCHAR(20),
    is_active BOOLEAN NOT NULL DEFAULT true,
    rate_limit INTEGER DEFAULT 100, -- API调用频率限制（次/小时）
    retry_count INTEGER DEFAULT 3, -- 重试次数
    timeout_seconds INTEGER DEFAULT 30, -- 超时时间
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP
);

-- 4. 客户云平台配置表
CREATE TABLE customer_cloud_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customers(id),
    provider_id UUID NOT NULL REFERENCES cloud_providers(id),
    api_key_encrypted TEXT NOT NULL, -- 加密存储的API密钥
    secret_key_encrypted TEXT, -- 加密存储的Secret密钥
    account_id VARCHAR(100), -- 云平台账户ID
    is_active BOOLEAN NOT NULL DEFAULT true,
    sync_enabled BOOLEAN NOT NULL DEFAULT true, -- 是否启用数据同步
    last_sync_at TIMESTAMP, -- 最后同步时间
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP,
    UNIQUE(customer_id, provider_id)
);

-- 5. 合同表
CREATE TABLE contracts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_no VARCHAR(50) UNIQUE NOT NULL, -- 合同编号
    customer_id UUID NOT NULL REFERENCES customers(id),
    title VARCHAR(200) NOT NULL,
    status contract_status NOT NULL DEFAULT 'draft',
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    settlement_cycle INTEGER NOT NULL DEFAULT 1, -- 结算周期（月）
    payment_terms TEXT, -- 付款条款
    contract_amount DECIMAL(15,2), -- 合同金额
    discount_rate DECIMAL(5,4), -- 折扣率（0.8500表示85%折扣）
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP
);

-- 6. 返佣规则表
CREATE TABLE commission_rules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL REFERENCES contracts(id),
    provider cloud_provider NOT NULL,
    service_type VARCHAR(50) NOT NULL, -- 服务类型（compute, storage, network等）
    tier_min DECIMAL(15,2) NOT NULL DEFAULT 0, -- 阶梯最小值
    tier_max DECIMAL(15,2), -- 阶梯最大值（NULL表示无上限）
    commission_rate DECIMAL(5,4) NOT NULL, -- 返佣比例（0.0500表示5%）
    fixed_amount DECIMAL(15,2), -- 固定返佣金额
    is_active BOOLEAN NOT NULL DEFAULT true,
    effective_date DATE NOT NULL,
    expiry_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP
);

-- 7. 账单数据表（按月分表）
-- 主表模板
CREATE TABLE billing_data_template (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customers(id),
    provider cloud_provider NOT NULL,
    account_id VARCHAR(100) NOT NULL,
    service_type VARCHAR(50) NOT NULL,
    resource_id VARCHAR(200),
    resource_name VARCHAR(200),
    usage_amount DECIMAL(15,6) NOT NULL, -- 使用量
    usage_unit VARCHAR(20), -- 使用单位
    unit_price DECIMAL(15,6) NOT NULL, -- 单价
    original_cost DECIMAL(15,2) NOT NULL, -- 原始费用
    discounted_cost DECIMAL(15,2) NOT NULL, -- 折扣后费用
    currency VARCHAR(3) NOT NULL DEFAULT 'CNY',
    billing_date DATE NOT NULL,
    billing_period VARCHAR(10) NOT NULL, -- 计费周期（YYYY-MM）
    region VARCHAR(50),
    zone VARCHAR(50),
    tags JSONB, -- 标签信息
    raw_data JSONB, -- 原始API响应数据
    sync_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 8. 返佣记录表
CREATE TABLE commission_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    billing_data_id UUID NOT NULL, -- 关联账单数据ID
    customer_id UUID NOT NULL REFERENCES customers(id),
    contract_id UUID NOT NULL REFERENCES contracts(id),
    rule_id UUID NOT NULL REFERENCES commission_rules(id),
    provider cloud_provider NOT NULL,
    service_type VARCHAR(50) NOT NULL,
    base_amount DECIMAL(15,2) NOT NULL, -- 计算基数
    commission_rate DECIMAL(5,4) NOT NULL, -- 适用返佣比例
    commission_amount DECIMAL(15,2) NOT NULL, -- 返佣金额
    status commission_status NOT NULL DEFAULT 'pending',
    billing_period VARCHAR(10) NOT NULL, -- 计费周期（YYYY-MM）
    calculated_at TIMESTAMP,
    paid_at TIMESTAMP,
    payment_reference VARCHAR(100), -- 支付参考号
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    deleted_at TIMESTAMP
);

-- 9. 数据同步日志表
CREATE TABLE sync_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customers(id),
    provider cloud_provider NOT NULL,
    sync_type VARCHAR(50) NOT NULL, -- 同步类型（billing, usage等）
    sync_period VARCHAR(10), -- 同步周期（YYYY-MM）
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    status VARCHAR(20) NOT NULL, -- 状态（running, success, failed）
    records_count INTEGER DEFAULT 0, -- 处理记录数
    error_message TEXT, -- 错误信息
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 10. 系统配置表
CREATE TABLE system_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    config_key VARCHAR(100) UNIQUE NOT NULL,
    config_value TEXT NOT NULL,
    config_type VARCHAR(20) NOT NULL DEFAULT 'string', -- string, number, boolean, json
    description TEXT,
    is_encrypted BOOLEAN NOT NULL DEFAULT false, -- 是否加密存储
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id)
);

-- 创建索引
-- 用户表索引
CREATE INDEX idx_users_username ON users(username) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_role ON users(role);

-- 客户表索引
CREATE INDEX idx_customers_code ON customers(customer_code) WHERE deleted_at IS NULL;
CREATE INDEX idx_customers_status ON customers(status);
CREATE INDEX idx_customers_parent ON customers(parent_id);
CREATE INDEX idx_customers_level ON customers(level);

-- 合同表索引
CREATE INDEX idx_contracts_customer ON contracts(customer_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_contracts_status ON contracts(status);
CREATE INDEX idx_contracts_dates ON contracts(start_date, end_date);

-- 返佣规则表索引
CREATE INDEX idx_commission_rules_contract ON commission_rules(contract_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_commission_rules_provider ON commission_rules(provider);
CREATE INDEX idx_commission_rules_service ON commission_rules(service_type);

-- 账单数据表索引（将在分表中创建）
-- CREATE INDEX idx_billing_data_customer ON billing_data_YYYYMM(customer_id);
-- CREATE INDEX idx_billing_data_provider ON billing_data_YYYYMM(provider);
-- CREATE INDEX idx_billing_data_period ON billing_data_YYYYMM(billing_period);
-- CREATE INDEX idx_billing_data_date ON billing_data_YYYYMM(billing_date);

-- 返佣记录表索引
CREATE INDEX idx_commission_records_billing ON commission_records(billing_data_id);
CREATE INDEX idx_commission_records_customer ON commission_records(customer_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_commission_records_period ON commission_records(billing_period);
CREATE INDEX idx_commission_records_status ON commission_records(status);

-- 数据同步日志索引
CREATE INDEX idx_sync_logs_customer ON sync_logs(customer_id);
CREATE INDEX idx_sync_logs_provider ON sync_logs(provider);
CREATE INDEX idx_sync_logs_period ON sync_logs(sync_period);
CREATE INDEX idx_sync_logs_status ON sync_logs(status);

-- 系统配置表索引
CREATE INDEX idx_system_configs_key ON system_configs(config_key);

-- 创建更新时间触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 为所有需要的表创建更新时间触发器
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_customers_updated_at BEFORE UPDATE ON customers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_cloud_providers_updated_at BEFORE UPDATE ON cloud_providers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_customer_cloud_configs_updated_at BEFORE UPDATE ON customer_cloud_configs FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_contracts_updated_at BEFORE UPDATE ON contracts FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_commission_rules_updated_at BEFORE UPDATE ON commission_rules FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_commission_records_updated_at BEFORE UPDATE ON commission_records FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_system_configs_updated_at BEFORE UPDATE ON system_configs FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- 插入初始数据

-- 初始化云服务商配置
INSERT INTO cloud_providers (name, provider, api_endpoint, api_version) VALUES
('腾讯云', 'tencent', 'https://partners.tencentcloudapi.com', 'v1'),
('阿里云', 'alibaba', 'https://ecs.aliyuncs.com', '2014-05-26'),
('华为云', 'huawei', 'https://ecs.myhuaweicloud.com', 'v1'),
('AWS', 'aws', 'https://aws.amazon.com', '2016-11-15');

-- 初始化系统配置
INSERT INTO system_configs (config_key, config_value, config_type, description) VALUES
('system.name', 'XCloud多云对账平台', 'string', '系统名称'),
('system.version', '1.0.0', 'string', '系统版本'),
('billing.default_currency', 'CNY', 'string', '默认货币单位'),
('sync.default_retry_count', '3', 'number', '默认重试次数'),
('sync.default_timeout', '30', 'number', '默认超时时间（秒）'),
('commission.precision', '4', 'number', '返佣计算精度'),
('report.max_export_records', '100000', 'number', '报表导出最大记录数');

-- 创建分表函数
CREATE OR REPLACE FUNCTION create_billing_data_partition(partition_date DATE)
RETURNS VOID AS $$
DECLARE
    table_name TEXT;
    start_date DATE;
    end_date DATE;
BEGIN
    -- 计算分表名称
    table_name := 'billing_data_' || to_char(partition_date, 'YYYYMM');
    
    -- 计算分表范围
    start_date := date_trunc('month', partition_date);
    end_date := start_date + INTERVAL '1 month';
    
    -- 创建分表（如果不存在）
    EXECUTE format('
        CREATE TABLE IF NOT EXISTS %I (
            LIKE billing_data_template INCLUDING ALL,
            CHECK (billing_date >= %L AND billing_date < %L)
        ) INHERITS (billing_data_template)',
        table_name, start_date, end_date);
    
    -- 创建索引
    EXECUTE format('CREATE INDEX IF NOT EXISTS %I ON %I(customer_id)', 
        'idx_' || table_name || '_customer', table_name);
    EXECUTE format('CREATE INDEX IF NOT EXISTS %I ON %I(provider)', 
        'idx_' || table_name || '_provider', table_name);
    EXECUTE format('CREATE INDEX IF NOT EXISTS %I ON %I(billing_period)', 
        'idx_' || table_name || '_period', table_name);
    EXECUTE format('CREATE INDEX IF NOT EXISTS %I ON %I(billing_date)', 
        'idx_' || table_name || '_date', table_name);
        
    -- 创建更新时间触发器
    EXECUTE format('CREATE TRIGGER update_%I_updated_at BEFORE UPDATE ON %I FOR EACH ROW EXECUTE FUNCTION update_updated_at_column()', 
        table_name, table_name);
        
    RAISE NOTICE '分表 % 创建完成', table_name;
END;
$$ LANGUAGE plpgsql;

-- 创建当前月份和下个月的分表
SELECT create_billing_data_partition(CURRENT_DATE);
SELECT create_billing_data_partition(CURRENT_DATE + INTERVAL '1 month');

COMMENT ON DATABASE xcloud IS 'XCloud多云对账平台数据库 - 包含用户管理、客户管理、合同管理、返佣计算等核心功能的完整数据结构';