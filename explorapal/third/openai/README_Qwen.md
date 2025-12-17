# 阿里云Qwen AI模型集成指南

## 概述

本项目已集成阿里云的Qwen系列AI模型，完全替代了原有的Azure AI/OpenAI模型。Qwen系列模型在中文理解、多模态处理和推理能力方面具有显著优势。

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

## 阿里云语音服务集成

语音功能现在由 **qwen3-omni-flash** 多模态大模型提供，支持：

### 语音转文字 (ASR)
- **支持**: 119种语言文本交互
- **特点**: 实时语音识别、多语言支持、高准确率

### 文字转语音 (TTS)
- **支持**: 20种语言语音交互
- **特点**: 生成类人语音、跨语言精准沟通、多种音色可选

### 配置方式
语音功能通过DashScope API统一配置，无需额外配置。

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
1. **API密钥错误**: 检查API Key是否正确配置
2. **余额不足**: 前往阿里云控制台充值
3. **模型不可用**: 确认模型名称是否正确
4. **超时错误**: 适当调整timeout参数

### 获取帮助
- [DashScope文档](https://help.aliyun.com/zh/dashscope/)
- [Qwen模型介绍](https://qwen.aliyun.com/)
- [阿里云工单](https://workorder.aliyun.com/)
