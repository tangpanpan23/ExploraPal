# AI对话RPC服务

## 生成Protobuf代码

在运行服务之前，需要先生成protobuf代码：

```bash
cd explorapal/app/ai-dialogue/rpc

# 使用protoc生成Go代码
protoc --go_out=. --go-grpc_out=. ai-dialogue.proto

# 或者使用goctl生成（推荐）
goctl rpc protoc ai-dialogue.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## 启动服务

```bash
cd explorapal/app/ai-dialogue/rpc
go run aidialogueservice.go
```

## 服务接口

- `AnalyzeImage`: 图片分析
- `GenerateQuestions`: 生成引导问题
- `PolishNote`: 笔记润色
- `GenerateReport`: 生成研究报告
