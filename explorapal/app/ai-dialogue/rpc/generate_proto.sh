#!/bin/bash

# 生成AI对话服务的protobuf代码

echo "🔧 生成AI对话服务的protobuf代码..."

cd "$(dirname "$0")"

# 检查protoc是否安装
if ! command -v protoc &> /dev/null; then
    echo "❌ protoc未安装，请先安装protoc"
    echo "安装方法:"
    echo "  macOS: brew install protobuf"
    echo "  Linux: sudo apt-get install protobuf-compiler"
    exit 1
fi

# 安装protoc插件
echo "📦 安装protoc插件..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 确保PATH包含Go bin目录
export PATH="$PATH:$(go env GOPATH)/bin"

# 创建输出目录
mkdir -p aidialogue

# 删除占位符文件（如果存在）
if [ -f "aidialogue/placeholder.go" ]; then
    echo "🗑️  删除占位符文件..."
    rm aidialogue/placeholder.go
fi

# 生成protobuf代码
echo "📦 生成Go代码..."
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ai-dialogue.proto

if [ $? -eq 0 ]; then
    echo "✅ Protobuf代码生成成功！"
    echo "📁 生成的文件："
    ls -la aidialogue/*.pb.go 2>/dev/null || echo "   (文件列表)"
    echo ""
    echo "🚀 现在可以运行服务了："
    echo "   go run aidialogueservice.go"
else
    echo "❌ Protobuf代码生成失败！"
    echo "请检查："
    echo "  1. protoc是否已正确安装"
    echo "  2. protoc-gen-go和protoc-gen-go-grpc插件是否已安装"
    echo "  3. Go环境变量是否正确设置"
    exit 1
fi
