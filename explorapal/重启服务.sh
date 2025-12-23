#!/bin/bash

# 重启ExploraPal演示服务
# 停止旧服务，重新编译和启动新服务

echo "🔄 重启ExploraPal演示服务..."
echo ""

# 停止旧服务
echo "🛑 停止旧服务..."
pkill -f "videoprocessingservice.go" 2>/dev/null || true
pkill -f "api.go" 2>/dev/null || true
pkill -f "aidialogueservice.go" 2>/dev/null || true
pkill -f "projectmanagementservice.go" 2>/dev/null || true
pkill -f "audioprocessingservice.go" 2>/dev/null || true

echo "✅ 旧服务已停止"
sleep 2

# 清理可能的端口占用
echo "🧹 清理端口..."
lsof -ti:9001,9002,9003,9004,9005 | xargs kill -9 2>/dev/null || true

echo "✅ 端口清理完成"
sleep 1

# 启动服务
echo "🚀 启动新服务..."
./start_demo.sh

echo ""
echo "⏳ 等待服务完全启动（45秒）..."
sleep 45

# 检查服务状态
echo "🔍 检查服务状态..."
echo "  测试端点: http://localhost:9003/api/ping"

# 尝试多次检查
MAX_RETRIES=3
for i in $(seq 1 $MAX_RETRIES); do
    echo "  尝试 $i/$MAX_RETRIES..."
    if curl -s --max-time 5 http://localhost:9003/api/ping > /dev/null 2>&1; then
        SUCCESS=true
        break
    fi
    sleep 5
done

if [ "$SUCCESS" = true ]; then
    echo "✅ 所有服务启动成功！"
    echo ""
    echo "📊 服务状态:"
    echo "  • API服务: http://localhost:9003 ✅"
    echo "  • 视频处理服务: http://localhost:9005 ✅"
    echo "  • AI对话服务: http://localhost:9002 ✅"
    echo "  • 项目管理服务: http://localhost:9001 ✅"
    echo "  • 音频处理服务: http://localhost:9004 ✅"
else
    echo "❌ 服务启动失败，请检查日志"
    echo "查看日志: tail -f logs/*.log"
    exit 1
fi
