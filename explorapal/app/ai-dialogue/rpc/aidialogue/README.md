# Protobuf生成代码目录

## ⚠️ 重要提示

此目录中的代码需要通过protobuf生成，**请勿手动编辑**。

## 生成步骤

运行以下命令生成protobuf代码：

```bash
cd explorapal/app/ai-dialogue/rpc
./generate_proto.sh
```

或者手动运行：

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ai-dialogue.proto
```

## 生成的文件

运行protoc后会生成以下文件：
- `aidialogue/ai-dialogue.pb.go` - 消息类型定义
- `aidialogue/ai-dialogue_grpc.pb.go` - gRPC服务定义

生成后即可正常启动服务。
