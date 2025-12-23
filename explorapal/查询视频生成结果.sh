#!/bin/bash

# ExploraPal 异步视频生成结果查询脚本
# 用于查询豆包视频生成任务的状态和结果

echo "🎬 ExploraPal 异步视频生成结果查询"
echo "======================================"

# 检查参数
if [ $# -eq 0 ]; then
    echo "使用方法: $0 <任务ID>"
    echo "  任务ID: 视频生成任务的ID"
    echo ""
    echo "示例:"
    echo "  $0 ea430074-4dfb-4726-99f5-056d047ba610"
    echo ""
    echo "也可以从任务文件中读取:"
    echo "  $0 \$(cat video_generation_task_*.txt | tail -1)"
    exit 1
fi

TASK_ID="$1"
ASYNC_API_URL="http://apx-api.tal.com/v1/async/results/$TASK_ID"
API_KEY="300000712:9ffb0776d5409f4131f0a314fd5cb80e"  # 请替换为真实的API密钥 (格式: appId:appKey)

echo "📋 任务ID: $TASK_ID"
echo "🔗 查询URL: $ASYNC_API_URL"
echo ""

# 查询任务状态
echo "🔍 查询任务状态..."

# 检查是否为模拟任务ID
if [[ "$TASK_ID" == mock-* ]]; then
    echo "🎭 检测到模拟任务ID，使用模拟响应"

    # 模拟任务状态变化
    MOCK_TIMESTAMP=$(echo "$TASK_ID" | cut -d- -f2)
    CURRENT_TIME=$(date +%s)
    ELAPSED_TIME=$((CURRENT_TIME - MOCK_TIMESTAMP))

    if [ $ELAPSED_TIME -lt 30 ]; then
        # 30秒内：等待中
        TASK_STATUS=1
        STATUS_TEXT="等待中"
    elif [ $ELAPSED_TIME -lt 60 ]; then
        # 30-60秒：处理中
        TASK_STATUS=2
        STATUS_TEXT="处理中"
    else
        # 60秒后：完成
        TASK_STATUS=3
        STATUS_TEXT="已完成"
    fi

    # 计算创建时间 (macOS兼容)
    CREATED_TIME=$(date -r $MOCK_TIMESTAMP '+%Y-%m-%d %H:%M:%S' 2>/dev/null || date '+%Y-%m-%d %H:%M:%S')

    HTTP_STATUS="200"
    ACTUAL_RESPONSE="{
        \"id\":\"$TASK_ID\",
        \"_id\":12345,
        \"created_at\":\"$CREATED_TIME\",
        \"updated_at\":\"$(date '+%Y-%m-%d %H:%M:%S')\",
        \"status\":$TASK_STATUS,
        \"response\":{
            \"video_url\":\"https://example.com/mock_video_$TASK_ID.mp4\",
            \"video_cover_url\":\"https://example.com/mock_cover_$TASK_ID.jpg\"
        },
        \"request\":{}
    }"

    echo "🎭 模拟任务状态: $STATUS_TEXT ($TASK_STATUS)"
else
    # 真实的API调用
    RESPONSE=$(curl -s -w "\nHTTP_STATUS:%{http_code}\n" -X GET "$ASYNC_API_URL" \
      -H "api-key: $API_KEY" \
      -H "X-APX-Model: doubao-seedance-1.0-lite-t2v")

    # 提取HTTP状态码和响应体
    HTTP_STATUS=$(echo "$RESPONSE" | grep "HTTP_STATUS:" | cut -d: -f2)
    ACTUAL_RESPONSE=$(echo "$RESPONSE" | grep -v "HTTP_STATUS:")
fi

echo "📊 HTTP状态码: $HTTP_STATUS"
echo ""

if [ "$HTTP_STATUS" != "200" ]; then
    echo "❌ 查询失败，HTTP状态码: $HTTP_STATUS"
    echo "响应内容: $ACTUAL_RESPONSE"
    exit 1
fi

# 解析JSON响应
if ! echo "$ACTUAL_RESPONSE" | jq . 2>/dev/null; then
    echo "❌ 响应不是有效的JSON格式"
    echo "原始响应: $ACTUAL_RESPONSE"
    exit 1
