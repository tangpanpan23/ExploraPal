# AI对话RPC服务设置指南

## ⚠️ 重要：首次运行前必须生成Protobuf代码

在运行 `aidialogueservice.go` 之前，**必须先生成protobuf代码**，否则会出现以下错误：

```
package explorapal/app/ai-dialogue/rpc/aidialogue is not in std
```

## 快速开始

### 步骤1：生成Protobuf代码

```bash
cd explorapal/app/ai-dialogue/rpc

# 方法1：使用生成脚本（推荐）
chmod +x generate_proto.sh
./generate_proto.sh

# 方法2：手动运行protoc
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ai-dialogue.proto
```

### 步骤2：验证生成的文件

生成成功后，`aidialogue/` 目录下应该有以下文件：
- `ai-dialogue.pb.go` - 消息类型定义
- `ai-dialogue_grpc.pb.go` - gRPC服务定义

### 步骤3：启动服务

```bash
go run aidialogueservice.go
```

## 依赖安装

如果遇到 `protoc: command not found` 错误：

### macOS
```bash
brew install protobuf
```

### Linux
```bash
sudo apt-get install protobuf-compiler
```

### 安装Go插件
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 故障排除

### 错误：`package ... is not in std`

**原因**：protobuf代码未生成

**解决**：运行 `./generate_proto.sh` 生成代码

### 错误：`protoc: command not found`

**原因**：protoc未安装

**解决**：按照上面的"依赖安装"步骤安装protoc

### 错误：`protoc-gen-go: program not found`

**原因**：Go插件未安装或PATH未设置

**解决**：
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```
