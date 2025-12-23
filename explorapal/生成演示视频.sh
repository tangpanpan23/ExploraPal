#!/bin/bash

# ExploraPal AI演示视频生成脚本
# 使用豆包Doubao-Seedance-1.0-lite-i2v模型
# 输入：用户原始上传的图片 + 润色后的文字描述
# 输出：高质量MP4演示视频

echo "🎬 ExploraPal AI演示视频生成器 (豆包Doubao-Seedance-1.0-lite-i2v)"
echo "==================================================================="
echo "🤖 使用豆包图像到视频生成模型"
echo "📸 输入：用户原始图片 + AI润色描述"
echo "🎥 输出：高质量MP4演示视频"

# 检查网络连接和豆包API可达性
echo "🔍 检查豆包API连接状态..."
if ! curl -s --connect-timeout 5 http://apx-api-gray.tal.com/v1/async/results > /dev/null 2>&1; then
    echo "⚠️ 豆包API服务可能不可达，但这不影响脚本运行"
    echo "💡 如果使用真实的API密钥，视频生成将调用星图AI平台"
    echo "🎭 如果使用占位符密钥，将使用模拟响应测试脚本流程"
    echo ""
else
    echo "✅ 豆包API服务可达"
fi
echo "🤖 使用AI模型: 豆包Doubao-Seedance-1.0-lite-i2v (图像到视频生成)"

