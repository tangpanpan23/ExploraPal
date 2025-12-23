#!/bin/bash

# 测试豆包异步API连接的脚本

echo "🧪 测试豆包异步API连接"
echo "========================"

# API配置 - 使用测试环境
ASYNC_API_URL="http://apx-api.tal.com/v1/async/chat"
API_KEY="300000712:9ffb0776d5409f4131f0a314fd5cb80e"  # 请替换为真实的API密钥

echo "🔗 API端点: $ASYNC_API_URL"
echo "🔑 API密钥: $API_KEY"
echo ""

# 测试1: 检查API端点是否可达
echo "测试1: 检查API端点连通性"
echo "curl -I $ASYNC_API_URL"
CURL_TEST=$(curl -I -s "$ASYNC_API_URL" 2>/dev/null | head -1)
echo "响应: $CURL_TEST"
echo ""

# 测试2: 发送最小化请求 (不包含图片)
echo "测试2: 发送最小化测试请求"
TEST_JSON=$(mktemp /tmp/test_request_XXXXXX.json)
cat > "$TEST_JSON" << EOF
{
  "model": "doubao-seedance-1.0-lite-t2v",
  "prompt": "生成一个蓝天白云的美丽风景视频",
  "duration": "5"
}
EOF

echo "📄 测试请求内容:"
cat "$TEST_JSON"
echo ""

echo "🚀 发送测试请求..."
TEST_RESPONSE=$(curl -s -w "\nHTTP_STATUS:%{http_code}\n" -X POST "$ASYNC_API_URL" \
  -H "Content-Type: application/json" \
  -H "api-key: $API_KEY" \
  -H "X-APX-Model: doubao-seedance-1.0-lite-t2v" \
  -d @"$TEST_JSON")

HTTP_STATUS=$(echo "$TEST_RESPONSE" | grep "HTTP_STATUS:" | cut -d: -f2)
ACTUAL_RESPONSE=$(echo "$TEST_RESPONSE" | grep -v "HTTP_STATUS:")

echo "📊 HTTP状态码: $HTTP_STATUS"
echo "📄 响应内容: $ACTUAL_RESPONSE"
echo ""

# 清理
rm -f "$TEST_JSON"

echo "测试完成"
