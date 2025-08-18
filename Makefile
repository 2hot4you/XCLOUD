.PHONY: help dev build test clean docker-up docker-down

# 默认目标
help:
	@echo "XCloud 多云对账平台 - 可用命令："
	@echo "  dev          - 启动开发环境（数据库+后端+前端）"
	@echo "  build        - 构建前后端应用"
	@echo "  test         - 运行所有测试"
	@echo "  clean        - 清理构建文件"
	@echo "  docker-up    - 启动Docker服务（数据库等）"
	@echo "  docker-down  - 停止Docker服务"
	@echo "  migrate-up   - 运行数据库迁移"
	@echo "  migrate-down - 回滚数据库迁移"

# 启动开发环境
dev: docker-up
	@echo "启动后端服务..."
	@cd backend && go run cmd/main.go &
	@echo "启动前端服务..."
	@cd frontend && npm run dev &
	@echo "开发环境启动完成！"

# 构建应用
build:
	@echo "构建后端..."
	@cd backend && go build -o bin/xcloud ./cmd
	@echo "构建前端..."
	@cd frontend && npm run build

# 运行测试
test:
	@echo "运行后端测试..."
	@cd backend && go test ./...
	@echo "运行前端测试..."
	@cd frontend && npm run test

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -rf backend/bin/
	@rm -rf frontend/dist/
	@rm -rf frontend/node_modules/

# 启动Docker服务
docker-up:
	@echo "启动Docker服务..."
	@mkdir -p docker/data/postgres docker/data/redis docker/data/rabbitmq
	@docker-compose up -d
	@echo "等待数据库启动..."
	@sleep 5

# 停止Docker服务
docker-down:
	@echo "停止Docker服务..."
	@docker-compose down

# 数据库迁移
migrate-up:
	@echo "运行数据库迁移..."
	@cd backend && go run cmd/migrate.go up

migrate-down:
	@echo "回滚数据库迁移..."
	@cd backend && go run cmd/migrate.go down