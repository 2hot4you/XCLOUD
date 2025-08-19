# XCloud 服务管理脚本使用指南

## 概述

`manage.sh` 是 XCloud 项目的一站式服务管理工具，提供了启动、停止、重启前后端服务以及 Docker 基础设施的完整功能。

## 特性

- 🎨 **彩色交互界面** - 用户友好的菜单操作
- 📊 **实时状态监控** - 实时显示各服务运行状态
- 🔧 **PID 管理** - 智能进程管理，避免重复启动
- 📝 **日志管理** - 统一的日志查看和管理
- 🚀 **一键部署** - 支持一键启动完整开发环境
- 📱 **命令行支持** - 支持直接命令行调用

## 快速开始

### 交互式菜单

```bash
# 启动交互式菜单
./scripts/manage.sh
```

### 命令行使用

```bash
# 查看服务状态
./scripts/manage.sh status

# 启动完整开发环境 (Docker + 后端 + 前端)
./scripts/manage.sh start-dev

# 停止所有服务
./scripts/manage.sh stop-all
```

## 主要功能

### 服务管理
- **后端服务** - 启动/停止/重启 Go 后端服务
- **前端服务** - 启动/停止/重启 Vue 前端服务  
- **Docker 服务** - 管理 PostgreSQL、Redis、RabbitMQ

### 日志查看
- **实时日志** - 支持查看后端、前端、Docker 服务日志
- **日志文件** - 自动保存服务日志到 `logs/` 目录

### 开发工具
- **项目构建** - 一键构建前后端项目
- **测试运行** - 运行全套测试用例
- **数据库迁移** - 执行数据库结构更新
- **项目清理** - 清理构建文件和缓存

## 可用命令

### 服务管理命令
```bash
./scripts/manage.sh start-backend      # 启动后端服务
./scripts/manage.sh stop-backend       # 停止后端服务
./scripts/manage.sh restart-backend    # 重启后端服务

./scripts/manage.sh start-frontend     # 启动前端服务
./scripts/manage.sh stop-frontend      # 停止前端服务
./scripts/manage.sh restart-frontend   # 重启前端服务

./scripts/manage.sh start-docker       # 启动Docker服务
./scripts/manage.sh stop-docker        # 停止Docker服务
./scripts/manage.sh restart-docker     # 重启Docker服务
```

### 快捷操作
```bash
./scripts/manage.sh start-dev          # 启动完整开发环境
./scripts/manage.sh stop-all           # 停止所有服务
./scripts/manage.sh status             # 查看服务状态
```

### 日志查看
```bash
./scripts/manage.sh logs-backend       # 查看后端日志
./scripts/manage.sh logs-frontend      # 查看前端日志
./scripts/manage.sh logs-docker        # 查看Docker日志
```

### 开发工具
```bash
./scripts/manage.sh build              # 构建项目
./scripts/manage.sh test               # 运行测试
./scripts/manage.sh migrate            # 数据库迁移
./scripts/manage.sh clean              # 清理项目
```

## 文件结构

脚本运行后会创建以下目录结构：

```
XCloud/
├── .pids/                 # 进程PID文件
│   ├── backend.pid
│   └── frontend.pid
├── logs/                  # 服务日志文件
│   ├── backend.log
│   └── frontend.log
└── scripts/
    ├── manage.sh          # 主管理脚本
    └── README.md          # 使用指南
```

## 服务端口

- **前端开发服务器**: http://localhost:3000
- **后端API服务**: http://localhost:8080
- **Swagger文档**: http://localhost:8080/swagger/index.html
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379
- **RabbitMQ**: localhost:5672 (管理界面: localhost:15672)

## 典型工作流程

### 初次开发环境搭建
```bash
# 1. 启动完整开发环境
./scripts/manage.sh start-dev

# 2. 等待服务启动完成，然后访问
# 前端: http://localhost:3000
# 后端API: http://localhost:8080
# Swagger文档: http://localhost:8080/swagger/index.html
```

### 日常开发
```bash
# 查看服务状态
./scripts/manage.sh status

# 重启后端服务 (代码更新后)
./scripts/manage.sh restart-backend

# 查看后端日志
./scripts/manage.sh logs-backend

# 运行测试
./scripts/manage.sh test

# 停止所有服务
./scripts/manage.sh stop-all
```

### 构建和部署
```bash
# 构建项目
./scripts/manage.sh build

# 运行完整测试
./scripts/manage.sh test

# 数据库迁移
./scripts/manage.sh migrate
```

## 故障排除

### 常见问题

1. **端口占用**
   ```bash
   # 检查端口占用
   lsof -i :3000  # 前端端口
   lsof -i :8080  # 后端端口
   ```

2. **Docker 服务启动失败**
   ```bash
   # 检查Docker状态
   docker ps
   docker-compose ps
   
   # 查看Docker日志
   ./scripts/manage.sh logs-docker
   ```

3. **后端服务启动失败**
   ```bash
   # 检查Go环境
   go version
   
   # 查看后端日志
   ./scripts/manage.sh logs-backend
   ```

4. **前端服务启动失败**
   ```bash
   # 检查Node环境
   node --version
   npm --version
   
   # 重新安装依赖
   cd frontend && npm install
   ```

### 日志查看

所有服务日志都保存在 `logs/` 目录下：
- `logs/backend.log` - 后端服务日志
- `logs/frontend.log` - 前端服务日志

使用 `tail -f logs/service.log` 可以实时查看日志更新。

## 高级用法

### 自定义配置

可以通过环境变量自定义配置：

```bash
# 自定义日志目录
export XCLOUD_LOG_DIR="/custom/logs"

# 自定义PID目录  
export XCLOUD_PID_DIR="/custom/pids"

./scripts/manage.sh start-dev
```

### 集成到其他工具

脚本支持返回标准退出码，可以集成到其他构建工具中：

```bash
#!/bin/bash
./scripts/manage.sh start-dev
if [ $? -eq 0 ]; then
    echo "开发环境启动成功"
else
    echo "开发环境启动失败"
    exit 1
fi
```

## 注意事项

1. **权限**: 确保脚本有执行权限 (`chmod +x scripts/manage.sh`)
2. **依赖**: 需要安装 Go、Node.js、Docker 环境
3. **端口**: 确保相关端口没有被其他服务占用
4. **内存**: Docker 服务需要足够的内存资源

## 支持与反馈

如果在使用过程中遇到问题，请：
1. 查看相关服务日志
2. 检查系统环境和依赖
3. 提交 Issue 或联系开发团队