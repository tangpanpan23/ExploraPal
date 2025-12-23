#!/bin/bash

# ExploraPal 完整学习流程演示
# 包含TALect MCP教学资源服务

set -e

echo "🎯 ExploraPal 完整学习流程演示"
echo "=============================="
echo ""

# 检查参数
if [ $# -lt 1 ]; then
    echo "用法: $0 <图片文件路径> [年级级别]"
    echo "示例: $0 dinosaur.jpg grade_3"
    exit 1
fi

IMAGE_PATH="$1"
GRADE_LEVEL="${2:-grade_3}"

echo "📸 输入图片: $IMAGE_PATH"
echo "🎓 年级级别: $GRADE_LEVEL"
echo ""

# 检查TALect MCP服务状态
if [ ! -f "mcp_config.yaml" ]; then
    echo "ℹ️  MCP服务配置未检测到"
    MCP_ENABLED=false
else
    MCP_ENABLED=true
    echo "✅ MCP服务已配置"
fi

echo ""
echo "🚀 开始完整AI学习流程..."
echo ""

# 步骤1: AI图像分析 (ExploraPal能力)
echo "🔍 步骤1: AI图像分析..."
echo "   使用 qwen3-vl-plus 进行多模态深度理解"

# 这里应该调用实际的AI服务
# 模拟分析结果
ANALYSIS_RESULT="{
    \"object_name\": \"三角龙化石\",
    \"category\": \"古生物学\",
    \"confidence\": 0.95,
    \"description\": \"这是一块保存完好的三角龙化石，具有典型的三角龙特征：头顶三个角、颈部骨板、强壮的四肢\",
    \"key_features\": [
        \"三个突出的头顶角\",
        \"颈部和背部排列的骨板\",
        \"强壮的四肢和尾巴\",
        \"保存完好的骨骼结构\"
    ],
    \"scientific_name\": \"Triceratops\"
}"

echo "   ✅ 分析完成: 三角龙化石识别成功"
echo "   📊 置信度: 95%"
echo ""

# 步骤2: 教学资源搜索 (TALect MCP服务)
if [ "$MCP_ENABLED" = true ]; then
    echo "📚 步骤2: 教学资源搜索..."
    echo "   调用TALect MCP服务搜索相关教学资源"

    # 解析分析结果，构建搜索查询
    OBJECT_NAME=$(echo $ANALYSIS_RESULT | jq -r '.object_name')
    CATEGORY=$(echo $ANALYSIS_RESULT | jq -r '.category')
    SEARCH_QUERY="$OBJECT_NAME $CATEGORY 学习资料"

    echo "   🔍 搜索关键词: $SEARCH_QUERY"
    echo "   🎓 年级筛选: $GRADE_LEVEL"
    echo "   📖 学科: science"

    # 这里应该调用MCP客户端进行搜索
    # 模拟搜索结果
    SEARCH_RESULTS="找到 5 个相关教学素材：
1. 三角龙科普视频 (ID: mat_001)
2. 恐龙时代教学课件 (ID: mat_002)
3. 古生物化石观察指南 (ID: mat_003)
4. 恐龙进化知识图谱 (ID: mat_004)
5. 实地考察学习活动 (ID: mat_005)"

    echo "   ✅ 搜索完成"
    echo "$SEARCH_RESULTS"
    echo ""
else
    echo "ℹ️  使用AI生成教学内容"
    echo ""
fi

# 步骤3: 生成个性化问题 (ExploraPal能力)
echo "❓ 步骤3: 生成个性化问题..."
echo "   基于布鲁姆分类学的层次化问题设计"

# 模拟问题生成
QUESTIONS="[
    {
        \"content\": \"你能描述一下这块三角龙化石最明显的特征吗？\",
        \"type\": \"观察\",
        \"difficulty\": \"easy\",
        \"purpose\": \"培养观察能力和描述能力\"
    },
    {
        \"content\": \"三角龙的三个头顶角有什么作用？\",
        \"type\": \"推理\",
        \"difficulty\": \"medium\",
        \"purpose\": \"发展科学推理思维\"
    },
    {
        \"content\": \"如果让你设计一个关于三角龙的科学实验，你会怎么做？\",
        \"type\": \"实验设计\",
        \"difficulty\": \"hard\",
        \"purpose\": \"培养科学研究能力\"
    }
]"

