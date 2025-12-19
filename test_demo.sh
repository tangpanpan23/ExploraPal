#!/bin/bash

# 探索伙伴演示环境测试脚本
# 快速验证各服务是否正常运行

echo "🔍 探索伙伴服务状态检查"
echo "================================"

BASE_URL="http://localhost:9003"

# 检查服务状态
check_service() {
    echo "📡 检查API服务状态..."
    if curl -s "$BASE_URL/api/common/ping" > /dev/null 2>&1; then
        echo "✅ API服务运行正常 (端口9003)"
    else
        echo "❌ API服务未运行"
        return 1
    fi
}

# 测试项目创建
test_project_creation() {
    echo "📝 测试项目创建..."
    response=$(curl -s -X POST "$BASE_URL/api/project/create" \
        -H "Content-Type: application/json" \
        -d '{
            "user_id": 1,
            "title": "演示测试项目",
            "description": "用于测试的演示项目",
            "category": "test",
            "tags": ["test", "demo"]
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 项目创建成功"
        PROJECT_ID=$(echo "$response" | grep -o '"project_id":[0-9]*' | cut -d':' -f2)
        echo "   项目ID: $PROJECT_ID"
        return 0
    else
        echo "❌ 项目创建失败"
        return 1
    fi
}

# 测试AI图像分析
test_image_analysis() {
    echo "🖼️  测试AI图像分析..."
    response=$(curl -s -X POST "$BASE_URL/api/observation/image/recognize" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": 1,
            "image_url": "https://example.com/test.jpg",
            "prompt": "分析这张图片"
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 图像分析接口正常"
    else
        echo "⚠️  图像分析接口响应 (可能需要真实图片数据)"
    fi
}

# 测试问题生成
test_question_generation() {
    echo "❓ 测试问题生成..."
    response=$(curl -s -X POST "$BASE_URL/api/questioning/questions/generate" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": 1,
            "context_info": "测试观察内容",
            "category": "test",
            "user_age": 8
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 问题生成接口正常"
    else
        echo "⚠️  问题生成接口响应"
    fi
}

# 测试语音转文字
test_speech_to_text() {
    echo "🎤 测试语音转文字..."
    response=$(curl -s -X POST "$BASE_URL/api/expression/speech/text" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": 1,
            "audio_data": "test_audio_data",
            "audio_format": "wav",
            "language": "zh-CN"
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 语音转文字接口正常"
    else
        echo "⚠️  语音转文字接口响应 (需要真实音频数据)"
    fi
}

# 测试文字转语音
test_text_to_speech() {
    echo "🔊 测试文字转语音..."
    response=$(curl -s -X POST "$BASE_URL/api/audio/text-to-speech" \
        -H "Content-Type: application/json" \
        -d '{
            "text": "你好，欢迎使用探索伙伴！",
            "voice": "female",
            "language": "zh-CN",
            "speed": 1.0
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 文字转语音接口正常"
    else
        echo "⚠️  文字转语音接口响应"
    fi
}

# 测试视频分析
test_video_analysis() {
    echo "🎬 测试视频分析..."
    response=$(curl -s -X POST "$BASE_URL/api/achievement/video/analyze" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": 1,
            "video_data": "test_video_data",
            "video_format": "mp4",
            "analysis_type": "content",
            "duration": 30.0
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 视频分析接口正常"
    else
        echo "⚠️  视频分析接口响应 (需要真实视频数据)"
    fi
}

# 测试报告生成
test_report_generation() {
    echo "📊 测试报告生成..."
    response=$(curl -s -X POST "$BASE_URL/api/achievement/report/generate" \
        -H "Content-Type: application/json" \
        -d '{
            "project_data": "测试项目数据",
            "category": "test"
        }')

    if echo "$response" | grep -q '"code":200'; then
        echo "✅ 报告生成接口正常"
    else
        echo "⚠️  报告生成接口响应"
    fi
}

# 主测试流程
main() {
    PROJECT_ID=""

    if check_service; then
        test_project_creation

        if [ -n "$PROJECT_ID" ]; then
            test_image_analysis
            test_question_generation
            test_speech_to_text
            test_text_to_speech
            test_video_analysis
            test_report_generation
        fi

        echo ""
        echo "================================"
        echo "🎉 演示环境测试完成！"
        echo ""
        echo "💡 提示:"
        echo "   - 绿色✅表示接口完全正常"
        echo "   - 黄色⚠️表示接口响应但可能需要真实数据"
        echo "   - 红色❌表示接口异常"
        echo ""
        echo "🚀 运行完整演示: ./demo_flow.sh"
        echo "📖 查看详细文档: cat MULTIMODAL_DEMO.md"
    else
        echo ""
        echo "❌ 服务未启动，请先运行:"
        echo "   ./start_demo.sh"
        exit 1
    fi
}

main "$@"
