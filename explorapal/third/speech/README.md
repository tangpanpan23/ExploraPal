# 阿里云语音服务集成指南

## 概述

本项目集成了阿里云智能语音服务，提供语音转文字(ASR)和文字转语音(TTS)功能，完全替代了原有的Azure Speech Services。

## 服务介绍

### 1. 语音转文字 (ASR - Automatic Speech Recognition)
- **实时语音识别**: 支持实时流式语音识别
- **单句识别**: 支持短音频文件一次性识别
- **多语言支持**: 中文、英文、日语、韩语等
- **噪音过滤**: 自动过滤环境噪音，提高识别准确率
- **情感识别**: 可选的情感识别功能

### 2. 文字转语音 (TTS - Text To Speech)
- **多音色选择**: 多种音色，包括儿童友好音色
- **语速调节**: 支持语速调整(-500到500)
- **音调调节**: 支持音调调整(-500到500)
- **多种格式**: 支持WAV、MP3、PCM等格式输出

## 配置步骤

### 1. 开通阿里云智能语音服务
1. 访问 [阿里云智能语音服务](https://ai.aliyun.com/nls)
2. 开通智能语音服务
3. 创建应用，获取AppKey
4. 获取AccessKey ID和AccessKey Secret

### 2. 配置参数
```yaml
# 语音服务配置
SpeechService:
  AccessKeyId: "your-access-key-id"          # 阿里云AccessKey ID
  AccessKeySecret: "your-access-key-secret" # 阿里云AccessKey Secret
  AppKey: "your-app-key"                     # 语音服务AppKey
  Region: "cn-shanghai"                      # 服务地域
```

### 3. 验证配置
```bash
# 测试语音服务连接（需要先实现测试代码）
```

## 使用方法

### 语音转文字
```go
client := speech.NewClient(&speech.Config{
    AccessKeyId:     "your-access-key-id",
    AccessKeySecret: "your-access-key-secret",
    AppKey:          "your-app-key",
    Region:          "cn-shanghai",
})

// 语音转文字
text, err := client.SpeechToText(ctx, audioData, "wav", 16000, "zh-CN")
if err != nil {
    log.Fatal(err)
}
fmt.Println("识别结果:", text)
```

### 文字转语音
```go
// 文字转语音
audioData, err := client.TextToSpeech(ctx, "你好，这是测试文本", "xiaoyun", "mp3")
if err != nil {
    log.Fatal(err)
}

// 保存音频文件
err = ioutil.WriteFile("output.mp3", audioData, 0644)
if err != nil {
    log.Fatal(err)
}
```

## API接口说明

### 语音识别接口
- **URL**: `https://nls-gateway-{region}.aliyuncs.com/stream/v1/asr`
- **方法**: POST
- **认证**: AccessKey签名 + AppKey
- **数据格式**: JSON

### 语音合成接口
- **URL**: `https://nls-gateway-{region}.aliyuncs.com/stream/v1/tts`
- **方法**: POST
- **认证**: AccessKey签名 + AppKey
- **数据格式**: JSON

## 功能参数

### ASR参数
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| format | string | 是 | 音频格式：wav, mp3, m4a, flv, amr, webm |
| sample_rate | int | 是 | 采样率：8000, 16000 |
| language | string | 否 | 语言代码：zh-CN, en-US, ja-JP, ko-KR |
| audio_base64 | string | 是 | base64编码的音频数据 |

### TTS参数
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| text | string | 是 | 待合成的文本 |
| voice | string | 否 | 音色：xiaoyun, xiaogang, etc. |
| format | string | 否 | 输出格式：wav, mp3, pcm |
| speech_rate | int | 否 | 语速：-500~500 |
| pitch_rate | int | 否 | 音调：-500~500 |

## 最佳实践

### 1. 音频质量优化
- **采样率**: 推荐使用16000Hz
- **格式**: 优先使用WAV格式，无损质量
- **长度**: 单次识别建议不超过60秒
- **噪音**: 尽量在安静环境中录音

### 2. 性能优化
- **连接复用**: 复用HTTP客户端连接
- **并发控制**: 控制并发请求数量
- **错误重试**: 实现合理的重试机制
- **超时设置**: 根据音频长度设置合适超时时间

### 3. 成本控制
- **按需使用**: 只在需要时调用语音服务
- **批量处理**: 对于大量音频，考虑批量处理
- **缓存机制**: 缓存常用语音合成结果

## 错误处理

### 常见错误码
- **400**: 请求参数错误
- **401**: 认证失败
- **403**: 权限不足
- **429**: 请求频率过高
- **500**: 服务内部错误

### 错误处理策略
```go
text, err := client.SpeechToText(ctx, audioData, "wav", 16000, "zh-CN")
if err != nil {
    // 分类处理不同错误
    if strings.Contains(err.Error(), "认证失败") {
        // 检查AccessKey配置
    } else if strings.Contains(err.Error(), "参数错误") {
        // 检查音频参数
    } else {
        // 其他错误，重试或降级处理
    }
}
```

## 监控和告警

### 关键指标监控
- **识别准确率**: 监控语音识别准确性
- **响应时间**: 监控API响应延迟
- **成功率**: 监控API调用成功率
- **使用量**: 监控每日调用次数和费用

### 告警规则
- API调用失败率 > 5%
- 平均响应时间 > 30秒
- 费用超出预算阈值

## 安全注意事项

### 1. 音频数据保护
- 传输过程使用HTTPS加密
- 音频数据不落地存储
- 及时清理临时文件

### 2. 访问控制
- 使用最小权限原则
- 定期轮换AccessKey
- 监控异常访问行为

### 3. 合规要求
- 遵守当地隐私保护法规
- 获得用户录音授权
- 支持用户数据删除请求

## 技术支持

- [阿里云智能语音服务文档](https://help.aliyun.com/document_detail/90723.html)
- [阿里云工单系统](https://workorder.aliyun.com/)
- [阿里云技术论坛](https://developer.aliyun.com/)
