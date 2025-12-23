# 探索伙伴 MVP 演示指南

## 产品概述

**探索伙伴** 是一个专为儿童设计的AI学习产品，通过"观察-提问-表达"的互动流程，将孩子的兴趣爱好转化为深度研究项目。

**MVP版本**: "发现口袋 - 恐龙篇"

## 核心功能流程

```
观察阶段 → 提问引导 → 表达阶段 → 成果生成 → 多媒体创作
   ↓          ↓          ↓          ↓          ↓
 图片分析    AI问题     语音交互   研究报告   视频创作
   ↓          ↓          ↓          ↓          ↓
 视频分析    深度思考   内容润色   纪录片     AI视频
```

## 快速开始

### 1. 启动服务

```bash
# 确保数据库已设置
./setup_database.sh

# 启动所有服务（建议按顺序启动）
go run app/project-management/rpc/projectmanagementservice.go    # 9001
go run app/ai-dialogue/rpc/aidialogueservice.go                 # 9002
go run app/audio-processing/rpc/main.go                         # 9004
go run app/video-processing/rpc/main.go                         # 9005
go run app/api/api.go                                           # 9003
```

### 2. 验证服务状态

```bash
# 检查API服务
curl http://localhost:9003/api/common/ping
```

## 完整演示流程

### 场景：小明探索恐龙世界

#### 步骤1：创建探索项目

```bash
curl -X POST "http://localhost:9003/api/project/create" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "title": "小明的恐龙探索之旅",
    "description": "跟随小明一起探索古老的恐龙世界",
    "category": "dinosaur",
    "tags": ["恐龙", "古生物", "探索"]
  }'
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "project_id": 1,
    "project_code": "P202412180001",
    "status": "active"
  }
}
```

#### 步骤2：观察阶段 - 分析恐龙图片

```bash
curl -X POST "http://localhost:9003/api/observation/image/recognize" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "image_url": "https://example.com/dinosaur-fossil.jpg",
    "prompt": "请分析这张恐龙化石图片，识别恐龙种类，描述特征",
    "category": "dinosaur"
  }'
```

**响应示例**:
```json
{
  "object_name": "三角龙化石",
  "category": "dinosaur",
  "confidence": 0.95,
  "description": "这是一块保存完好的三角龙头骨化石，三角龙是白垩纪晚期的草食性恐龙，以头部的大骨板和角著称。",
  "key_features": [
    "头部有三只角",
    "颈部有大骨板",
    "牙齿呈叶片状"
  ],
  "scientific_name": "Triceratops"
}
```

#### 步骤3：提问引导 - 生成探索问题

```bash
curl -X POST "http://localhost:9003/api/questioning/questions/generate" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "observation_id": 1,
    "context_info": "小明观察到了一块三角龙化石，上面有三只角和骨板",
    "category": "dinosaur",
    "user_age": 8
  }'
```

**响应示例**:
```json
{
  "questions": [
    {
      "content": "你觉得三角龙的角有什么用处？",
      "type": "reasoning",
      "difficulty": "intermediate",
      "purpose": "培养推理能力和想象力"
    },
    {
      "content": "我们可以从化石上看到三角龙吃什么吗？",
      "type": "observation",
      "difficulty": "basic",
      "purpose": "学习观察细节"
    },
    {
      "content": "如果我们能见到活的三角龙，你最想做什么实验？",
      "type": "experiment",
      "difficulty": "advanced",
      "purpose": "激发科学探索精神"
    }
  ]
}
```

#### 步骤4：表达阶段 - 语音转文字

```bash
curl -X POST "http://localhost:9003/api/expression/speech/text" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "audio_data": "UklGRnoGAABXQVZFZm10IAAAAAEAAQARAAAAEAAAAAEACABkYXRhAgAAAAEA",
    "audio_format": "wav",
    "language": "zh-CN"
  }'
```

**响应示例**:
```json
{
  "text": "三角龙有三只角，脖子上有骨板，是草食性恐龙",
  "confidence": 0.95,
  "language": "zh-CN",
  "duration": 3.2
}
```

