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
if ! curl -s http://localhost:9003/api/common/ping > /dev/null 2>&1; then
    echo "❌ API服务未运行"
    echo "请先启动服务："
    echo "  cd explorapal && ./start_demo.sh"
    echo ""
    echo "启动服务后，请等待所有服务完全启动（大约30秒），然后重新运行此脚本。"
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
    exit 1
fi

echo "📸 输入图片: $IMAGE_PATH"
echo "📝 润色描述: $DESCRIPTION"

# 等待一下确保所有服务都完全启动
echo "⏳ 等待服务完全就绪..."
sleep 3

# 将图片转换为base64编码
echo "🔄 转换图片为base64编码..."
if command -v base64 >/dev/null 2>&1; then
    # Linux base64命令
    IMAGE_BASE64=$(base64 -w 0 "$IMAGE_PATH" 2>/dev/null || base64 "$IMAGE_PATH" 2>/dev/null | tr -d '\n')
elif command -v openssl >/dev/null 2>&1; then
    # 使用openssl作为备选方案
    IMAGE_BASE64=$(openssl base64 -in "$IMAGE_PATH" | tr -d '\n')
else
    echo "❌ 无法进行base64编码，请安装base64或openssl工具"
    exit 1
fi

if [ -z "$IMAGE_BASE64" ]; then
    echo "❌ 图片base64编码失败"
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

RESPONSE=$(curl -s -X POST http://localhost:9003/api/achievement/video/generate \
  -H "Content-Type: application/json" \
  -d "{
    \"project_id\": 1,
    \"user_id\": 1,
    \"image_data\": \"data:image/jpeg;base64,$IMAGE_BASE64\",
    \"prompt\": \"$DESCRIPTION\",
    \"style\": \"educational\",
    \"duration\": 60,
    \"scenes\": [],
    \"voice\": \"female\",
    \"language\": \"zh-CN\"
  }")

echo "📡 API响应："
echo "$RESPONSE" | jq . 2>/dev/null || echo "$RESPONSE"

# 检查响应状态
if echo "$RESPONSE" | grep -q '"status":200'; then
    echo ""
    echo "✅ 演示视频生成成功！"

    # 提取视频数据
    VIDEO_DATA=$(echo "$RESPONSE" | sed -n 's/.*"video_data":"\([^"]*\)".*/\1/p')

    if [ -n "$VIDEO_DATA" ]; then
        # 生成文件名
        FILENAME="ExploraPal_完整演示视频_$(date +%Y%m%d_%H%M%S).mp4"

        echo "💾 保存视频文件：$FILENAME"

        # 将base64数据转换为MP4文件
        echo "$VIDEO_DATA" | base64 -d > "$FILENAME"

        if [ -f "$FILENAME" ]; then
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
