#!/bin/bash

# 配置测试脚本
# 用于验证配置文件和API密钥读取是否正常

echo "🧪 测试视频生成配置"
echo "===================="

CONFIG_FILE="video_generation_config.yaml"

# 检查配置文件是否存在
if [ ! -f "$CONFIG_FILE" ]; then
    echo "❌ 配置文件 '$CONFIG_FILE' 不存在"
    echo "请复制 video_generation_config.yaml.example 为 $CONFIG_FILE 并填入真实配置"
    exit 1
fi

echo "✅ 配置文件存在: $CONFIG_FILE"

# 测试各个配置项
echo ""
echo "🔍 测试配置读取:"

echo -n "  API Key: "
API_KEY="$(./read_config.sh api_key)"
if [ -n "$API_KEY" ] && [[ "$API_KEY" == *":"* ]]; then
    echo "✅ $API_KEY"
else
    echo "❌ 读取失败"
fi

echo -n "  Base URL: "
BASE_URL="$(./read_config.sh base_url)"
if [ -n "$BASE_URL" ]; then
    echo "✅ $BASE_URL"
else
    echo "❌ 读取失败"
fi

echo -n "  Model: "
MODEL="$(./read_config.sh model)"
if [ -n "$MODEL" ]; then
    echo "✅ $MODEL"
else
    echo "❌ 读取失败"
fi

echo -n "  Duration: "
DURATION="$(./read_config.sh duration)"
if [ -n "$DURATION" ]; then
    echo "✅ ${DURATION}秒"
else
    echo "❌ 读取失败"
fi

# 检查API连通性
echo ""
echo "🌐 测试API连通性:"
echo -n "  连接测试: "
if curl -s --connect-timeout 5 "$BASE_URL/v1/async/results" > /dev/null 2>&1; then
    echo "✅ API端点可达"
else
    echo "⚠️ API端点不可达 (可能是网络或配置问题)"
fi

echo ""
echo "📋 配置测试完成"
echo ""
echo "💡 使用说明:"
echo "  1. 确保 $CONFIG_FILE 中的API密钥正确"
echo "  2. 运行视频生成: ./生成演示视频.sh <图片> <描述>"
echo "  3. 查询结果: ./查询视频生成结果.sh <任务ID>"