#### 步骤5：表达阶段 - AI润色笔记

```bash
curl -X POST "http://localhost:9003/api/expression/note/polish" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "raw_content": "三角龙有3只角 头上有骨板 吃植物 很厉害",
    "content_type": "speech",
    "context_info": {
      "observation_results": "三角龙化石，三只角，骨板",
      "previous_answers": "防御，草食性",
      "project_category": "dinosaur"
    },
    "category": "dinosaur",
    "user_age": 8
  }'
```

**响应示例**:
```json
{
  "title": "三角龙探秘笔记",
  "summary": "小明学习了三角龙的基本特征和生活习性",
  "key_points": [
    "三角龙有三只角和颈部骨板",
    "是草食性恐龙",
    "生活在白垩纪晚期"
  ],
  "formatted_text": "今天我学习了三角龙！三角龙是一种非常有趣的恐龙，它头上长着三只角，脖子后面有大大的骨板。这些特征让三角龙看起来很威武！三角龙主要吃植物，是白垩纪晚期最常见的恐龙之一。",
  "suggestions": [
    "可以添加三角龙的生活环境描述",
    "可以画一幅三角龙的画像"
  ]
}
```

#### 步骤6：成果生成 - 创建研究报告

```bash
curl -X POST "http://localhost:9003/api/achievement/report/generate" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "project_data": "小明探索了三角龙，学习了恐龙的特征、生活习性和灭绝原因",
    "category": "dinosaur"
  }'
```

**响应示例**:
```json
{
  "title": "小明的三角龙探索报告",
  "abstract": "通过观察化石图片和AI引导提问，小明深入了解了三角龙的生活习性",
  "content": "详细的研究报告内容...",
  "conclusion": "三角龙是白垩纪最有趣的恐龙之一",
  "next_steps": [
    "探索其他种类的恐龙",
    "学习恐龙灭绝的原因",
    "制作恐龙模型"
  ]
}
```

#### 步骤6：查看项目列表和详情

```bash
# 获取项目列表
curl -X POST "http://localhost:9003/api/project/list" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "category": "dinosaur"
  }'

# 获取项目详情
curl -X POST "http://localhost:9003/api/project/detail" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1
  }'

# 完成项目
curl -X POST "http://localhost:9003/api/project/status/update" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "status": "completed"
  }'
```

#### 步骤6：语音交互 - 文字转语音播报

```bash
# 将研究总结转换为语音播报
curl -X POST "http://localhost:9003/api/audio/text-to-speech" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "text": "小明通过探索三角龙，学习到了恐龙的基本特征和生活习性。原来三角龙是草食性恐龙，有三只角和颈部骨板来保护自己。",
    "voice": "female",
    "language": "zh-CN",
    "speed": 1.0
  }'
```

**响应示例**:
```json
{
  "status": 200,
  "msg": "文字转语音成功",
  "audio_data": "UklGRnoGAABXQVZFZm10IBAAAAABAAEAQB8AAEAfAAABAAgAZGF0YQoGAACBhYqFbF1fdJivrJBhNjVgodDbq2EcBj+a2/LDciUFLIHO8tiJNwgZaLvt559NEAxQp+PwtmMcBjiR1/LMeSwFJHfH8N2QQAoUXrTp66hVFApGn+DyvmQbBzeR1/LNeS0F",
  "format": "wav",
  "duration": 12.5,
  "expression_id": 0
}
```

#### 步骤7：多媒体创作 - AI视频生成

