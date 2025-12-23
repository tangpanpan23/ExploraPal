#!/bin/bash

# 一键启动MVP演示环境

echo "🚀 启动探索伙伴 MVP 演示环境..."
echo ""

# 创建日志目录
echo "📁 创建日志目录..."
mkdir -p logs
echo "✅ 日志目录创建完成"

# 检查数据库
echo "📋 检查数据库连接..."
if ! mysql -u root -ptangpanpan314 -e "SELECT 1;" > /dev/null 2>&1; then
    echo "❌ 数据库连接失败，请先运行: ./setup_database.sh"
    exit 1
fi
echo "✅ 数据库连接正常"

# 清理可能占用的端口
echo ""
echo "🧹 清理端口..."
./tools/kill_port.sh 9001 > /dev/null 2>&1
./tools/kill_port.sh 9002 > /dev/null 2>&1
./tools/kill_port.sh 9003 > /dev/null 2>&1
./tools/kill_port.sh 9004 > /dev/null 2>&1
./tools/kill_port.sh 9005 > /dev/null 2>&1
echo "✅ 端口清理完成"

# 启动服务
echo ""
echo "🔄 启动服务..."

# 启动项目管理RPC服务
echo "启动项目管理RPC服务 (端口9001)..."
cd app/project-management/rpc
go run projectmanagementservice.go > ../../../logs/project-rpc.log 2>&1 &
PROJECT_PID=$!
cd ../../..

# 等待一下让服务启动
sleep 2

# 启动AI对话RPC服务
echo "启动AI对话RPC服务 (端口9002)..."
cd app/ai-dialogue/rpc
go run aidialogueservice.go > ../../../logs/ai-rpc.log 2>&1 &
AI_PID=$!
cd ../../..

# 等待一下让服务启动
sleep 2

# 启动语音处理RPC服务
echo "启动语音处理RPC服务 (端口9004)..."
cd app/audio-processing/rpc
go run audioprocessingservice.go > ../../../logs/audio-rpc.log 2>&1 &
AUDIO_PID=$!
cd ../../..

# 等待一下让服务启动
sleep 2

# 启动视频处理RPC服务
echo "启动视频处理RPC服务 (端口9005)..."
cd app/video-processing/rpc
go run videoprocessingservice.go > ../../../logs/video-rpc.log 2>&1 &
VIDEO_PID=$!
cd ../../..

# 等待一下让服务启动
sleep 2

# 启动API服务
echo "启动API服务 (端口9003)..."
cd app/api
go run api.go > ../../logs/api.log 2>&1 &
API_PID=$!
cd ..

# 等待服务完全启动
sleep 3

echo ""
echo "🎉 服务启动完成！"
echo ""
echo "📊 服务状态:"
echo "  项目管理RPC: http://localhost:9001 (PID: $PROJECT_PID)"
echo "  AI对话RPC:    http://localhost:9002 (PID: $AI_PID)"
echo "  语音处理RPC:  http://localhost:9004 (PID: $AUDIO_PID)"
echo "  视频处理RPC:  http://localhost:9005 (PID: $VIDEO_PID)"
echo "  API服务:      http://localhost:9003 (PID: $API_PID)"
echo ""
echo "🧪 测试API:"
echo "  curl http://localhost:9003/api/common/ping"
echo ""
echo "📖 查看演示文档:"
echo "  cat MVP_DEMO.md"
echo ""
echo "🛑 停止服务:"
echo "  kill $PROJECT_PID $AI_PID $AUDIO_PID $VIDEO_PID $API_PID"
echo ""

# 保存进程ID到文件
echo "$PROJECT_PID $AI_PID $AUDIO_PID $VIDEO_PID $API_PID" > .demo_pids

echo "💡 提示: 运行演示时保持这个终端窗口开启"
echo "   或者使用 Ctrl+C 停止所有服务"
