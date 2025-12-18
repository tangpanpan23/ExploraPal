#!/bin/bash

# 生成语音处理服务的proto文件

cd "$(dirname "$0")"

# 生成Go代码
protoc --go_out=. --go-grpc_out=. audio-processing.proto

echo "✅ Proto文件生成完成！"