echo "   ✅ 生成个性化问题:"
echo "$QUESTIONS" | jq '.[] | "   • \(.content)"' 2>/dev/null || echo "   • 问题生成完成"
echo ""

# 步骤4: 教学教案生成 (TALect MCP服务)
if [ "$MCP_ENABLED" = true ]; then
    echo "📝 步骤4: 教学教案生成..."
    echo "   调用TALect MCP服务生成个性化教案"

    # 构建教案参数
    OBJECTIVES="[
        \"理解三角龙的基本特征和生活习性\",
        \"培养观察古生物化石的能力\",
        \"激发对古生物学的学习兴趣\"
    ]"

    echo "   📋 教学目标:"
    echo "$OBJECTIVES" | jq -r '.[] | "     • \(.")' 2>/dev/null || echo "     • 教案生成中..."

    # 这里应该调用MCP客户端生成教案
    # 模拟教案生成
    LESSON_PLAN="# 个性化教学教案

## 教学目标
- 理解三角龙的基本特征和生活习性
- 培养观察古生物化石的能力
- 激发对古生物学的学习兴趣

## 教学活动建议
- 化石观察与特征分析
- 小组讨论与知识分享
- 创意表达与作品创作

---
*基于学而思教研标准生成*"

    echo "   ✅ 教案生成完成"
    echo ""
else
    echo "ℹ️  使用AI能力生成教学教案"
    echo ""
fi

# 步骤5: 生成学习报告 (ExploraPal能力)
echo "📊 步骤5: 生成学习报告..."
echo "   使用 qwen3-max 进行深度分析和报告生成"

# 模拟报告生成
REPORT_TITLE="三角龙化石观察与学习报告"
REPORT_SUMMARY="通过对三角龙化石的观察和学习，学生深入了解了古生物的基本特征和发展规律。"

echo "   📄 报告标题: $REPORT_TITLE"
echo "   📝 摘要: $REPORT_SUMMARY"
echo "   ✅ 报告生成完成"
echo ""

# 步骤6: 生成个性化学习建议
echo "🎯 步骤6: 生成个性化学习建议..."
echo "   基于学习分析的个性化推荐"

RECOMMENDATIONS="[
    {
        \"type\": \"exploration\",
        \"title\": \"古生物博物馆参观\",
        \"description\": \"组织实地参观，近距离观察更多化石标本\",
        \"difficulty\": \"medium\",
        \"duration\": 120
    },
    {
        \"type\": \"reading\",
        \"title\": \"扩展阅读\",
        \"description\": \"阅读《恐龙世界》相关章节\",
        \"difficulty\": \"easy\",
        \"duration\": 30
    },
    {
        \"type\": \"creative\",
        \"title\": \"创意表达\",
        \"description\": \"绘制三角龙复原图，写一篇短文\",
        \"difficulty\": \"medium\",
        \"duration\": 45
    }
]"

echo "   ✅ 个性化建议:"
echo "$RECOMMENDATIONS" | jq '.[] | "   • \(.title): \(.description)"' 2>/dev/null || echo "   • 建议生成完成"
echo ""

echo "🎉 完整AI学习流程完成!"
echo ""
echo "📈 学习成果总结:"
echo "   • 识别并分析了三角龙化石特征"
echo "   • 完成了个性化问题引导"
echo "   • 生成了学习报告和建议"
if [ "$MCP_ENABLED" = true ]; then
    echo "   • 获得了5个相关教学资源"
    echo "   • 生成了个性化教学教案"
fi
echo ""
echo "💡 ExploraPal的核心价值："
echo "   • 🎯 专注于儿童探索式学习"
echo "   • 🤖 强大的多模态AI能力"
echo "   • 🎨 创意学习体验设计"
echo "   • 📚 集成TALect专业教学资源"
echo ""
echo "🌟 ExploraPal为孩子创造精彩的AI学习体验！"
