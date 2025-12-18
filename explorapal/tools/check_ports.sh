#!/bin/bash

# 检查所有服务端口状态的脚本

echo "🔍 检查所有服务端口状态..."
echo ""

SERVICES=(
    "项目管理RPC服务:9001"
    "AI对话RPC服务:9002"
    "API服务:9003"
)

for service in "${SERVICES[@]}"; do
    NAME=$(echo $service | cut -d: -f1)
    PORT=$(echo $service | cut -d: -f2)

    echo "📋 $NAME (端口 $PORT):"

    # 检查端口是否被占用
    if command -v lsof &> /dev/null; then
        PID=$(lsof -ti :$PORT 2>/dev/null)
        if [ -z "$PID" ]; then
            echo "   ✅ 端口 $PORT 未被占用"
        else
            PROCESS_INFO=$(lsof -i :$PORT | head -2 | tail -1)
            echo "   ⚠️  端口 $PORT 被占用: $PROCESS_INFO"
        fi
    elif command -v netstat &> /dev/null; then
        PID=$(netstat -tlnp 2>/dev/null | grep ":$PORT " | awk '{print $7}' | cut -d'/' -f1)
        if [ -z "$PID" ]; then
            echo "   ✅ 端口 $PORT 未被占用"
        else
            echo "   ⚠️  端口 $PORT 被进程 $PID 占用"
        fi
    else
        echo "   ❓ 无法检查端口状态（未找到 lsof 或 netstat）"
    fi

    echo ""
done

echo "💡 如果端口被占用，可以使用以下命令清理："
echo "   ./tools/kill_port.sh 9001  # 清理项目管理RPC服务端口"
echo "   ./tools/kill_port.sh 9002  # 清理AI对话RPC服务端口"
echo "   ./tools/kill_port.sh 9003  # 清理API服务端口"
