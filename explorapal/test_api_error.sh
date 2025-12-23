#!/bin/bash

# 测试API错误响应的脚本

echo "🧪 测试API错误响应处理"
echo "========================"

# 模拟API错误响应
MOCK_ERROR_RESPONSE='{"error_code":110000,"message":"Header 模型不存在"}'

echo "📡 模拟API错误响应："
echo "$MOCK_ERROR_RESPONSE"
echo ""

# 测试错误处理逻辑
if echo "$MOCK_ERROR_RESPONSE" | jq -e '.error_code' >/dev/null 2>&1; then
    ERROR_CODE=$(echo "$MOCK_ERROR_RESPONSE" | jq -r '.error_code // empty')
    ERROR_MESSAGE=$(echo "$MOCK_ERROR_RESPONSE" | jq -r '.message // empty')

    echo "✅ 错误检测成功！"
    echo "  • 错误码: $ERROR_CODE"
    echo "  • 错误信息: $ERROR_MESSAGE"
    echo ""

    # 提供针对性的解决建议
    case "$ERROR_CODE" in
        "110000")
            echo "🔧 解决方案："
            echo "  1. 确认应用已配置为异步应用（需要平台管理员操作）"
            echo "  2. 检查模型名称是否正确: doubao-seedance-1.0-lite-i2v"
            echo "  3. 确认API密钥有效且格式正确"
            echo "  4. 联系星图AI平台技术支持"
            ;;
        *)
            echo "🔧 通用解决方案："
            echo "  1. 检查API密钥是否正确"
            echo "  2. 确认网络连接正常"
            echo "  3. 查看星图AI平台文档"
            ;;
    esac
else
    echo "❌ 错误检测失败"
fi

echo ""
echo "测试完成 ✅"
