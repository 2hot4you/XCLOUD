-- XCloud数据库初始化脚本
-- 创建必要的扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 创建用户和权限
-- （生产环境中应该使用更安全的密码）
-- ALTER USER xcloud WITH SUPERUSER;

-- 设置时区
SET timezone = 'Asia/Shanghai';

-- 创建数据库（如果不存在）
-- CREATE DATABASE xcloud;

-- 添加注释
COMMENT ON DATABASE xcloud IS 'XCloud多云对账平台数据库';