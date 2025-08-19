#!/bin/bash

# XCloud 服务管理脚本
# 作者: XCloud开发团队
# 版本: 1.0

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PID_DIR="$PROJECT_ROOT/.pids"
LOG_DIR="$PROJECT_ROOT/logs"

# 确保目录存在
mkdir -p "$PID_DIR" "$LOG_DIR"

# PID文件路径
BACKEND_PID_FILE="$PID_DIR/backend.pid"
FRONTEND_PID_FILE="$PID_DIR/frontend.pid"

# 日志文件路径
BACKEND_LOG_FILE="$LOG_DIR/backend.log"
FRONTEND_LOG_FILE="$LOG_DIR/frontend.log"

# 打印彩色文本
print_color() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# 打印标题
print_title() {
    echo
    print_color $BLUE "=========================================="
    print_color $CYAN "          XCloud 服务管理工具"
    print_color $BLUE "=========================================="
    echo
}

# 检查进程是否运行
is_running() {
    local pid_file=$1
    if [ -f "$pid_file" ]; then
        local pid=$(cat "$pid_file")
        if ps -p "$pid" > /dev/null 2>&1; then
            return 0
        else
            rm -f "$pid_file"
            return 1
        fi
    fi
    return 1
}

# 获取服务状态
get_service_status() {
    local service=$1
    local pid_file=""
    
    case $service in
        "backend")
            pid_file="$BACKEND_PID_FILE"
            ;;
        "frontend")
            pid_file="$FRONTEND_PID_FILE"
            ;;
    esac
    
    if is_running "$pid_file"; then
        print_color $GREEN "●运行中"
    else
        print_color $RED "●已停止"
    fi
}

# 获取Docker服务状态
get_docker_status() {
    if command -v docker &> /dev/null && command -v docker-compose &> /dev/null; then
        if docker-compose -f "$PROJECT_ROOT/docker-compose.yml" ps 2>/dev/null | grep -q "Up"; then
            print_color $GREEN "●运行中"
        else
            print_color $RED "●已停止"
        fi
    elif command -v docker &> /dev/null; then
        # 尝试使用docker compose (新版本)
        if docker compose -f "$PROJECT_ROOT/docker-compose.yml" ps 2>/dev/null | grep -q "Up"; then
            print_color $GREEN "●运行中"
        else
            print_color $RED "●已停止"
        fi
    else
        print_color $RED "●Docker未安装"
    fi
}

# 显示服务状态
show_status() {
    print_color $YELLOW "=== 服务状态 ==="
    echo -e "后端服务:   $(get_service_status backend)"
    echo -e "前端服务:   $(get_service_status frontend)"
    echo -e "Docker服务: $(get_docker_status)"
    echo
}

# 启动后端服务
start_backend() {
    print_color $YELLOW "启动后端服务..."
    
    if is_running "$BACKEND_PID_FILE"; then
        print_color $GREEN "后端服务已在运行中"
        return 0
    fi
    
    cd "$PROJECT_ROOT/backend"
    
    # 检查Go模块
    if [ ! -f "go.mod" ]; then
        print_color $RED "错误: 找不到go.mod文件"
        return 1
    fi
    
    # 启动后端服务并记录PID
    nohup go run cmd/main.go > "$BACKEND_LOG_FILE" 2>&1 &
    local pid=$!
    echo $pid > "$BACKEND_PID_FILE"
    
    # 等待服务启动
    sleep 3
    
    if is_running "$BACKEND_PID_FILE"; then
        print_color $GREEN "✓ 后端服务启动成功 (PID: $pid)"
        print_color $CYAN "  - API地址: http://localhost:8080"
        print_color $CYAN "  - Swagger文档: http://localhost:8080/swagger/index.html"
        print_color $CYAN "  - 日志文件: $BACKEND_LOG_FILE"
    else
        print_color $RED "✗ 后端服务启动失败"
        return 1
    fi
}

# 停止后端服务
stop_backend() {
    print_color $YELLOW "停止后端服务..."
    
    if ! is_running "$BACKEND_PID_FILE"; then
        print_color $GREEN "后端服务未运行"
        return 0
    fi
    
    local pid=$(cat "$BACKEND_PID_FILE")
    kill $pid
    rm -f "$BACKEND_PID_FILE"
    
    print_color $GREEN "✓ 后端服务已停止"
}

# 启动前端服务
start_frontend() {
    print_color $YELLOW "启动前端服务..."
    
    if is_running "$FRONTEND_PID_FILE"; then
        print_color $GREEN "前端服务已在运行中"
        return 0
    fi
    
    cd "$PROJECT_ROOT/frontend"
    
    # 检查package.json
    if [ ! -f "package.json" ]; then
        print_color $RED "错误: 找不到package.json文件"
        return 1
    fi
    
    # 检查node_modules
    if [ ! -d "node_modules" ]; then
        print_color $YELLOW "安装依赖包..."
        npm install
    fi
    
    # 启动前端服务并记录PID
    nohup npm run dev > "$FRONTEND_LOG_FILE" 2>&1 &
    local pid=$!
    echo $pid > "$FRONTEND_PID_FILE"
    
    # 等待服务启动
    sleep 5
    
    if is_running "$FRONTEND_PID_FILE"; then
        print_color $GREEN "✓ 前端服务启动成功 (PID: $pid)"
        print_color $CYAN "  - 访问地址: http://localhost:3000"
        print_color $CYAN "  - 日志文件: $FRONTEND_LOG_FILE"
    else
        print_color $RED "✗ 前端服务启动失败"
        return 1
    fi
}