```bash
# 使用豆包Doubao-Seedance-1.0-lite-i2v模型生成教学视频
# 输入：用户原始图片 + AI润色后的描述
curl -X POST "http://localhost:9003/api/achievement/video/generate" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "image_data": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wAARCAABAAEDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAv/xAAhEAACAQMDBQAAAAAAAAAAAAABAgMABAUGIWGRkqGx0f/EABUBAQEAAAAAAAAAAAAAAAAAAAMF/8QAGhEAAgIDAAAAAAAAAAAAAAAAAAECEgMRkf/aAAwDAQACEQMRAD8AltJagyeH0AthI5xdrLcNM91BF5pX2HaH9bcfaSXWGaRmknyJckliyjqTzSlT54b6bk+h0R+IRjWjBqO6O2mhP//Z",
    "prompt": "小葫芦观察到了一只三角龙化石，上面有三只角和坚硬的骨板，看起来非常威武。这是白垩纪时期的古老生物，生活在大约7000万年前。",
    "style": "educational",
    "duration": 60.0,
    "scenes": [
      "三角龙外形介绍",
      "生活习性展示",
      "生存环境再现"
    ],
    "voice": "female",
    "language": "zh-CN"
  }'
```

**响应示例**:
```json
{
  "status": 200,
  "msg": "视频生成成功",
  "video_data": "AAAFMHN1Ym1oYXJ0YmVhdAFtZXRhZGF0YQABAA...",
  "format": "mp4",
  "duration": 60.0,
  "metadata": {
    "title": "三角龙探秘视频",
    "description": "AI生成的三角龙教学视频",
    "scenes": [
      "三角龙外形介绍",
      "生活习性展示",
      "生存环境再现"
    ],
    "audio_language": "zh-CN",
    "resolution": "1920x1080"
  },
  "achievement_id": 0
}
```

#### 步骤8：内容分析 - 深度视频理解

```bash
# 分析用户上传的恐龙视频
curl -X POST "http://localhost:9003/api/achievement/video/analyze" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "video_data": "AAAFMHN1Ym1oYXJ0YmVhdAFtZXRhZGF0YQABAA...",
    "video_format": "mp4",
    "analysis_type": "content",
    "duration": 30.0
  }'
```

**响应示例**:
```json
{
  "status": 200,
  "msg": "视频分析成功",
  "video_analysis": {
    "scenes": [
      {
        "timestamp": 0.0,
        "scene_type": "educational",
        "confidence": 0.92,
        "description": "教室场景，老师在讲解恐龙知识"
      }
    ],
    "objects": [
      {
        "timestamp": 5.0,
        "object_name": "三角龙模型",
        "confidence": 0.88,
        "bbox": {
          "x": 200,
          "y": 150,
          "width": 400,
          "height": 300
        }
      }
    ],
    "texts": [
      {
        "timestamp": 10.0,
        "text": "三角龙Triceratops",
        "language": "zh-CN",
        "confidence": 0.95,
        "bbox": {
          "x": 100,
          "y": 50,
          "width": 300,
          "height": 60
        }
      }
    ],
    "audio": [
      {
        "timestamp": 15.0,
        "transcription": "三角龙是白垩纪晚期的恐龙",
        "language": "zh-CN",
        "confidence": 0.91
      }
    ],
    "summary": {
      "title": "恐龙纪录片分析",
      "description": "视频主要介绍了三角龙的特征和习性",
      "keywords": ["三角龙", "恐龙", "白垩纪", "草食性"],
      "category": "educational",
      "duration": 30.0
    }
  },
  "achievement_id": 0
}
```

## API 接口列表

### 项目管理
- `POST /api/project/create` - 创建项目
- `POST /api/project/list` - 获取项目列表
- `POST /api/project/detail` - 获取项目详情
- `POST /api/project/status/update` - 更新项目状态

### 观察阶段
- `POST /api/observation/image/recognize` - 识别图片内容

### 提问引导
- `POST /api/questioning/questions/generate` - 生成引导问题

### 表达阶段
- `POST /api/expression/note/polish` - AI润色笔记

### 语音处理
- `POST /api/audio/text-to-speech` - 文字转语音

### 成果生成
- `POST /api/achievement/report/generate` - 生成研究报告
- `POST /api/achievement/video/analyze` - 视频内容分析
- `POST /api/achievement/video/generate` - AI视频生成

### 公共接口
- `GET /api/common/ping` - 健康检查

