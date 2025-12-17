# 内部AI服务集成指南 (基于阿里云Qwen)

## 概述

本项目通过TAL MLOps平台调用内部AI服务，使用阿里云Qwen系列大模型，完全替代了原有的Azure AI/OpenAI模型。内部服务提供统一的OpenAI兼容API接口，简化了集成复杂度。

## 调用方式

### API调用格式
```bash
curl --location 'http://ai-service.tal.com/openai-compatible/v1/chat/completions' \
--header "Authorization: Bearer ${TAL_MLOPS_APP_ID}:${TAL_MLOPS_APP_KEY}" \
--header 'Content-Type: application/json' \
--data '{
    "model": "qwen-flash",
    "messages": [
        {
            "role": "system",
            "content": "You are a helpful assistant."
        },
        {
            "role": "user",
            "content": "你是谁？"
        }
    ]
}'
```

### 配置参数
- `TAL_MLOPS_APP_ID`: TAL MLOps应用ID (配置文件中的唯一标识)
- `TAL_MLOPS_APP_KEY`: TAL MLOps应用密钥 (配置文件中的唯一标识)
- `BaseURL`: 内部AI服务端点 (默认: http://ai-service.tal.com/openai-compatible/v1)

## 推荐模型配置

### 1. 图像分析 (ModelImageAnalysis)
- **模型**: `qwen3-vl-plus`
- **上下文长度**: 256K
- **优势**:
  - 视觉智能体能力达到世界顶尖水平
  - 支持超长视频理解
  - 全面升级视觉coding、空间感知、多模态思考
- **适用场景**: 儿童照片分析、AR信息生成、视觉学习引导

### 2. 文本生成 (ModelTextGeneration)
- **模型**: `qwen-flash`
- **上下文长度**: 1048.576K
- **优势**:
  - 思考模式和非思考模式的有效融合
  - 复杂推理类任务性能优秀
  - 指令遵循、文本理解能力显著提高
- **适用场景**: 问题生成、笔记润色、儿童友好表达

### 3. 复杂推理 (ModelAdvancedReasoning)
- **模型**: `qwen3-max`
- **上下文长度**: 256K
- **优势**:
  - 智能体编程与工具调用专项升级
  - 达到领域SOTA水平
  - 适配复杂的智能体需求
- **适用场景**: 研究报告生成、深度学习分析、项目总结

### 4. 多模态语音交互 (ModelVoiceInteraction)
- **模型**: `qwen3-omni-flash`
- **上下文长度**: 48K
- **优势**:
  - 基于Thinker-Talker混合专家架构
  - 支持文本、图像、音频、视频的高效理解与语音生成
  - 119种语言文本交互，20种语言语音交互
  - 生成类人语音，实现跨语言精准沟通
  - 强大指令跟随与系统提示定制功能
- **适用场景**: 语音转文字、文字转语音、多模态对话、语音助手

## 配置步骤

### 1. 开通DashScope服务
1. 访问 [阿里云DashScope](https://dashscope.aliyuncs.com/)
2. 注册/登录阿里云账号
3. 开通DashScope服务
4. 获取API Key

### 2. 配置API密钥
```yaml
DashScope:
  APIKey: "your-dashscope-api-key-here"
  BaseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1"
  Timeout: 30
  MaxTokens: 2000
  Temperature: 0.7
```

### 3. 验证配置
```bash
# 测试API连接
curl -X POST "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen-flash",
    "messages": [
      {"role": "user", "content": "Hello, test message"}
    ]
  }'
```

## 模型特点对比

| 任务类型 | Qwen模型 | 优势 | 适用儿童学习场景 |
|---------|---------|------|------------------|
| 图像分析 | qwen3-vl-plus | 超长视频理解，空间感知强 | 观察学习，AR增强 |
| 问题生成 | qwen-flash | 思考模式灵活，推理优秀 | 个性化引导问题 |
| 内容创作 | qwen3-max | 工具调用强，结构化输出 | 研究报告，学习总结 |
| 语音交互 | qwen3-omni-flash | 多模态语音处理 | 语音转文字，文字转语音 |

## 语音功能集成

语音功能通过 **qwen3-omni-flash** 多模态大模型提供，同样通过内部AI服务调用：

### 语音转文字 (ASR)
- **支持**: 119种语言文本交互
- **特点**: 实时语音识别、多语言支持、高准确率

### 文字转语音 (TTS)
- **支持**: 20种语言语音交互
- **特点**: 生成类人语音、跨语言精准沟通、多种音色可选

### 调用方式
语音功能使用相同的API端点和认证方式，通过指定model为"qwen3-omni-flash"即可使用语音功能。

## 使用示例

### Go代码调用示例
```go
package main

import (
    "context"
    "fmt"
    "explorapal/third/openai"
)

func main() {
    // 初始化客户端
    config := &openai.Config{
        TAL_MLOPS_APP_ID:  "your-app-id",
        TAL_MLOPS_APP_KEY: "your-app-key",
        BaseURL:           "http://ai-service.tal.com/openai-compatible/v1",
    }

    client := openai.NewClient(config)

    // 调用文本生成功能
    messages := []openai.ChatCompletionMessage{
        {
            Role:    openai.ChatMessageRoleSystem,
            Content: "你是一个儿童教育助手",
        },
        {
            Role:    openai.ChatMessageRoleUser,
            Content: "帮我生成一个关于恐龙的简单问题",
        },
    }

    resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
        Model:    openai.ModelTextGeneration,
        Messages: messages,
    })

    if err != nil {
        fmt.Printf("调用失败: %v\n", err)
        return
    }

    fmt.Printf("AI回复: %s\n", resp.Choices[0].Message.Content)
}
```

## 注意事项

### 1. API调用限制
- 免费额度有限，建议及时充值
- 注意并发请求限制
- 监控API使用量和费用

### 2. 模型选择建议
- 儿童教育场景优先选择支持思考模式的模型
- 复杂推理任务使用qwen3-max
- 快速响应场景使用qwen-flash

### 3. 最佳实践
- 根据具体任务选择最合适的模型
- 合理设置temperature参数控制创造性
- 利用模型的思考模式进行深度学习引导

## 故障排除

### 常见问题
1. **认证失败**: 检查TAL_MLOPS_APP_ID和TAL_MLOPS_APP_KEY是否正确配置
2. **网络连接**: 确认内部AI服务 (ai-service.tal.com) 网络可达
3. **权限不足**: 联系TAL MLOps管理员确认账户权限
4. **模型不可用**: 确认模型名称是否正确，当前支持的模型包括qwen3-vl-plus、qwen-flash、qwen3-max、qwen3-omni-flash
5. **超时错误**: 适当调整timeout参数，或检查内部服务负载情况

### 获取帮助
- 联系TAL MLOps平台管理员
- 查看内部AI服务文档
- 提交工单到TAL技术支持
