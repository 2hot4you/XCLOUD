# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

本文件为Claude Code在此代码库中工作时提供指导，任何问答都使用中文形式。

## 项目概述

XCloud是为公有云代理商设计的多云对账和返佣平台。集成腾讯云、阿里云、华为云和AWS等云服务商，实现自动化对账和返佣计算。

## 系统架构

采用微服务架构：
- **后端**: 基于Go的微服务，使用Gin框架
- **前端**: Vue 3 + TypeScript + Vue DevUI组件库
- **数据库**: PostgreSQL（主库）+ Redis（缓存）
- **消息队列**: RabbitMQ用于数据同步
- **外部API**: 多个云服务商合作伙伴API

## 常用开发命令

### 项目启动
```bash
# 启动完整开发环境（数据库+后端+前端）
make dev

# 仅启动Docker服务（数据库、Redis、RabbitMQ）
make docker-up

# 停止Docker服务
make docker-down
```

### 后端开发 (Go)
```bash
# 在backend目录下：
# 启动后端服务
go run cmd/main.go

# 运行所有测试
go test ./...

# 构建后端
go build -o bin/xcloud ./cmd

# 数据库迁移
go run cmd/migrate.go up
go run cmd/migrate.go down

# 生成Swagger文档
swag init
```

### 前端开发 (Vue 3)
```bash
# 在frontend目录下：
# 启动开发服务器（端口3000）
npm run dev

# 生产构建
npm run build

# 运行测试
npm run test

# 代码检查和修复
npm run lint

# 类型检查
npm run type-check
```

### 构建和测试
```bash
# 构建前后端应用
make build

# 运行所有测试
make test

# 清理构建文件
make clean
```

## 技术栈详情

### 前端技术栈（严格遵循）
- **框架**: Vue 3.3+ + TypeScript
- **UI组件库**: Vue DevUI 1.6+ （华为云开源组件库）
- **状态管理**: Pinia 2.1+
- **路由**: Vue Router 4.2+
- **图表库**: ECharts 5.4+ + vue-echarts 6.6+
- **构建工具**: Vite 4.4+
- **HTTP客户端**: Axios 1.5+
- **工具库**: dayjs 1.11+（日期处理）

### 后端技术栈（严格遵循）
- **语言**: Go 1.23+
- **Web框架**: Gin 1.9+
- **ORM**: GORM 1.25+
- **数据库**: PostgreSQL 14+ + Redis 6+
- **消息队列**: RabbitMQ 3+
- **API文档**: Swagger/Swaggo 1.16+
- **认证**: JWT (golang-jwt/jwt/v5 5.3+)
- **配置管理**: Viper 1.16+
- **日志**: Logrus 1.9+

## 项目结构和关键文件

### 后端架构 (backend/)
- `cmd/main.go` - 应用程序入口点，包含服务器启动逻辑和路由配置
- `internal/` - 内部业务逻辑模块：
  - `auth/` - 认证和授权处理
  - `customer/` - 客户管理业务逻辑
  - `contract/` - 合同管理业务逻辑
  - `user/` - 用户管理业务逻辑
  - `commission/` - 返佣计算服务
  - `sync/` - 数据同步服务
  - `report/` - 报表统计服务
- `pkg/` - 共享包：
  - `database/` - 数据库连接和Redis配置
  - `middleware/` - 中间件（CORS、JWT、日志、恢复）
  - `jwt/` - JWT token处理
  - `logger/` - 日志配置
- `configs/config.yaml` - 应用配置文件
- `docs/` - Swagger API文档

### 前端架构 (frontend/src/)
- `main.ts` - Vue应用入口点
- `App.vue` - 根组件
- `api/` - API调用封装（Axios配置和类型定义）
- `components/` - 可复用组件
- `views/` - 页面组件：
  - `layout/MainLayout.vue` - 主布局组件
  - `dashboard/` - 仪表板页面
  - `customer/` - 客户管理页面
  - `contract/` - 合同管理页面
  - `auth/` - 认证相关页面
- `store/` - Pinia状态管理
- `router/` - Vue Router路由配置

### 数据库设计特点
- 使用UUID作为主键
- 软删除机制（deleted_at字段）
- 审计字段（created_at, updated_at, created_by, updated_by）
- 账单数据按月分表（billing_data_YYYYMM）
- 支持多云服务商的统一数据模型
- 灵活的返佣规则配置

## 部署配置

### Docker服务
- **PostgreSQL**: 端口5432，数据库名xcloud
- **Redis**: 端口6379，用于缓存
- **RabbitMQ**: 端口5672（AMQP），15672（管理界面）

### 服务端口
- **后端API**: 8080端口
- **前端开发**: 3000端口
- **Swagger文档**: http://localhost:8080/swagger/index.html

## 开发注意事项

### API集成模式
每个云服务商实现统一接口的适配器模式，处理：
- 云服务商API认证和密钥管理
- 频率限制和重试逻辑
- 数据格式标准化
- 错误处理和日志记录

### 安全考虑
- API密钥加密存储
- JWT token认证
- 审计日志记录
- 敏感数据加密

### 性能优化
- 数据库索引优化
- Redis缓存热点数据
- 按月分表策略处理大数据量
- 异步数据同步处理

## 配置文件说明

### 后端配置 (backend/configs/config.yaml)
主要配置项包括：
- 应用基本信息（名称、版本、运行模式）
- 数据库连接配置（PostgreSQL、Redis）
- JWT认证配置（密钥、过期时间）
- 云平台API配置（各个云服务商的端点和参数）
- 消息队列配置（RabbitMQ）
- 数据同步配置（同步间隔、批处理大小）

### 前端配置 (frontend/)
- `vite.config.ts` - Vite构建配置，包含代理设置
- `tsconfig.json` - TypeScript配置
- `package.json` - 依赖管理和脚本定义

## 测试和检查

### 后端测试
```bash
cd backend
go test ./... -v        # 详细测试输出
go test -cover ./...    # 测试覆盖率
```

### 前端测试和检查
```bash
cd frontend
npm run test           # 运行单元测试
npm run type-check     # TypeScript类型检查
npm run lint          # ESLint代码检查
```

## API文档访问

启动后端服务后，可以通过以下地址访问API文档：
- Swagger UI: http://localhost:8080/swagger/index.html
- 健康检查: http://localhost:8080/health

## 常见开发任务

### 添加新的云服务商支持
1. 在`internal/`下创建对应的适配器模块
2. 实现统一的API接口
3. 更新配置文件添加新的云服务商配置
4. 添加相应的数据库枚举值

### 添加新的API端点
1. 在对应的`internal/模块`下添加handler
2. 在`routes.go`中注册新路由
3. 添加Swagger注释生成API文档
4. 编写单元测试

### 数据库迁移
```bash
# 创建新的分表
SELECT create_billing_data_partition('2024-12-01'::DATE);

# 手动运行迁移
cd backend && go run cmd/migrate.go up
```