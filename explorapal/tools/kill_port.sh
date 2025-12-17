#!/bin/bash

# 强制终止占用指定端口的进程

PORT=${1:-8082}

echo "🔍 查找占用端口 $PORT 的进程..."

if command -v lsof &> /dev/null; then
    PIDS=$(lsof -ti :$PORT)
    if [ -z "$PIDS" ]; then
        echo "✅ 端口 $PORT 未被占用"
        exit 0
    fi
    
    echo "⚠️  发现以下进程占用端口 $PORT:"
    lsof -i :$PORT
    
    for PID in $PIDS; do
        echo "🛑 终止进程 $PID..."
        kill -9 $PID 2>/dev/null
        if [ $? -eq 0 ]; then
            echo "✅ 已终止进程 $PID"
        else
            echo "❌ 无法终止进程 $PID (可能需要sudo权限)"
        fi
    done
    
    # 再次检查
    sleep 1
    REMAINING=$(lsof -ti :$PORT)
    if [ -z "$REMAINING" ]; then
        echo "✅ 端口 $PORT 已释放"
    else
        echo "⚠️  仍有进程占用端口，可能需要sudo权限:"
        echo "   sudo kill -9 $REMAINING"
        exit 1
    fi
else
    echo "❌ 未找到 lsof 命令，请安装: brew install lsof (macOS) 或 apt-get install lsof (Linux)"
    exit 1
fi
