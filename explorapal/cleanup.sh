#!/bin/bash

# 清理测试和调试文件脚本
# 执行前请备份重要文件

echo "🧹 清理测试和调试文件..."

# 需要删除的文件列表
FILES_TO_REMOVE=(
    "api_examples.md"
    "demo_flow.sh"
    "test_demo.sh"
    "TROUBLESHOOTING.md"
)

# 需要删除的目录
DIRS_TO_REMOVE=(
    "tools"
)

echo "📋 将删除以下文件："
for file in "${FILES_TO_REMOVE[@]}"; do
    if [ -f "$file" ]; then
        echo "  - $file"
    fi
done

echo "📋 将删除以下目录："
for dir in "${DIRS_TO_REMOVE[@]}"; do
    if [ -d "$dir" ]; then
        echo "  - $dir/"
    fi
done

echo ""
read -p "⚠️  确定要删除这些文件吗？(y/N): " -n 1 -r
echo ""

if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "🗑️  开始删除..."

    # 删除文件
    for file in "${FILES_TO_REMOVE[@]}"; do
        if [ -f "$file" ]; then
            rm -f "$file"
            echo "✅ 删除文件: $file"
        fi
    done

    # 删除目录
    for dir in "${DIRS_TO_REMOVE[@]}"; do
        if [ -d "$dir" ]; then
            rm -rf "$dir"
            echo "✅ 删除目录: $dir/"
        fi
    done

    echo ""
    echo "🎉 清理完成！"
    echo ""
    echo "📝 已清理的内容："
    echo "  - 旧版API示例文档"
    echo "  - 测试和演示脚本"
    echo "  - 故障排除指南"
    echo "  - 调试工具脚本"
    echo ""
    echo "📚 保留的核心文档："
    echo "  - README.md (项目说明)"
    echo "  - MVP_DEMO.md (演示指南)"
    echo "  - MULTIMODAL_DEMO.md (多模态功能)"
else
    echo "❌ 操作已取消"
fi