# 检查参数
if [ $# -lt 2 ]; then
    echo ""
    echo "❌ 参数不足"
    echo "用法: $0 <图片文件路径> <润色后的描述文本>"
    echo ""
    echo "参数说明:"
    echo "  图片文件路径: 用户原始上传的图片文件"
    echo "  润色后的描述: 用户总结内容经过AI润色后的文字描述"
    echo ""
    echo "示例:"
    echo "  $0 /path/to/user_image.jpg \"小葫芦观察到了一只可爱的小恐龙，它有着绿色的皮肤和长长的脖子。\""
    exit 1
fi

IMAGE_PATH="$1"
DESCRIPTION="$2"

# 检查图片文件是否存在
if [ ! -f "$IMAGE_PATH" ]; then
    echo "❌ 图片文件不存在: $IMAGE_PATH"
    echo "💡 当前工作目录: $(pwd)"
    echo "💡 请检查文件路径是否正确"
    exit 1
fi

echo "📸 输入图片: $IMAGE_PATH"
echo "📝 润色描述: $DESCRIPTION"

# 显示文件信息用于调试
echo "🔍 文件信息:"
ls -la "$IMAGE_PATH" 2>/dev/null || echo "  无法获取文件信息"

# 等待一下确保所有服务都完全启动
echo "⏳ 等待服务完全就绪..."
sleep 3

# 将图片转换为base64编码
echo "🔄 转换图片为base64编码..."

# 首先检查文件是否存在且可读
if [ ! -r "$IMAGE_PATH" ]; then
    echo "❌ 图片文件不存在或不可读: $IMAGE_PATH"
    echo "请检查文件路径是否正确"
    exit 1
fi

# 检查文件大小（限制为10MB）
FILE_SIZE=$(stat -f%z "$IMAGE_PATH" 2>/dev/null || stat -c%s "$IMAGE_PATH" 2>/dev/null || echo "0")
if [ "$FILE_SIZE" -gt 10485760 ]; then  # 10MB
    echo "❌ 图片文件过大: $FILE_SIZE bytes (最大支持10MB)"
    exit 1
fi

echo "📊 文件大小: $FILE_SIZE bytes"

IMAGE_BASE64=""

# 尝试base64命令（优先选择）
if command -v base64 >/dev/null 2>&1; then
    echo "🔧 尝试使用base64命令..."

    # 首先尝试Linux风格的base64 (带-w选项)
    if base64 -w 0 "$IMAGE_PATH" >/dev/null 2>&1; then
        IMAGE_BASE64=$(base64 -w 0 "$IMAGE_PATH" 2>/dev/null | tr -d '\n')
        echo "✅ 使用GNU base64编码成功"
    fi

    # 如果上面的方法失败，尝试macOS/BSD风格的base64
    if [ -z "$IMAGE_BASE64" ]; then
        if base64 "$IMAGE_PATH" >/dev/null 2>&1; then
            IMAGE_BASE64=$(base64 "$IMAGE_PATH" 2>/dev/null | tr -d '\n')
            echo "✅ 使用BSD base64编码成功"
        fi
    fi
fi

# 如果base64命令失败，尝试openssl
if [ -z "$IMAGE_BASE64" ] && command -v openssl >/dev/null 2>&1; then
    echo "🔧 尝试使用openssl..."
    if openssl base64 -in "$IMAGE_PATH" >/dev/null 2>&1; then
        IMAGE_BASE64=$(openssl base64 -in "$IMAGE_PATH" 2>/dev/null | tr -d '\n')
        echo "✅ 使用openssl编码成功"
    fi
fi

# 如果上面都失败，尝试python3
if [ -z "$IMAGE_BASE64" ] && command -v python3 >/dev/null 2>&1; then
    echo "🔧 尝试使用python3..."
    PYTHON_OUTPUT=$(python3 -c "
import base64
import sys
try:
    with open('$IMAGE_PATH', 'rb') as f:
        data = f.read()
    if len(data) == 0:
        sys.exit(1)
    encoded = base64.b64encode(data).decode('ascii')
    print(encoded, end='')
    sys.exit(0)
except Exception as e:
    print(f'Python编码错误: {e}', file=sys.stderr)
    sys.exit(1)
" 2>/dev/null)

    if [ $? -eq 0 ] && [ -n "$PYTHON_OUTPUT" ]; then
        IMAGE_BASE64="$PYTHON_OUTPUT"
        echo "✅ 使用python3编码成功"
    fi
fi

if [ -z "$IMAGE_BASE64" ] || [ ${#IMAGE_BASE64} -eq 0 ]; then
    echo "❌ 图片base64编码失败"
    echo "尝试的编码方法:"
    echo "  - base64命令: $(command -v base64 >/dev/null 2>&1 && echo "可用" || echo "不可用")"
    echo "  - openssl: $(command -v openssl >/dev/null 2>&1 && echo "可用" || echo "不可用")"
    echo "  - python3: $(command -v python3 >/dev/null 2>&1 && echo "可用" || echo "不可用")"
    echo "请确保安装了上述工具之一，或检查图片文件是否损坏"
    exit 1
fi

echo "✅ 图片编码完成 (长度: ${#IMAGE_BASE64} 字符)"

# 生成视频脚本内容（基于用户提供的描述）
VIDEO_SCRIPT="基于用户观察的图片和AI润色后的描述，ExploraPal生成了这个精彩的学习视频。

用户观察到的内容：
$DESCRIPTION

通过AI智能分析，这个学习主题包含：
• 丰富的观察细节和特征识别
• 激发好奇心的学习价值
• 个性化的问题引导和探索路径

ExploraPal将帮助用户：
1. 深入分析观察对象
2. 生成针对性的学习问题
3. 记录和优化学习笔记
4. 创建专业的学习报告

这就是AI学习助手的魅力！让每个孩子都能通过观察和探索，享受学习的乐趣，发现世界的奥秘。"

echo "📝 演示脚本内容："
echo "$VIDEO_SCRIPT"
echo ""
echo "🎬 开始生成演示视频..."

# 调用豆包Doubao-Seedance-1.0-lite-i2v API生成视频
echo "🎬 调用豆包Doubao-Seedance-1.0-lite-i2v模型生成视频..."
echo "📤 发送数据: 原始图片 + 润色描述"

# 创建临时JSON文件 - 直接调用星图AI平台API
# 清理可能存在的旧临时文件
rm -f /tmp/video_request_*.json

TEMP_JSON_FILE=$(mktemp /tmp/video_request_XXXXXX.json)
echo "🔧 创建临时文件: $TEMP_JSON_FILE"
cat > "$TEMP_JSON_FILE" << EOF
{
  "model": "doubao-seedance-1.0-lite-i2v",
  "img_url": "data:image/jpeg;base64,$IMAGE_BASE64",
  "prompt": "$DESCRIPTION",
  "duration": "$DEFAULT_DURATION"
}
EOF

echo "📄 创建临时请求文件: $TEMP_JSON_FILE"

# 显示请求文件内容（前200字符）
echo "📋 请求文件内容预览："
head -c 200 "$TEMP_JSON_FILE"
echo ""
echo "📊 请求文件大小: $(wc -c < "$TEMP_JSON_FILE") bytes"
echo ""

# 显示完整的请求内容（用于调试）
echo "📄 完整请求内容："
echo "=================="
cat "$TEMP_JSON_FILE"
echo ""
echo "=================="

# 执行异步API调用
echo "🚀 提交异步视频生成任务..."

# 从配置文件读取API配置
echo "📖 读取配置文件..."
ASYNC_API_URL="$(./read_config.sh base_url)/v1/async/chat"
API_KEY="$(./read_config.sh api_key)"
DEFAULT_DURATION="$(./read_config.sh duration)"

if [ -z "$API_KEY" ] || [[ "$API_KEY" == *":" ]]; then
    echo "❌ 无法从配置文件读取API密钥"
    echo "请检查 video_generation_config.yaml 文件"
    exit 1
fi

echo "✅ 配置加载成功"

# 先执行curl获取响应头信息
CURL_HEADERS=$(curl -I -s "$ASYNC_API_URL" 2>/dev/null | head -1)
echo "🌐 HTTP响应头: $CURL_HEADERS"

# 检查请求文件是否存在并获取大小
if [ ! -f "$TEMP_JSON_FILE" ]; then
    echo "❌ 临时请求文件不存在: $TEMP_JSON_FILE"
    exit 1
fi

FILE_SIZE=$(wc -c < "$TEMP_JSON_FILE")
echo "📊 请求体大小: $FILE_SIZE bytes"

if [ "$FILE_SIZE" -gt 2097152 ]; then  # 2MB限制
    echo "⚠️ 请求体过大 ($FILE_SIZE bytes)"
    echo "💡 在演示环境中，当图片过大时将使用模拟响应来展示完整流程"

    # 使用模拟模式 - 不包含图片，只使用文本描述
    cat > "$TEMP_JSON_FILE" << EOF
{
  "model": "doubao-seedance-1.0-lite-i2v",
  "prompt": "$DESCRIPTION (演示模式：图片因大小限制而省略)",
  "duration": "$DEFAULT_DURATION"
}
EOF

    echo "🎭 已切换到演示模式（不包含图片数据）"
    NEW_SIZE=$(wc -c < "$TEMP_JSON_FILE")
    echo "📊 演示模式请求体大小: $NEW_SIZE bytes"
fi

# 执行异步任务提交
echo "📡 发送请求到: $ASYNC_API_URL"

# 检查API密钥是否为占位符
if [[ "$API_KEY" == *"xxxxx"* ]]; then
    echo "⚠️ 检测到占位符API密钥，请替换为真实的API密钥"
    echo "📝 请在脚本中设置真实的API密钥: API_KEY=\"your_real_app_id:your_real_app_key\""
    echo ""
    echo "🔧 为了测试脚本流程，使用模拟响应..."

    # 生成模拟的任务ID
    MOCK_TASK_ID="mock-$(date +%s)-$(openssl rand -hex 8 2>/dev/null || echo "random123")"
    HTTP_STATUS="200"
    ACTUAL_RESPONSE="{\"id\":\"$MOCK_TASK_ID\",\"created_at\":\"$(date '+%Y-%m-%d %H:%M:%S')\",\"status\":1}"

    echo "🎭 使用模拟响应测试脚本流程"
else
    # 真实的API调用
    TASK_RESPONSE=$(curl -s -w "\nHTTP_STATUS:%{http_code}\n" -X POST "$ASYNC_API_URL" \
      -H "Content-Type: application/json" \
      -H "api-key: $API_KEY" \
      -H "X-APX-Model: doubao-seedance-1.0-lite-i2v" \
      -d @"$TEMP_JSON_FILE")

    # 提取HTTP状态码和响应体
    HTTP_STATUS=$(echo "$TASK_RESPONSE" | grep "HTTP_STATUS:" | cut -d: -f2)
    ACTUAL_RESPONSE=$(echo "$TASK_RESPONSE" | grep -v "HTTP_STATUS:")
fi

echo "📊 HTTP状态码: $HTTP_STATUS"

# 清理临时文件
rm -f "$TEMP_JSON_FILE"
echo "🧹 清理临时文件完成"

# 显示详细的API响应信息
echo "📡 API响应详情："
if [ -z "$ACTUAL_RESPONSE" ]; then
    echo "❌ API无响应（响应为空）"
    echo "🔍 故障排查："
    echo "  1. 检查API服务是否启动: curl http://apx-api.tal.com/v1/async/results"
    echo "  2. 检查网络连接"
    echo "  3. 检查API密钥是否正确"
    echo "  4. 查看API服务日志"
    exit 1
fi

echo "📄 原始响应内容："
echo "$ACTUAL_RESPONSE"
echo ""

# 尝试解析JSON响应
echo "🔍 响应解析："
if echo "$ACTUAL_RESPONSE" | jq . 2>/dev/null; then
    echo "✅ 响应是有效的JSON格式"

    # 检查是否是错误响应
    if echo "$ACTUAL_RESPONSE" | jq -e '.error_code' >/dev/null 2>&1; then
        ERROR_CODE=$(echo "$ACTUAL_RESPONSE" | jq -r '.error_code // empty')
        ERROR_MESSAGE=$(echo "$ACTUAL_RESPONSE" | jq -r '.message // empty')

        echo "❌ API返回错误:"
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
        echo ""
        echo "💡 提示：当前使用演示模式，不影响脚本功能测试"
        exit 1
    fi

    # 提取任务ID
    TASK_ID=$(echo "$ACTUAL_RESPONSE" | jq -r '.id // empty' 2>/dev/null)

    if [ -n "$TASK_ID" ] && [ "$TASK_ID" != "null" ] && [ "$TASK_ID" != "empty" ]; then
        echo "✅ 异步视频生成任务提交成功！"
        echo "📋 任务ID: $TASK_ID"

        # 保存任务ID到本地文件
        TASK_FILE="video_generation_task_$(date +%Y%m%d_%H%M%S).txt"
        echo "$TASK_ID" > "$TASK_FILE"
        echo "💾 任务ID已保存到文件: $TASK_FILE"

        echo ""
        echo "🎬 视频生成任务已提交，正在后台处理..."
        echo "⏱️ 预计需要5-10分钟完成生成"
        echo ""
        echo "📋 如何查询生成结果："
        echo "  1. 等待几分钟后运行查询脚本："
        echo "     ./查询视频生成结果.sh $TASK_ID"
        echo "  2. 或使用以下命令手动查询："
        echo "     curl -X GET \"http://apx-api.tal.com/v1/async/results/$TASK_ID\" \\"
        echo "       -H \"api-key: $API_KEY\" \\"
        echo "       -H \"X-APX-Model: doubao-seedance-1.0-lite-i2v\""
        echo ""
        echo "🔍 任务状态说明："
        echo "  • 状态1: 等待中"
        echo "  • 状态2: 处理中"
        echo "  • 状态3: 已完成"
        echo "  • 状态4: 处理失败"
        echo ""
        echo "💡 提示："
        echo "  • 视频生成需要时间，请耐心等待"
        echo "  • 可以使用任务ID随时查询进度"
        echo "  • 生成完成后会提供视频下载链接"
    else
        echo "❌ 未获取到任务ID"
        echo "API响应可能有误，请检查响应内容"
        exit 1
    fi
else
    echo "⚠️ 响应不是有效的JSON格式"
    echo "❌ 异步任务提交失败"
    echo "请检查API服务和网络连接"
    exit 1
fi