## 服务架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   API Service   │    │ Project Mgmt    │    │  AI Dialogue    │    │ Audio Process   │
│   (Port 9003)   │◄──►│   RPC Service   │    │  RPC Service    │    │  RPC Service    │
│                 │    │   (Port 9001)   │    │   (Port 9002)   │    │   (Port 9004)   │
│ - RESTful API   │    │                 │    │                 │    │                 │
│ - 路由转发      │    │ - 项目CRUD      │    │ - AI 图像分析  │    │ - 语音转文字   │
│ - 请求验证      │    │ - 数据存储      │    │ - 问题生成     │    │ - 文字转语音   │
└─────────────────┘    └─────────────────┘    │ - 笔记润色     │    └─────────────────┘
                                              │ - 报告生成     │    ┌─────────────────┐
                                              └─────────────────┘    │ Video Process   │
                                                                     │  RPC Service    │
                                                                     │   (Port 9005)   │
                                                                     │                 │
                                                                     │ - 视频分析      │
                                                                     │ - 视频生成      │
                                                                     └─────────────────┘
```

## 技术栈

- **框架**: Go-Zero
- **数据库**: MySQL
- **缓存**: Redis
- **AI服务**: 阿里云 Qwen (qwen3-vl-plus, qwen-flash, qwen3-max, qwen3-omni-flash, qwen-vl-plus)
- **通信**: gRPC (内部) + RESTful API (外部)

## AI服务降级机制

### 🔄 自动降级
当外部AI服务不可用时，系统会自动降级到模拟响应模式：

- **图像分析**: 根据图片URL关键词提供相应的分析结果
- **问题生成**: 返回预设的教育性问题库
- **笔记润色**: 保持原始内容并添加基本格式化
- **报告生成**: 生成模板化的研究报告结构

### 📊 降级示例

```bash
# 即使AI服务不可用，API仍会返回合理的结果
curl -X POST "http://localhost:9003/api/observation/image/recognize" \
  -H "Content-Type: application/json" \
  -d '{
    "image_url": "https://example.com/dinosaur-fossil.jpg",
    "category": "dinosaur"
  }'

# 返回模拟响应
{
  "object_name": "恐龙化石",
  "category": "dinosaur",
  "confidence": 0.8,
  "description": "这看起来像是一块恐龙化石，包含了古代生物的遗骸。",
  "key_features": ["骨骼结构", "化石纹理", "年代久远"],
  "scientific_name": "恐龙类"
}
```

这样确保在任何环境下都能进行完整的MVP演示！

## 已实现功能扩展

### 语音处理功能 ✅
- **语音转文字**: 支持音频文件转文字，支持多种格式
- **文字转语音**: 支持文字转语音播报，支持多种语音和语速

### 多媒体支持功能 ✅
- **视频分析**: 深度分析视频内容，提取场景、物体、情感、文字、音频等多维度信息
- **视频生成**: 基于脚本自动生成AI视频，支持教育、故事、动画等多种风格

## 下一步扩展

1. **个性化学习**: 根据孩子年龄和兴趣推荐内容
2. **社交功能**: 孩子间分享探索成果
3. **家长监控**: 为家长提供学习进度和建议
4. **实时协作**: 支持多人同时探索同一个主题
5. **离线模式**: 支持网络不稳定环境下的学习

---

**演示完成！** 🎉

这个MVP展示了完整的AI辅助儿童学习生态系统：

#### 🔍 核心学习流程
- **观察阶段**: 图像识别 + 视频内容分析
- **提问引导**: AI生成个性化学习问题
- **表达阶段**: 语音交互 + 内容创作
- **成果生成**: 多媒体内容创作

#### 🎵 多媒体创新
- **语音交互**: 语音转文字 + 文字转语音
- **视频创作**: AI视频生成 + 深度内容分析
- **沉浸式学习**: 视觉、听觉、交互多维体验

#### 🤖 AI能力展示
- **多模态理解**: 文本、图像、语音、视频全覆盖
- **个性化学习**: 根据儿童年龄和兴趣定制内容
- **创作赋能**: 从被动学习到主动创作的转变

探索伙伴不仅是学习工具，更是儿童创意表达的AI伙伴！🚀✨
