# AI对话RPC服务

## ⚠️ 重要：生成Protobuf代码

**在运行服务之前，必须先生成protobuf代码！**

### 方法1：使用生成脚本（推荐）

```bash
cd explorapal/app/ai-dialogue/rpc
chmod +x generate_proto.sh
./generate_proto.sh
```

### 方法2：手动运行protoc

```bash
cd explorapal/app/ai-dialogue/rpc

# 安装protoc插件（如果未安装）
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 生成protobuf代码
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ai-dialogue.proto
```

### 方法3：使用goctl（如果已安装）

```bash
cd explorapal/app/ai-dialogue/rpc
goctl rpc protoc ai-dialogue.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## 启动服务

生成protobuf代码后：

```bash
cd explorapal/app/ai-dialogue/rpc
go run aidialogueservice.go
```

## 服务接口

- `AnalyzeImage`: 图片分析
- `GenerateQuestions`: 生成引导问题
- `PolishNote`: 笔记润色
- `GenerateReport`: 生成研究报告

## 故障排除

如果遇到 `package explorapal/app/ai-dialogue/rpc/aidialogue is not in std` 错误：
1. 确保已运行 `generate_proto.sh` 生成protobuf代码
2. 检查 `aidialogue/` 目录下是否有生成的 `.pb.go` 文件
3. 如果存在 `placeholder.go`，删除它（protobuf生成后会替换）
