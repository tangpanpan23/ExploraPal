#!/bin/bash

# 探索伙伴多模态AI学习平台 - 完整演示流程脚本
# 使用方法: ./demo_flow.sh

set -e

echo "🎨 探索伙伴多模态AI学习平台 - 完整演示流程"
echo "========================================================"

BASE_URL="http://localhost:9003"
PROJECT_ID=""
USER_ID="1"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_step() {
    echo -e "\n${BLUE}[步骤 $1]${NC} $2"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_info() {
    echo -e "${YELLOW}ℹ $1${NC}"
}

# 检查服务是否运行
check_service() {
    if curl -s "$BASE_URL/api/common/ping" > /dev/null 2>&1; then
        print_success "服务运行正常"
    else
        echo -e "${RED}✗ 服务未运行，请先启动服务${NC}"
        echo "运行命令: ./start_demo.sh"
        exit 1
    fi
}

# 步骤1: 创建探索项目
create_project() {
    print_step "1" "创建恐龙探索项目"

    response=$(curl -s -X POST "$BASE_URL/api/project/create" \
        -H "Content-Type: application/json" \
        -d '{
            "user_id": '$USER_ID',
            "title": "小明的恐龙探索之旅",
            "description": "跟随小明一起探索古老的恐龙世界，了解恐龙的特征、生活习性和进化历程",
            "category": "dinosaur",
            "tags": ["恐龙", "古生物", "进化", "探索"]
        }')

    if echo "$response" | grep -q '"code":200'; then
        PROJECT_ID=$(echo "$response" | grep -o '"project_id":[0-9]*' | cut -d':' -f2)
        print_success "项目创建成功 (ID: $PROJECT_ID)"
        echo "响应: $response"
    else
        echo -e "${RED}项目创建失败${NC}"
        echo "响应: $response"
        exit 1
    fi
}

# 步骤2: AI图像分析
analyze_image() {
    print_step "2" "AI图像分析 - 识别恐龙化石"

    response=$(curl -s -X POST "$BASE_URL/api/observation/image/recognize" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID',
            "image_url": "https://example.com/dinosaur-fossil.jpg",
            "prompt": "分析这块恐龙化石，识别种类并描述特征"
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "图像分析完成"
        echo "识别结果: $(echo "$response" | grep -o '"object_name":"[^"]*"' | cut -d'"' -f4)"
    else
        echo -e "${YELLOW}图像分析模拟 (使用示例数据)${NC}"
        echo "实际环境中需要有效的图片URL"
    fi
}

# 步骤3: AI生成问题
generate_questions() {
    print_step "3" "AI生成个性化问题"

    response=$(curl -s -X POST "$BASE_URL/api/questioning/questions/generate" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID',
            "context_info": "小明观察到了一块三角龙化石，上面有三只角和骨板",
            "category": "dinosaur",
            "user_age": 8
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "问题生成完成"
        echo "生成的问题数量: $(echo "$response" | grep -o '"questions"' | wc -l)"
    else
        echo -e "${YELLOW}问题生成模拟${NC}"
    fi
}

# 步骤4: 语音转文字演示
speech_to_text() {
    print_step "4" "语音转文字演示"

    print_info "模拟语音数据处理..."
    print_info "实际使用时需要上传真实的音频文件"

    # 这里模拟语音转文字的请求
    response=$(curl -s -X POST "$BASE_URL/api/expression/speech/text" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID',
            "audio_data": "UklGRnoGAABXQVZFZm10IAAAAAEAAQARAAAAEAAAAAEACABkYXRhAgAAAAEA",
            "audio_format": "wav",
            "language": "zh-CN"
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "语音处理完成"
    else
        print_info "语音处理需要真实的音频数据"
    fi
}

# 步骤5: AI润色笔记
polish_note() {
    print_step "5" "AI润色笔记"

    response=$(curl -s -X POST "$BASE_URL/api/expression/note/polish" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID',
            "raw_content": "三角龙好厉害啊！它有三只角，可以保护自己不被别的恐龙吃掉。骨板也很厚，看起来像铠甲一样。",
            "content_type": "speech",
            "context_info": {
                "observation_results": "三角龙化石，三只角，骨板",
                "previous_answers": "防御，草食性",
                "project_category": "dinosaur"
            }
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "笔记润色完成"
        echo "润色后的标题: $(echo "$response" | grep -o '"title":"[^"]*"' | cut -d'"' -f4)"
    else
        echo -e "${YELLOW}笔记润色功能演示${NC}"
    fi
}

# 步骤6: 文字转语音
text_to_speech() {
    print_step "6" "文字转语音演示"

    response=$(curl -s -X POST "$BASE_URL/api/audio/text-to-speech" \
        -H "Content-Type: application/json" \
        -d '{
            "text": "欢迎来到恐龙世界！三角龙是一种非常有趣的恐龙。",
            "voice": "female",
            "language": "zh-CN",
            "speed": 1.0
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "语音合成完成"
    else
        print_info "语音合成功能演示"
    fi
}

# 步骤7: AI视频生成
generate_video() {
    print_step "7" "AI视频生成演示"

    print_info "正在生成教学视频..."
    print_info "这可能需要一些时间，请耐心等待"

    response=$(curl -s -X POST "$BASE_URL/api/achievement/video/generate" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID',
            "script": "欢迎来到恐龙世界！今天我们来学习三角龙。三角龙是一种古老的爬行动物，有三只角和坚硬的骨板...",
            "style": "educational",
            "duration": 30,
            "scenes": [
                "三角龙外形介绍",
                "生活习性展示"
            ],
            "voice": "female",
            "language": "zh-CN"
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "视频生成完成"
        echo "视频时长: $(echo "$response" | grep -o '"duration":[0-9.]*' | cut -d':' -f2)秒"
    else
        print_info "视频生成功能演示（需要AI模型支持）"
    fi
}

# 步骤8: 视频分析
analyze_video() {
    print_step "8" "视频内容分析"

    print_info "正在分析视频内容..."

    response=$(curl -s -X POST "$BASE_URL/api/achievement/video/analyze" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID',
            "video_data": "AAAAHGZ0eXBtcDQyAAACAGlzb21pc28yYXZjMQAAAAhmcmVlAAAGF21kYXQ",
            "video_format": "mp4",
            "analysis_type": "content",
            "duration": 30.0
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "视频分析完成"
    else
        print_info "视频分析功能演示（需要真实的视频数据）"
    fi
}

# 步骤9: 生成研究报告
generate_report() {
    print_step "9" "生成研究报告"

    response=$(curl -s -X POST "$BASE_URL/api/achievement/report/generate" \
        -H "Content-Type: application/json" \
        -d '{
            "project_data": "小明通过观察三角龙化石，学习了恐龙的特征、生活习性和防御机制",
            "category": "dinosaur"
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "研究报告生成完成"
        echo "报告标题: $(echo "$response" | grep -o '"title":"[^"]*"' | cut -d'"' -f4)"
    else
        print_info "报告生成功能演示"
    fi
}

# 步骤10: 查看项目成果
view_project() {
    print_step "10" "查看完整项目成果"

    response=$(curl -s -X POST "$BASE_URL/api/project/detail" \
        -H "Content-Type: application/json" \
        -d '{
            "project_id": '$PROJECT_ID',
            "user_id": '$USER_ID'
        }')

    if echo "$response" | grep -q '"code":200'; then
        print_success "项目详情获取成功"
        echo "项目状态: $(echo "$response" | grep -o '"status":"[^"]*"' | cut -d'"' -f4)"
    else
        echo -e "${YELLOW}项目详情查看${NC}"
    fi
}

# 主流程
main() {
    echo "开始执行完整演示流程..."
    echo "========================================================"

    check_service

    create_project
    analyze_image
    generate_questions
    speech_to_text
    polish_note
    text_to_speech
    generate_video
    analyze_video
    generate_report
    view_project

    echo ""
    echo "========================================================"
    print_success "演示流程执行完成！"
    echo ""
    echo -e "${BLUE}🎉 恭喜！你已经体验了探索伙伴的完整AI学习流程${NC}"
    echo ""
    echo -e "${YELLOW}📖 详细文档: MULTIMODAL_DEMO.md${NC}"
    echo -e "${YELLOW}🌐 API接口: http://localhost:9003${NC}"
    echo -e "${YELLOW}📊 项目ID: $PROJECT_ID${NC}"
}

# 执行主流程
main "$@"
