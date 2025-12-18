#!/bin/bash

# 停止MVP演示环境

echo "🛑 停止探索伙伴 MVP 演示环境..."

# 如果存在PID文件，读取并停止进程
if [ -f ".demo_pids" ]; then
    PIDS=$(cat .demo_pids)
    if [ ! -z "$PIDS" ]; then
        echo "停止服务进程: $PIDS"
        kill $PIDS 2>/dev/null
        rm .demo_pids
        echo "✅ 服务已停止"
    else
        echo "⚠️  未找到运行中的服务进程"
    fi
else
    # 查找可能的服务进程
    echo "查找可能的服务进程..."
    PROJECT_PID=$(pgrep -f "projectmanagementservice" | head -1)
    AI_PID=$(pgrep -f "aidialogueservice" | head -1)
    API_PID=$(pgrep -f "api.go" | head -1)

    if [ ! -z "$PROJECT_PID" ] || [ ! -z "$AI_PID" ] || [ ! -z "$API_PID" ]; then
        echo "发现运行中的服务，正在停止..."
        [ ! -z "$PROJECT_PID" ] && kill $PROJECT_PID && echo "停止项目管理RPC服务 (PID: $PROJECT_PID)"
        [ ! -z "$AI_PID" ] && kill $AI_PID && echo "停止AI对话RPC服务 (PID: $AI_PID)"
        [ ! -z "$API_PID" ] && kill $API_PID && echo "停止API服务 (PID: $API_PID)"
        echo "✅ 所有服务已停止"
    else
        echo "⚠️  未发现运行中的演示服务"
    fi
fi

echo ""
echo "🎯 如需重启演示，请运行:"
echo "  ./start_demo.sh"
