#!/bin/bash

# ExploraPal AI演示视频生成脚本
# 基于demo.html的完整演示流程，自动生成MP4演示视频

echo "🎬 ExploraPal AI演示视频生成器"
echo "=================================="

# 检查服务是否运行
echo "🔍 检查video-processing服务状态..."
if ! curl -s http://localhost:8080/api/video/generate > /dev/null; then
    echo "❌ video-processing服务未运行"
    echo "请先启动服务："
    echo "  cd explorapal && ./start_demo.sh"
    exit 1
fi

echo "✅ 服务运行正常"

# 定义演示视频脚本内容
VIDEO_SCRIPT="大家好！欢迎体验ExploraPal - 多模态AI学习平台的全流程演示。

第一步：智能拍照分析
小葫芦拍下一张恐龙图片，AI立即识别出这是三角龙化石，并分析其特征：三只角、骨板、生存环境等。

第二步：项目自动创建
AI基于图片分析结果，自动创建'小葫芦的恐龙探索之旅'项目，包含合适的年龄设置和学习目标。

第三步：AI智能提问
系统生成针对性的学习问题：
- 三角龙有哪些特殊的身体特征？
- 三角龙生活在哪个地质时期？
- 三角龙的食性是什么？
- 三角龙如何保护自己？

第四步：内容润色优化
孩子口述的想法被转换为流畅的文字，并经过AI润色，形成完整的学习笔记。

第五步：语音交互体验
支持语音录入和文字转语音播报，让学习过程更加自然和生动。

第六步：自动生成报告
AI根据整个学习过程，生成专业的学习报告，包含发现总结和进一步探索建议。

第七步：视频内容分析
上传学习相关的视频，AI自动分析内容、场景和主题，为学习提供更多维度。

第八步：AI视频创作
基于学习内容和报告，AI自动生成教学视频，包含动画演示和讲解。

这就是ExploraPal的完整学习魔法！通过多模态AI技术，让每个孩子都能享受个性化的探索式学习体验。

感谢观看！让我们一起开启AI学习的精彩旅程！"

echo "📝 演示脚本内容："
echo "$VIDEO_SCRIPT"
echo ""
echo "🎬 开始生成演示视频..."

# 调用video-processing API生成视频
RESPONSE=$(curl -s -X POST http://localhost:8080/api/video/generate \
  -H "Content-Type: application/json" \
  -d "{
    \"script\": \"$VIDEO_SCRIPT\",
    \"style\": \"educational\",
    \"duration\": 180,
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