fi

# 提取任务信息
TASK_STATUS=$(echo "$ACTUAL_RESPONSE" | jq -r '.status // 0')
CREATED_AT=$(echo "$ACTUAL_RESPONSE" | jq -r '.created_at // "未知"')
UPDATED_AT=$(echo "$ACTUAL_RESPONSE" | jq -r '.updated_at // "未知"')

echo "📊 任务状态信息："
echo "  • 任务ID: $TASK_ID"
echo "  • 创建时间: $CREATED_AT"
echo "  • 最后更新: $UPDATED_AT"
echo "  • 当前状态: $TASK_STATUS"

# 根据状态显示不同信息
case $TASK_STATUS in
    1)
        echo "  • 状态说明: ⏳ 等待中"
        echo ""
        echo "💡 任务正在队列中等待处理，请稍后再次查询"
        ;;
    2)
        echo "  • 状态说明: 🔄 处理中"
        echo ""
        echo "💡 视频正在生成中，请稍后再次查询"
        ;;
    3)
        echo "  • 状态说明: ✅ 已完成"
        echo ""

        # 提取视频URL
        VIDEO_URL=$(echo "$ACTUAL_RESPONSE" | jq -r '.response.video_url // empty')

        if [ -n "$VIDEO_URL" ] && [ "$VIDEO_URL" != "null" ] && [ "$VIDEO_URL" != "empty" ]; then
            echo "🎬 视频生成成功！"
            echo "📹 视频URL: $VIDEO_URL"
            echo ""

            # 下载视频文件
            echo "💾 开始下载视频文件..."
            FILENAME="ExploraPal_演示视频_$(date +%Y%m%d_%H%M%S).mp4"

            if curl -s -o "$FILENAME" "$VIDEO_URL"; then
                if [ -f "$FILENAME" ] && [ -s "$FILENAME" ]; then
                    FILE_SIZE=$(stat -f%z "$FILENAME" 2>/dev/null || stat -c%s "$FILENAME" 2>/dev/null || echo "0")
                    echo "✅ 视频文件下载成功！"
                    echo "📁 文件名: $FILENAME"
                    echo "📊 文件大小: $FILE_SIZE bytes"
                    echo "📍 保存位置: $(pwd)/$FILENAME"
                    echo ""
                    echo "🎥 视频规格："
                    echo "  • 时长：约60秒"
                    echo "  • 风格：教育风格"
                    echo "  • 格式：MP4"
                    echo "  • 模型：豆包Doubao-Seedance-1.0-lite-t2v"
                    echo ""
                    echo "💡 播放提示："
                    echo "  • 使用 ffplay $FILENAME 或 vlc $FILENAME 播放"
                    echo "  • 或直接在浏览器中打开此文件"
                else
                    echo "❌ 文件下载失败或文件为空"
                    echo "请检查网络连接或手动下载："
                    echo "curl -o video.mp4 \"$VIDEO_URL\""
                fi
            else
                echo "❌ 视频下载失败"
                echo "请手动下载："
                echo "curl -o video.mp4 \"$VIDEO_URL\""
            fi
        else
            echo "❌ 未找到视频URL"
            echo "完整的响应内容："
            echo "$ACTUAL_RESPONSE" | jq .
        fi
        ;;
    4)
        echo "  • 状态说明: ❌ 处理失败"
        echo ""

        # 提取错误信息
        ERROR_MSG=$(echo "$ACTUAL_RESPONSE" | jq -r '.response.message // "未知错误"')
        echo "❌ 视频生成失败"
        echo "📝 错误信息: $ERROR_MSG"
        echo ""
        echo "🔍 请检查："
        echo "  • 输入参数是否正确"
        echo "  • API服务状态"
        echo "  • 网络连接"
        ;;
    *)
        echo "  • 状态说明: ❓ 未知状态"
        echo ""
        echo "⚠️ 收到未知的任务状态，请联系技术支持"
        echo "完整响应: $ACTUAL_RESPONSE"
        ;;
esac

echo ""
echo "======================================"
echo "查询完成"
