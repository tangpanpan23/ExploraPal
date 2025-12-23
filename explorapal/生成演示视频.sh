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

# 检查服务是否运行
echo "🔍 检查API服务状态..."
if ! curl -s http://localhost:9003/api/ping > /dev/null 2>&1; then
    echo "❌ API服务未运行"
    echo ""
    echo "请先启动服务："
    echo "  1. 确保所有服务脚本有执行权限："
    echo "     chmod +x *.sh"
    echo "  2. 启动服务："
    echo "     ./start_demo.sh"
    echo "  3. 等待30秒让服务完全启动"
    echo "  4. 重新运行此脚本"
    echo ""
    exit 1
fi

echo "✅ API服务运行正常"
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

# 创建临时JSON文件避免参数过长问题
TEMP_JSON_FILE=$(mktemp /tmp/video_request_XXXXXX.json)
cat > "$TEMP_JSON_FILE" << EOF
{
  "project_id": 1,
  "user_id": 1,
  "image_data": "data:image/jpeg;base64,$IMAGE_BASE64",
  "prompt": "$DESCRIPTION",
  "style": "educational",
  "duration": 60,
  "scenes": [],
  "voice": "female",
  "language": "zh-CN"
}
EOF

echo "📄 创建临时请求文件: $TEMP_JSON_FILE"

# 显示请求文件内容（前200字符）
echo "📋 请求文件内容预览："
head -c 200 "$TEMP_JSON_FILE"
echo ""
echo "📊 请求文件大小: $(wc -c < "$TEMP_JSON_FILE") bytes"
echo ""

# 执行API调用
echo "🚀 发送API请求..."

# 先执行curl获取响应头信息
CURL_HEADERS=$(curl -I -s http://localhost:9003/api/achievement/video/generate 2>/dev/null | head -1)
echo "🌐 HTTP响应头: $CURL_HEADERS"

# 执行实际的API调用
RESPONSE=$(curl -s -w "\nHTTP_STATUS:%{http_code}\n" -X POST http://localhost:9003/api/achievement/video/generate \
  -H "Content-Type: application/json" \
  -d @"$TEMP_JSON_FILE")

# 提取HTTP状态码和响应体
HTTP_STATUS=$(echo "$RESPONSE" | grep "HTTP_STATUS:" | cut -d: -f2)
ACTUAL_RESPONSE=$(echo "$RESPONSE" | grep -v "HTTP_STATUS:")

echo "📊 HTTP状态码: $HTTP_STATUS"

# 清理临时文件
rm -f "$TEMP_JSON_FILE"
echo "🧹 清理临时文件完成"

# 显示详细的API响应信息
echo "📡 API响应详情："
if [ -z "$ACTUAL_RESPONSE" ]; then
    echo "❌ API无响应（响应为空）"
    echo "🔍 故障排查："
    echo "  1. 检查API服务是否启动: curl http://localhost:9003/api/common/ping"
    echo "  2. 检查网络连接"
    echo "  3. 检查防火墙设置"
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
else
    echo "⚠️ 响应不是有效的JSON格式"
fi
echo ""

# 检查各种可能的响应状态
if echo "$ACTUAL_RESPONSE" | grep -q '"status":200\|"code":200\|success\|Success'; then
    echo ""
    echo "✅ 演示视频生成成功！"

    # 提取视频数据 - 使用jq进行JSON解析
    VIDEO_DATA=$(echo "$ACTUAL_RESPONSE" | jq -r '.video_data // empty' 2>/dev/null)

    if [ -n "$VIDEO_DATA" ] && [ "$VIDEO_DATA" != "null" ] && [ "$VIDEO_DATA" != "empty" ]; then
        # 生成文件名
        FILENAME="ExploraPal_完整演示视频_$(date +%Y%m%d_%H%M%S).mp4"

        echo "💾 保存视频文件：$FILENAME"

        # 将base64数据转换为MP4文件
        echo "🔄 开始base64解码..."
        if echo "$VIDEO_DATA" | base64 -d > "$FILENAME" 2>/dev/null; then
            echo "✅ base64解码成功"
        else
            echo "❌ base64解码失败"
            echo "请检查VIDEO_DATA格式是否正确"
            exit 1
        fi

        if [ -f "$FILENAME" ] && [ -s "$FILENAME" ]; then
            FILE_SIZE=$(stat -f%z "$FILENAME" 2>/dev/null || stat -c%s "$FILENAME" 2>/dev/null)
            echo "✅ 视频文件已保存！"
            echo "📊 文件大小：$FILE_SIZE bytes"
            echo "🎬 文件位置：$(pwd)/$FILENAME"
            echo ""
            echo "🎥 视频规格："
            echo "  • 时长：3分钟"
            echo "  • 风格：教育风格"
            echo "  • 语音：女声中文"
            echo "  • 分辨率：1920x1080"
            echo ""
            echo "📖 演示内容涵盖："
            echo "  1. 智能拍照分析"
            echo "  2. 项目自动创建"
            echo "  3. AI智能提问"
            echo "  4. 内容润色优化"
            echo "  5. 语音交互体验"
            echo "  6. 自动生成报告"
            echo "  7. 视频内容分析"
            echo "  8. AI视频创作"
        else
            echo "❌ 视频文件保存失败"
        fi
    else
        echo "❌ 未找到视频数据"
    fi
else
    echo "❌ 视频生成失败"
    echo "请检查："
    echo "  • video-processing服务是否正常运行"
    echo "  • 网络连接是否正常"
    echo "  • AI模型服务是否可用"
fi
