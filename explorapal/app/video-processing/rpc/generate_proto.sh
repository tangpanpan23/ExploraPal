#!/bin/bash

# 生成视频处理服务的proto文件

cd "$(dirname "$0")"

# 生成Go代码
protoc --go_out=. --go-grpc_out=. video-processing.proto

echo "✅ Proto文件生成完成！"
