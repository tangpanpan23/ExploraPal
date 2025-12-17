# 探索伙伴 (ExploraPal) - AI学习产品

## 项目简介

探索伙伴是一个基于AI的儿童学习产品，采用"从教知识点转向辅助探索"的理念，帮助孩子通过多模态AI技术，将兴趣爱好转化为深度研究项目。

## 核心理念

- **从填充式学习到探索式学习**：不直接灌输知识，而是通过AI引导孩子自主探索
- **兴趣驱动**：从孩子的真实兴趣出发（如恐龙、火箭、Minecraft）
- **多模态交互**：支持图像识别、语音处理、AI对话
- **项目式学习**：将探索过程组织为完整的"研究项目"

## MVP功能 - 发现口袋：恐龙篇

### 核心流程：观察 → 提问 → 表达

1. **观察阶段**
   - 孩子拍照上传恐龙图片
   - AI图像识别分析恐龙特征
   - 生成AR增强信息和观察建议

2. **提问引导**
   - 基于观察结果生成个性化问题
   - 问题类型：观察、推理、实验、比较
   - AI提供答案和延伸思考

3. **表达阶段**
   - 语音转文字记录想法
   - AI润色生成研究笔记
   - 自动生成分享卡片

## 技术架构

### 后端技术栈
- **框架**: Go + go-zero 微服务框架
- **数据库**: MySQL
- **缓存**: Redis
- **AI服务**: OpenAI GPT-4 Vision, Claude
- **架构**: API服务 + RPC服务 微服务架构

### 项目结构
```
explorapal/
├── app/
│   ├── api/                    # REST API服务
│   ├── project-management/rpc/ # 项目管理RPC服务
│   ├── image-recognition/rpc/  # 图像识别RPC服务
│   ├── audio-processing/rpc/   # 语音处理RPC服务
│   └── ai-dialogue/rpc/        # AI对话RPC服务
├── common/                     # 通用工具
├── constant/                   # 常量定义
├── database/migrations/        # 数据库迁移
├── pkg/                        # 通用包
├── third/                      # 第三方服务集成
│   ├── openai/                 # OpenAI客户端
│   ├── claude/                 # Claude客户端
│   └── azure-speech/           # 语音服务
└── storage/                    # 文件存储
```

## 核心数据模型

- **用户(Users)**: 用户基本信息
- **项目(Projects)**: 探索项目
- **观察记录(Observations)**: 图像分析结果
- **问题记录(Questions)**: AI生成的问题和回答
- **表达记录(Expressions)**: 孩子的表达内容
- **成果(Achievements)**: 生成的研究报告、纪录片等
- **项目活动(ProjectActivities)**: 用户操作记录

## API接口设计

### 项目管理
- `POST /api/project/create` - 创建项目
- `POST /api/project/list` - 获取项目列表
- `POST /api/project/detail` - 获取项目详情

### 观察阶段
- `POST /api/observation/image/upload` - 上传观察图片
- `POST /api/observation/image/recognize` - 识别图片内容

### 提问引导
- `POST /api/questioning/questions/generate` - 生成引导问题
- `POST /api/questioning/question/select` - 选择问题并获取回答

### 表达阶段
- `POST /api/expression/speech/text` - 语音转文字
- `POST /api/expression/note/polish` - AI润色笔记

### 成果生成
- `POST /api/achievement/report/generate` - 生成研究报告
- `POST /api/achievement/documentary/generate` - 生成纪录片
- `POST /api/achievement/poster/generate` - 生成学术海报

## AI能力集成

### OpenAI集成
- **GPT-4 Vision**: 图像识别和分析
- **GPT-4**: 问题生成、笔记润色、报告撰写
- **DALL-E**: 生成视觉辅助内容

### Claude集成
- **Claude-3**: 复杂推理和知识梳理
- **儿童友好表达**: 优化AI回答的儿童适应性

### 语音处理
- **Azure Speech Services**: 语音转文字
- **文字转语音**: 生成音频内容

## 部署和运行

### 环境要求
- Go 1.22+
- MySQL 8.0+
- Redis 6.0+
- OpenAI API Key

### 启动服务
```bash
# 启动项目管理服务
go run app/project-management/rpc/projectmanagementservice.go

# 启动API服务
go run app/api/api.go

# 启动其他RPC服务
go run app/ai-dialogue/rpc/aidialogueservice.go
```

### 数据库初始化
```bash
# 执行数据库迁移
go run database/migrations/
```

## 开发计划

### MVP阶段 (当前)
- [x] 项目架构设计
- [x] 数据库模型设计
- [x] API接口定义
- [x] 基础RPC服务实现
- [x] AI服务集成
- [ ] MVP功能实现
- [ ] 测试和验证

### 后续迭代
- [ ] 多主题支持（火箭、海洋、恐龙之外）
- [ ] 高级AI能力（跨项目关联、个性化学习路径）
- [ ] 家长协作功能
- [ ] 教育机构管理系统
- [ ] 移动App开发

## 团队和贡献

欢迎对AI教育产品感兴趣的开发者加入！

### 开发规范
- 遵循 go-zero 开发模式
- 使用 Git Flow 工作流
- 代码提交遵循 Conventional Commits
- 所有功能需要单元测试覆盖

---

*"用AI守护和激发每个孩子与生俱来的好奇心与创造力"* 🚀🦕🤖
