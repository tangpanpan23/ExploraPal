# 探索伙伴 MVP 演示指南

## 产品概述

**探索伙伴** 是一个专为儿童设计的AI学习产品，通过"观察-提问-表达"的互动流程，将孩子的兴趣爱好转化为深度研究项目。

**MVP版本**: "发现口袋 - 恐龙篇"

## 核心功能流程

```
观察阶段 → 提问引导 → 表达阶段 → 成果生成
   ↓          ↓          ↓          ↓
 图片分析    AI问题     笔记润色   研究报告
```

## 快速开始

### 1. 启动服务

```bash
# 确保数据库已设置
./setup_database.sh

# 启动所有服务（建议按顺序启动）
go run app/project-management/rpc/projectmanagementservice.go    # 9001
go run app/ai-dialogue/rpc/aidialogueservice.go                 # 9002
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

#### 步骤4：表达阶段 - AI润色笔记

```bash
curl -X POST "http://localhost:9003/api/expression/note/polish" \
  -H "Content-Type: application/json" \
  -d '{
    "raw_content": "三角龙有3只角 头上有骨板 吃植物 很厉害",
    "context_info": "孩子观察恐龙化石后的随手笔记",
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

#### 步骤5：成果生成 - 创建研究报告

```bash
curl -X POST "http://localhost:9003/api/achievement/report/generate" \
  -H "Content-Type: application/json" \
  -d '{
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

### 成果生成
- `POST /api/achievement/report/generate` - 生成研究报告

### 公共接口
- `GET /api/common/ping` - 健康检查

## 服务架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   API Service   │    │ Project Mgmt    │    │  AI Dialogue    │
│   (Port 9003)   │◄──►│   RPC Service   │    │  RPC Service    │
│                 │    │   (Port 9001)   │    │   (Port 9002)   │
│ - RESTful API   │    │                 │    │                 │
│ - 路由转发      │    │ - 项目CRUD      │    │ - AI 图像分析  │
│ - 请求验证      │    │ - 数据存储      │    │ - 问题生成     │
└─────────────────┘    └─────────────────┘    │ - 笔记润色     │
                                              │ - 报告生成     │
                                              └─────────────────┘
```

## 技术栈

- **框架**: Go-Zero
- **数据库**: MySQL
- **缓存**: Redis
- **AI服务**: 阿里云 Qwen (qwen3-vl-plus, qwen-flash, qwen3-max)
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

## 下一步扩展

1. **语音处理**: 集成语音转文字和文字转语音
2. **多媒体支持**: 支持视频分析和生成
3. **个性化学习**: 根据孩子年龄和兴趣推荐内容
4. **社交功能**: 孩子间分享探索成果
5. **家长监控**: 为家长提供学习进度和建议

---

**演示完成！** 🎉

这个MVP展示了完整的AI辅助儿童学习流程，从观察图片开始，通过AI引导生成问题，再到润色表达，最终生成研究报告，帮助孩子将兴趣转化为系统的学习成果。
