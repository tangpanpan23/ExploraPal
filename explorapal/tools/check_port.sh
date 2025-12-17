#!/bin/bash

# 检查端口占用脚本

PORT=${1:-8082}

echo "🔍 检查端口 $PORT 占用情况..."

# macOS/Linux 通用方法
if command -v lsof &> /dev/null; then
    PID=$(lsof -ti :$PORT)
    if [ -z "$PID" ]; then
        echo "✅ 端口 $PORT 未被占用"
        exit 0
    else
        echo "⚠️  端口 $PORT 被进程占用:"
        lsof -i :$PORT
        echo ""
        read -p "是否要终止占用端口的进程? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            kill -9 $PID
            echo "✅ 已终止进程 $PID"
        else
            echo "❌ 未终止进程，请手动处理或使用其他端口"
            exit 1
        fi
    fi
elif command -v netstat &> /dev/null; then
    # Linux 备用方法
    PID=$(netstat -tlnp 2>/dev/null | grep :$PORT | awk '{print $7}' | cut -d'/' -f1)
    if [ -z "$PID" ]; then
        echo "✅ 端口 $PORT 未被占用"
        exit 0
    else
        echo "⚠️  端口 $PORT 被进程占用: PID $PID"
        read -p "是否要终止占用端口的进程? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            kill -9 $PID
            echo "✅ 已终止进程 $PID"
        else
            echo "❌ 未终止进程，请手动处理或使用其他端口"
            exit 1
        fi
    fi
else
    echo "❌ 未找到 lsof 或 netstat 命令，无法检查端口占用"
    exit 1
fi
