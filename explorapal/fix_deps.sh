#!/bin/bash

echo "🔧 修复Go模块依赖..."

cd "$(dirname "$0")"

# 下载缺失的依赖
echo "📦 下载依赖模块..."
go mod download github.com/go-sql-driver/mysql
go mod download github.com/zeromicro/go-zero
go mod download github.com/sashabaranov/go-openai

# 整理依赖并生成go.sum
echo "📋 整理依赖..."
go mod tidy

# 验证
echo "✅ 验证依赖..."
go mod verify

echo "✅ 依赖修复完成！"
echo ""
echo "现在可以运行: go run migrate.go"