# 停止前端服务
stop_frontend() {
    print_color $YELLOW "停止前端服务..."
    
    if ! is_running "$FRONTEND_PID_FILE"; then
        print_color $GREEN "前端服务未运行"
        return 0
    fi
    
    local pid=$(cat "$FRONTEND_PID_FILE")
    # 杀死进程组以确保所有子进程都被停止
    pkill -P $pid 2>/dev/null || true
    kill $pid 2>/dev/null || true
    rm -f "$FRONTEND_PID_FILE"
    
    print_color $GREEN "✓ 前端服务已停止"
}

# 启动Docker服务
start_docker() {
    print_color $YELLOW "启动Docker服务..."
    
    cd "$PROJECT_ROOT"
    
    if ! command -v docker &> /dev/null; then
        print_color $RED "错误: 找不到docker命令"
        return 1
    fi
    
    # 创建数据目录
    mkdir -p docker/data/postgres docker/data/redis docker/data/rabbitmq
    
    # 启动Docker服务，优先使用docker-compose，然后尝试docker compose
    if command -v docker-compose &> /dev/null; then
        docker-compose up -d
    else
        docker compose up -d
    fi
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ Docker服务启动成功"
        print_color $CYAN "  - PostgreSQL: localhost:5432"
        print_color $CYAN "  - Redis: localhost:6379"
        print_color $CYAN "  - RabbitMQ: localhost:5672 (管理界面: localhost:15672)"
        
        print_color $YELLOW "等待数据库初始化..."
        sleep 5
    else
        print_color $RED "✗ Docker服务启动失败"
        return 1
    fi
}

# 停止Docker服务
stop_docker() {
    print_color $YELLOW "停止Docker服务..."
    
    cd "$PROJECT_ROOT"
    
    # 停止Docker服务，优先使用docker-compose，然后尝试docker compose
    if command -v docker-compose &> /dev/null; then
        docker-compose down
    else
        docker compose down
    fi
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ Docker服务已停止"
    else
        print_color $RED "✗ Docker服务停止失败"
        return 1
    fi
}

# 重启服务
restart_service() {
    local service=$1
    case $service in
        "backend")
            stop_backend
            sleep 2
            start_backend
            ;;
        "frontend")
            stop_frontend
            sleep 2
            start_frontend
            ;;
        "docker")
            stop_docker
            sleep 2
            start_docker
            ;;
    esac
}

# 查看日志
view_logs() {
    local service=$1
    local log_file=""
    
    case $service in
        "backend")
            log_file="$BACKEND_LOG_FILE"
            ;;
        "frontend")
            log_file="$FRONTEND_LOG_FILE"
            ;;
        "docker")
            cd "$PROJECT_ROOT"
            if command -v docker-compose &> /dev/null; then
                docker-compose logs -f
            else
                docker compose logs -f
            fi
            return
            ;;
    esac
    
    if [ -f "$log_file" ]; then
        print_color $CYAN "查看 $service 日志 (按Ctrl+C退出):"
        echo "----------------------------------------"
        tail -f "$log_file"
    else
        print_color $RED "日志文件不存在: $log_file"
    fi
}

# 构建项目
build_project() {
    print_color $YELLOW "构建项目..."
    
    cd "$PROJECT_ROOT"
    
    print_color $CYAN "构建后端..."
    cd backend && go build -o bin/xcloud ./cmd
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ 后端构建成功"
    else
        print_color $RED "✗ 后端构建失败"
        return 1
    fi
    
    print_color $CYAN "构建前端..."
    cd ../frontend && npm run build
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ 前端构建成功"
    else
        print_color $RED "✗ 前端构建失败"
        return 1
    fi
    
    print_color $GREEN "✓ 项目构建完成"
}

# 运行测试
run_tests() {
    print_color $YELLOW "运行测试..."
    
    cd "$PROJECT_ROOT"
    
    print_color $CYAN "运行后端测试..."
    cd backend && go test ./...
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ 后端测试通过"
    else
        print_color $RED "✗ 后端测试失败"
        return 1
    fi
    
    print_color $CYAN "运行前端测试..."
    cd ../frontend && npm run test
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ 前端测试通过"
    else
        print_color $RED "✗ 前端测试失败"
        return 1
    fi
    
    print_color $GREEN "✓ 所有测试通过"
}

# 数据库迁移
migrate_database() {
    print_color $YELLOW "运行数据库迁移..."
    
    cd "$PROJECT_ROOT/backend"
    go run cmd/migrate.go up
    
    if [ $? -eq 0 ]; then
        print_color $GREEN "✓ 数据库迁移完成"
    else
        print_color $RED "✗ 数据库迁移失败"
        return 1
    fi
}

