# ExploraPal TALect MCP服务配置指南

## 📖 概述

本指南介绍如何配置ExploraPal的TALect MCP服务集成。TALect MCP服务为ExploraPal提供了丰富的教学资源支持能力。

## 🏗️ 系统架构

```
┌─────────────────┐    MCP协议     ┌──────────────────────┐
│   ExploraPal    │◄─────────────►│     TALect MCP       │
│   AI学习引擎    │   JSON-RPC    │   教学资源服务       │
│                 │               │                      │
│ • 图像分析      │               │ • 教学素材搜索       │
│ • 问题生成      │               │ • 教案自动生成       │
│ • 报告创作      │               │ • 个性化推荐         │
│ • 视频创作      │               │ • 学习路径规划       │
└─────────────────┘               └──────────────────────┘
```

## ⚙️ 配置步骤

### 1. MCP集成配置

首先，配置TALect MCP服务连接：

```bash
# 复制配置文件模板
cp mcp_config.yaml.example mcp_config.yaml

# 编辑配置文件
vim mcp_config.yaml
```

关键配置项：

```yaml
mcp:
  enabled: true                    # 启用MCP集成
  base_url: "http://localhost:8080/mcp"  # TALect MCP服务地址
  api_token: "your_api_token"      # API认证令牌（如果需要）

talect:
  subject_mapping:
    science: "science"             # 学科映射
  grade_mapping:
    grade_3: "grade_3"            # 年级映射
```

### 2. AI服务配置

确保ExploraPal的AI服务正确配置：

```yaml
# explorapal/config.yaml
ai:
  tal_mlo_ps:
    app_id: "your_tal_app_id"
    app_key: "your_tal_app_key"
  doubao:
    api_key: "your_doubao_key"
    base_url: "https://api.doubai.com/v1"
```

### 3. 启动TALect MCP服务

```bash
# 进入TALect项目目录
cd ../TALect/future-mcp-server

# 启动MCP服务
go run cmd/server/main.go
```

确认服务运行状态：

```bash
curl http://localhost:8080/health
# 应该返回: {"status": "healthy"}
```

## 🚀 使用方法

### 基础使用

运行增强版学习流程演示：

```bash
cd explorapal

# 基础演示（AI-only模式）
./enhanced_learning_demo.sh dinosaur.jpg

# 指定年级
./enhanced_learning_demo.sh dinosaur.jpg grade_3
```

### 完整功能体验

#### 1. 启用TALect MCP服务

配置TALect MCP服务后，ExploraPal将启用完整的教学资源功能：

```bash
# 检查MCP连接
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {"tools": {}},
      "clientInfo": {"name": "ExploraPal", "version": "1.0.0"}
    }
  }'
```

#### 2. Go代码集成

在Go应用中使用增强版学习流程：

```go
package main

import (
    "fmt"
    "log"

    "github.com/your-org/explorapal/third/openai"
)

func main() {
    // 配置MCP集成
    mcpConfig := &openai.MCPIntegration{
        Enabled:  true,
        BaseURL:  "http://localhost:8080/mcp",
        APIToken: "your_token",
    }

    // 初始化MCP客户端
    err := openai.SetMCPIntegration(mcpConfig)
    if err != nil {
        log.Fatal("MCP初始化失败:", err)
    }

    // 创建AI客户端
    aiClient, err := openai.NewClient(&openai.Config{
        TAL_MLOPS_APP_ID:  "your_app_id",
        TAL_MLOPS_APP_KEY: "your_app_key",
    })
    if err != nil {
        log.Fatal("AI客户端创建失败:", err)
    }

    // 读取图片数据
    imageData, err := ioutil.ReadFile("dinosaur.jpg")
    if err != nil {
        log.Fatal("图片读取失败:", err)
    }

    // 执行增强版学习流程
    result, err := aiClient.EnhancedLearningFlow(
        context.Background(),
        imageData,
        "user123",
        "grade_3",
    )
    if err != nil {
        log.Fatal("学习流程执行失败:", err)
    }

    // 处理结果
    fmt.Printf("图像分析: %+v\n", result.ImageAnalysis)
    fmt.Printf("相关素材: %d 个\n", len(result.RelatedMaterials))
    fmt.Printf("个性化问题: %d 个\n", len(result.PersonalizedQuestions))
    fmt.Printf("学习建议: %d 项\n", len(result.Recommendations))
}
```

## 🎯 核心功能详解

### 六维融合学习流程

1. **AI多模态感知** 🧠
   - 使用 qwen3-vl-plus 进行深度图像分析
   - 识别学习内容和学生认知状态

2. **智能教学资源搜索** 📚
   - 通过MCP调用TALect的教学素材库
   - 基于语义搜索和知识图谱推荐

3. **认知个性化引导** ❓
   - 结合布鲁姆分类学的问题设计
   - 自适应学习路径规划

4. **沉浸式表达创作** 🎨
   - 多模态创作工具支持
   - AI创作助手和Doubao视频生成

5. **标准化教学实施** 📝
   - 基于学而思标准的教案自动生成
   - 5E教学模型的智能适配

6. **学习效果评估** 📊
   - 多维度学习效果量化
   - 个性化干预策略推荐

## 🔧 故障排除

### MCP连接问题

```bash
# 检查MCP服务状态
curl http://localhost:8080/health

# 测试MCP协议
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"ping"}'
```

### AI服务问题

```bash
# 检查AI服务配置
tail -f explorapal/logs/ai_service.log

# 测试AI连接
curl -H "Authorization: Bearer your_token" \
  https://ai-service.tal.com/openai-compatible/v1/models
```

### 常见问题

1. **TALect MCP服务连接失败**
   - 确认TALect服务已启动并可访问
   - 检查网络连接和端口配置
   - 验证API令牌和权限设置

2. **AI服务调用失败**
   - 检查TAL/AppID和AppKey配置
   - 确认网络连接和防火墙设置
   - 查看服务日志了解具体错误

3. **图片处理失败**
   - 确认图片格式支持（JPEG/PNG）
   - 检查图片大小（建议<10MB）
   - 验证base64编码正确性

## 📊 性能监控

### 关键指标

- **响应时间**: MCP调用 < 2秒，AI分析 < 15秒
- **成功率**: 整体流程成功率 > 95%
- **并发处理**: 支持100+并发学习流程

### 监控命令

```bash
# 查看系统状态
./enhanced_learning_demo.sh --status

# 检查MCP连接健康度
curl http://localhost:8080/metrics

# 查看AI服务性能
tail -f explorapal/logs/performance.log
```

## 🚀 扩展开发

### 添加新的MCP工具

在TALect服务中注册新工具：

```go
// 在TALect的tool_registry.go中添加
registry.RegisterTool(&types.ToolDefinition{
    Name:        "custom_educational_tool",
    Description: "自定义教育工具",
    Handler:     handleCustomTool,
    InputSchema: customSchema,
})
```

### 扩展AI能力

在ExploraPal中添加新的AI模型：

```go
// 在openai.go中添加
const ModelCustomAI = "custom-model"

// 添加相应处理函数
func (c *Client) CustomAIFunction(ctx context.Context, input string) (string, error) {
    // 实现自定义AI功能
}
```

## 📞 支持与反馈

- **项目主页**: [TALect Pro Repository]
- **文档中心**: [详细技术文档]
- **问题反馈**: [GitHub Issues]
- **社区讨论**: [开发者论坛]

---

*"通过TALect Pro，我们正在重新定义AI教育的可能性，将ExploraPal的创新精神与TALect的专业积淀完美融合。"*
