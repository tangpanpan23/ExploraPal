#!/bin/bash

# 测试API服务的简单脚本

echo "🔍 测试API服务..."

# API服务地址
API_HOST="localhost"
API_PORT="9003"
API_BASE_URL="http://$API_HOST:$API_PORT"

# 检查API服务是否运行
echo "📡 检查API服务状态..."
curl -s "$API_BASE_URL/api/common/ping" > /dev/null
if [ $? -ne 0 ]; then
    echo "❌ API服务未运行或不可访问"
    echo "请确保API服务正在运行: go run app/api/api.go"
    exit 1
fi

echo "✅ API服务正在运行"

# 测试项目创建API
echo ""
echo "📝 测试项目创建API..."
RESPONSE=$(curl -s -X POST "$API_BASE_URL/api/project/create" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "title": "测试项目",
    "description": "这是一个测试项目",
    "category": "dinosaur",
    "tags": ["test", "api"]
  }')

echo "响应: $RESPONSE"

# 检查响应是否成功
if echo "$RESPONSE" | grep -q "project_id"; then
    echo "✅ API调用成功！"
else
    echo "❌ API调用失败"
    echo "请检查:"
    echo "  1. 项目管理RPC服务是否在端口9001运行"
    echo "  2. API服务是否正确加载了路由"
    echo "  3. 查看API服务日志了解详细错误"
fi