# 清理项目
clean_project() {
    print_color $YELLOW "清理项目文件..."
    
    # 停止所有服务
    stop_backend
    stop_frontend
    
    # 清理构建文件
    rm -rf "$PROJECT_ROOT/backend/bin/"
    rm -rf "$PROJECT_ROOT/frontend/dist/"
    
    # 清理PID和日志文件
    rm -rf "$PID_DIR"
    rm -rf "$LOG_DIR"
    
    print_color $GREEN "✓ 项目清理完成"
}

# 显示主菜单
show_menu() {
    clear
    print_title
    show_status
    
    print_color $YELLOW "=== 服务管理 ==="
    echo "1.  启动后端服务"
    echo "2.  停止后端服务"
    echo "3.  重启后端服务"
    echo
    echo "4.  启动前端服务"
    echo "5.  停止前端服务"
    echo "6.  重启前端服务"
    echo
    echo "7.  启动Docker服务"
    echo "8.  停止Docker服务"
    echo "9.  重启Docker服务"
    echo
    print_color $YELLOW "=== 日志查看 ==="
    echo "10. 查看后端日志"
    echo "11. 查看前端日志"
    echo "12. 查看Docker日志"
    echo
    print_color $YELLOW "=== 开发工具 ==="
    echo "13. 构建项目"
    echo "14. 运行测试"
    echo "15. 数据库迁移"
    echo "16. 清理项目"
    echo
    print_color $YELLOW "=== 快捷操作 ==="
    echo "17. 启动完整开发环境"
    echo "18. 停止所有服务"
    echo
    echo "0.  退出"
    echo
    print_color $GREEN "请选择操作 (0-18): "
}

# 启动完整开发环境
start_dev_env() {
    print_color $YELLOW "启动完整开发环境..."
    start_docker
    sleep 3
    start_backend
    sleep 2
    start_frontend
    print_color $GREEN "✓ 开发环境启动完成!"
}

# 停止所有服务
stop_all_services() {
    print_color $YELLOW "停止所有服务..."
    stop_frontend
    stop_backend
    stop_docker
    print_color $GREEN "✓ 所有服务已停止"
}

# 主循环
main() {
    while true; do
        show_menu
        read -r choice
        
        case $choice in
            1) start_backend; read -p "按Enter继续..." ;;
            2) stop_backend; read -p "按Enter继续..." ;;
            3) restart_service backend; read -p "按Enter继续..." ;;
            4) start_frontend; read -p "按Enter继续..." ;;
            5) stop_frontend; read -p "按Enter继续..." ;;
            6) restart_service frontend; read -p "按Enter继续..." ;;
            7) start_docker; read -p "按Enter继续..." ;;
            8) stop_docker; read -p "按Enter继续..." ;;
            9) restart_service docker; read -p "按Enter继续..." ;;
            10) view_logs backend ;;
            11) view_logs frontend ;;
            12) view_logs docker ;;
            13) build_project; read -p "按Enter继续..." ;;
            14) run_tests; read -p "按Enter继续..." ;;
            15) migrate_database; read -p "按Enter继续..." ;;
            16) clean_project; read -p "按Enter继续..." ;;
            17) start_dev_env; read -p "按Enter继续..." ;;
            18) stop_all_services; read -p "按Enter继续..." ;;
            0) 
                print_color $GREEN "感谢使用XCloud服务管理工具!"
                exit 0
                ;;
            *)
                print_color $RED "无效选择，请重新输入"
                read -p "按Enter继续..."
                ;;
        esac
    done
}

# 脚本入口
if [ $# -eq 0 ]; then
    main
else
    # 支持命令行参数
    case $1 in
        "start-backend") start_backend ;;
        "stop-backend") stop_backend ;;
        "restart-backend") restart_service backend ;;
        "start-frontend") start_frontend ;;
        "stop-frontend") stop_frontend ;;
        "restart-frontend") restart_service frontend ;;
        "start-docker") start_docker ;;
        "stop-docker") stop_docker ;;
        "restart-docker") restart_service docker ;;
        "start-dev") start_dev_env ;;
        "stop-all") stop_all_services ;;
        "status") show_status ;;
        "build") build_project ;;
        "test") run_tests ;;
        "migrate") migrate_database ;;
        "clean") clean_project ;;
        "logs-backend") view_logs backend ;;
        "logs-frontend") view_logs frontend ;;
        "logs-docker") view_logs docker ;;
        *)
            echo "用法: $0 [命令]"
            echo "可用命令:"
            echo "  start-backend, stop-backend, restart-backend"
            echo "  start-frontend, stop-frontend, restart-frontend"
            echo "  start-docker, stop-docker, restart-docker"
            echo "  start-dev, stop-all, status"
            echo "  build, test, migrate, clean"
            echo "  logs-backend, logs-frontend, logs-docker"
            ;;
    esac
fi