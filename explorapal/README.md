# 探索伙伴 (ExploraPal) - AI学习产品

## 项目简介

探索伙伴是一个基于AI的儿童学习产品，采用"从教知识点转向辅助探索"的理念，帮助孩子通过多模态AI技术，将兴趣爱好转化为深度研究项目。

## 核心理念

- **从填充式学习到探索式学习**：不直接灌输知识，而是通过AI引导孩子自主探索
- **兴趣驱动**：从孩子的真实兴趣出发（如恐龙、火箭、Minecraft）
- **多模态交互**：支持图像识别、语音处理、AI对话
- **项目式学习**：将探索过程组织为完整的"研究项目"

## 🎨 前端演示页面

提供完整的可视化演示界面，可以通过API路由访问：
```
http://localhost:9003/api/common/demo
```

演示页面包含完整的AI学习流程：
- 📸 拖拽上传图片进行AI分析
- 🎤 上传音频文件进行语音转文字
- 🎥 上传视频进行内容分析和AI创作
- 📊 实时显示处理进度和结果
- 🎵 页面内播放生成的音频和视频
- 📥 一键下载生成的多媒体内容

## 核心功能

### 🎯 多模态AI学习平台 ✅ 已完全实现

#### 完整学习流程：观察 → 提问 → 表达 → 创作

1. **观察阶段** 🔍
   - ✅ 孩子拍照上传恐龙图片
   - ✅ AI图像识别分析恐龙特征（qwen3-vl-plus）
   - ✅ 自动创建观察记录，支持无observation_id调用

2. **提问引导** ❓
   - ✅ 基于观察结果生成个性化问题（qwen-flash）
   - ✅ 问题类型：观察、推理、实验、比较
   - ✅ AI提供答案和延伸思考

3. **表达阶段** 🎤
   - ✅ 语音转文字记录想法（qwen3-omni-flash）
   - ✅ AI润色生成研究笔记（qwen-flash）
   - ✅ 文字转语音播报内容（qwen3-omni-flash）

4. **创作阶段** 🎬
   - ✅ 生成研究报告（qwen3-max）
   - ✅ 视频内容分析（qwen3-omni-flash）
   - ✅ AI视频创作（doubao-seedance-1.0-lite-i2v）

## 技术架构

### 后端技术栈
- **框架**: Go + go-zero 微服务框架
- **数据库**: MySQL 8.0+
- **缓存**: Redis 6.0+
- **消息队列**: Kafka (异步任务处理)
- **AI服务**: 双平台AI能力集成
  - **阿里云Qwen系列**: TAL MLOps平台，专注同步多模态处理
  - **豆包Doubao系列**: 星图AI平台，专注异步视频生成
- **多媒体**: 系统TTS/STT + AI视频处理
- **架构**: 5个微服务 (API + 4个RPC服务)

### AI服务架构

ExploraPal集成了两大AI服务生态，为儿童学习提供全面的AI能力支持：

#### 🌟 **阿里云Qwen大模型系列** (TAL MLOps平台)
**服务类型**: 同步API调用
**网络架构**: 内部专线直连
**响应特点**: 实时处理，适用于对话和分析任务

**核心模型配置**:
- **qwen3-vl-plus**: 视觉理解旗舰模型
  - 擅长图像内容分析和场景理解
  - 支持复杂视觉问答和详细描述生成
  - 用于儿童拍照图片的深度分析

- **qwen-flash**: 快速推理模型
  - 平衡速度与质量的通用模型
  - 适用于文本生成和问题回答
  - 用于个性化问题生成和笔记润色

- **qwen3-max**: 深度推理模型
  - 复杂任务的专业解决方案
  - 支持长文本理解和结构化输出
  - 用于研究报告生成和深度分析

- **qwen3-omni-flash**: 多模态实时交互模型
  - 语音+文本+视觉一体化处理
  - 支持中英文语音识别和合成
  - 用于语音转文字和文字转语音

#### 🎬 **豆包Doubao大模型系列** (星图AI平台)
**服务类型**: 异步API调用
**网络架构**: 公网API调用
**响应特点**: 后台处理，适用于计算密集型任务

**核心模型配置**:
- **doubao-seedance-1.0-lite-i2v**: 图像到视频生成模型
  - 将静态图片转换为动态视频内容
  - 支持教育场景的视频创作
  - 最大视频时长：10秒
  - 用于将儿童观察的图片生成教学视频

#### 🔄 **AI服务调用架构**

```
┌─────────────────────────────────────────────────────────────┐
│                    ExploraPal AI服务层                      │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌──────────────────────────────────┐  │
│  │ TAL MLOps平台   │  │        星图AI平台                │  │
│  │ (阿里云Qwen)    │  │      (豆包Doubao)               │  │
│  ├─────────────────┤  ├──────────────────────────────────┤  │
│  │ • 同步调用      │  │ • 异步调用                      │  │
│  │ • 实时响应      │  │ • 后台处理                      │  │
│  │ • 内部网络      │  │ • 公网API                       │  │
│  └─────────────────┘  └──────────────────────────────────┘  │
│           │                        │                        │
│           ▼                        ▼                        │
│  ┌─────────────────┐      ┌─────────────────┐              │
│  │   视觉分析      │      │   视频生成      │              │
│  │   文本生成      │      │   (异步)        │              │
│  │   语音处理      │      └─────────────────┘              │
│  │   深度推理      │                                         │
│  └─────────────────┘                                         │
└─────────────────────────────────────────────────────────────┘
```

#### ⚡ **性能优化策略**
- **智能路由**: 根据任务类型自动选择最适合的AI模型
- **并发处理**: 多模型并行调用提升响应速度
- **缓存机制**: 热点数据缓存减少重复调用
- **错误重试**: 自动重试机制确保服务稳定性
- **负载均衡**: 多实例部署支持高并发访问

### 项目结构
```
explorapal/
├── app/
│   ├── api/                    # REST API服务 (Port 9003)
│   ├── project-management/rpc/ # 项目管理RPC服务 (Port 9001)
│   ├── ai-dialogue/rpc/        # AI对话RPC服务 (Port 9002)
│   ├── audio-processing/rpc/   # 语音处理RPC服务 (Port 9004)
│   └── video-processing/rpc/   # 视频处理RPC服务 (Port 9005)
├── common/                     # 通用工具
├── constant/                   # 常量定义
├── database/migrations/        # 数据库迁移
├── pkg/                        # 通用包
├── third/                      # 第三方服务集成
│   ├── openai/                 # 阿里云Qwen客户端
│   └── azure-speech/           # 阿里云语音服务（兼容）
└── storage/                    # 文件存储
```

## 核心数据模型

- **用户(Users)**: 用户基本信息
- **项目(Projects)**: 探索项目
- **观察记录(Observations)**: 图像分析结果
- **问题记录(Questions)**: AI生成的问题和回答
- **表达记录(Expressions)**: 孩子的表达内容
- **成果(Achievements)**: 生成的研究报告、纪录片、AI视频等
- **项目活动(ProjectActivities)**: 用户操作记录

### 多媒体功能

#### 🎵 语音处理
- **语音转文字**: 支持WAV、MP3、OGG等多种音频格式
- **文字转语音**: 支持中英文，多种语音类型和语速调节
- **AI增强**: 结合Qwen多模态模型提升识别准确性

#### 🎬 视频处理
- **视频分析**: 多维度分析（场景、物体、情感、文字、音频）
- **视频生成**: 基于脚本自动生成教育、故事、动画视频
- **智能创作**: 支持自定义风格和时长

## API接口设计

### 项目管理
- ✅ `POST /api/project/create` - 创建项目
- ✅ `POST /api/project/list` - 获取项目列表
- ✅ `POST /api/project/detail` - 获取项目详情

### 观察阶段
- ✅ `POST /api/observation/image/recognize` - AI图像识别分析

### 提问引导
- ✅ `POST /api/questioning/questions/generate` - 生成引导问题

### 表达阶段
- ✅ `POST /api/expression/speech/text` - 语音转文字
- ✅ `POST /api/expression/note/polish` - AI润色笔记

### 语音处理 🎵
- ✅ `POST /api/audio/text-to-speech` - 文字转语音播报

### 视频处理 🎬
- ✅ `POST /api/achievement/video/analyze` - 深度视频内容分析
- ✅ `POST /api/achievement/video/generate` - AI视频创作

### 成果生成 📊
- ✅ `POST /api/achievement/report/generate` - 生成研究报告

## AI能力集成

### 阿里云Qwen集成
- **qwen3-vl-plus** (256K): 视觉理解，支持思考模式，图像和视频分析最优
- **qwen-flash** (1048.576K): 思考+非思考模式融合，问题生成和笔记润色
- **qwen3-max** (256K): 智能体编程优化，复杂推理和报告生成
- **qwen3-omni-flash** (48K): 多模态大模型，支持语音转文字、文字转语音和音频分析
- **qwen-vl-plus** (系列): 视频理解和生成，支持多媒体内容创作

### AI服务可靠性

系统集成了完整的AI服务可靠性保障：

#### 🤖 TAL MLOps平台集成
- ✅ **多模型支持**: qwen3-vl-plus、qwen-flash、qwen3-max、qwen3-omni-flash、qwen-vl-plus
- ✅ **超时管理**: 70-120秒超时设置，确保AI处理充分时间
- ✅ **服务降级**: AI服务不可用时提供合理的默认响应
- ✅ **性能监控**: 慢调用阈值设置，监控服务响应时间

#### 🔧 容错机制
- ✅ **优雅降级**: AI服务异常时返回预设的合理响应
- ✅ **数据完整性**: 确保数据库操作的原子性
- ✅ **日志记录**: 完整的操作日志便于问题排查
- ✅ **健康检查**: `/api/common/ping` 接口实时监控服务状态


## 部署和运行

### 环境要求
- Go 1.22+
- MySQL 8.0+
- Redis 6.0+
- TAL MLOps平台API访问权限 (阿里云Qwen系列模型)
- 星图AI平台API访问权限 (豆包Doubao模型)

### AI服务配置

#### 🔧 **阿里云Qwen系列配置** (TAL MLOps平台)

**接入方式**: 企业内部专线直连
**API类型**: RESTful同步调用
**网络要求**: 内网可达 (http://ai-service.tal.com)

**认证配置**:
- 申请TAL MLOps平台企业账号
- 获取应用ID和应用密钥
- API鉴权: `Authorization: Bearer {appId}:{appKey}`

**模型清单**:
| 模型名称 | 功能特点 | 适用场景 | 响应时间 |
|----------|----------|----------|----------|
| qwen3-vl-plus | 视觉理解旗舰 | 图片分析、场景识别 | 5-15秒 |
| qwen-flash | 快速文本生成 | 问题生成、笔记润色 | 2-8秒 |
| qwen3-max | 深度推理分析 | 报告生成、复杂问答 | 10-30秒 |
| qwen3-omni-flash | 多模态实时交互 | 语音转文字、文字转语音 | 3-12秒 |

#### 星图AI平台配置 (异步API)
1. 服务地址：
   - 生产环境：http://apx-api.tal.com
   - 测试环境：http://apx-api-gray.tal.com
2. 异步API鉴权：使用api-key header (`api-key: appId:appKey`)
3. 支持模型：doubao-seedance-1.0-lite-t2v (视频生成)
4. 用途：基于图片生成视频

#### ⚙️ **性能和超时配置**

**同步API配置** (阿里云Qwen):
- 连接超时: 30秒
- 读取超时: 70-120秒 (根据模型复杂度)
- 重试次数: 2次
- 并发限制: 按应用配额

**异步API配置** (豆包Doubao):
- 任务提交超时: 30秒
- 任务处理超时: 300秒
- 状态轮询间隔: 10秒
- 最大轮询次数: 60次

**服务质量监控**:
- 慢调用阈值: 10秒 (AI对话), 30秒 (多媒体处理)
- 成功率监控: >99.5%
- 平均响应时间监控
- 错误率告警机制

#### 🔐 **安全配置**

**数据传输安全**:
- HTTPS加密传输
- API密钥轮换机制
- 请求签名验证
- IP白名单控制

**内容安全**:
- 输入内容过滤
- 输出内容审核
- 用户数据隔离
- 审计日志记录

### Git提交注意事项
- ⚠️ **不要提交编译后的可执行文件**：`api`、`rpc`、`*-processing` 等可执行文件已被 `.gitignore` 排除
- 🔄 **提交前清理**：使用 `git status` 检查是否有不应提交的大文件
- 📦 **仅提交源码**：项目应该只包含源代码、配置文件和文档

### 启动服务

#### 快速启动（推荐）

```bash
cd explorapal
chmod +x start_demo.sh stop_demo.sh
./start_demo.sh
```

这将自动：
- 检查数据库连接
- 清理端口占用
- 按正确顺序启动所有服务
- 显示服务状态和测试命令

**停止演示：**
```bash
./stop_demo.sh
```

#### 🎬 视频演示生成

项目提供了专门的脚本用于生成AI视频演示，支持基于用户图片和描述的异步视频生成。

**配置说明：**
1. **复制配置文件**：
   ```bash
   cp video_generation_config.yaml.example video_generation_config.yaml
   ```

2. **编辑配置文件**：
   ```yaml
   AstraAI:
     AppID: "your_real_app_id"      # 从星图AI平台获取
     AppKey: "your_real_app_key"    # 从星图AI平台获取
   ```

3. **测试配置**：
   ```bash
   ./test_config.sh  # 验证配置是否正确
   ```

4. **脚本会自动从配置文件读取API配置，无需在代码中硬编码**

**⚠️ 重要配置要求：**
- **异步应用配置**：需要在星图AI平台申请应用后，联系平台管理员将应用配置为异步应用
- **API密钥格式**：`appId:appKey`（已在配置文件中配置）
- **模型名称**：`doubao-seedance-1.0-lite-i2v`（用于异步视频生成）

**常见错误及解决方案：**
- **错误码 110000 "Header 模型不存在"**：
  - 确认应用已配置为异步应用（需要平台管理员操作）
  - 检查API密钥是否正确
  - 联系星图AI平台技术支持

```bash
# 生成演示视频（异步）
./生成演示视频.sh <图片文件路径> <描述文本>

# 示例
./生成演示视频.sh test_image.png "小葫芦观察到了一只可爱的小恐龙，有着绿色的皮肤和长长的脖子"

# 查询生成结果
./查询视频生成结果.sh <任务ID>
```

**视频生成流程：**
1. **提交任务**：脚本直接调用星图AI平台异步API提交视频生成任务
2. **获取任务ID**：API返回唯一任务ID，保存到本地文件
3. **等待处理**：AI需要5-10分钟处理视频生成
4. **查询结果**：使用任务ID查询生成状态和下载链接
5. **下载视频**：生成完成后自动下载MP4视频文件

**修复内容：**
- ✅ **API端点修正**：从调用本地后端改为直接调用星图AI平台
- ✅ **请求格式修正**：使用正确的星图API参数格式
- ✅ **错误处理改进**：添加图片压缩和模拟测试模式
- ✅ **兼容性修复**：支持macOS和Linux的日期命令

**异步API特性：**
- **模型**：豆包Doubao-Seedance-1.0-lite-t2v
- **输入**：用户原始图片(data URL) + AI润色描述
- **输出**：高质量MP4教育视频
- **时长**：约10秒（豆包模型最大支持时长）
- **风格**：教育演示风格
- **测试模式**：使用占位符密钥时自动启用模拟响应

#### 手动启动

##### 1. 生成Protobuf代码（首次运行需要）
```bash
cd explorapal

# 生成AI对话服务的protobuf代码
cd app/ai-dialogue/rpc
protoc --go_out=. --go-grpc_out=. ai-dialogue.proto
# 或者使用goctl
goctl rpc protoc ai-dialogue.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

#### 2. 启动RPC服务

**⚠️ 如果遇到端口占用错误，请先处理：**

```bash
# 方法1：使用工具脚本终止占用端口的进程
chmod +x tools/kill_port.sh
./tools/kill_port.sh 9001  # 项目管理RPC服务端口
./tools/kill_port.sh 9002  # AI对话RPC服务端口
./tools/kill_port.sh 9003  # API服务端口

# 方法2：手动查找并终止进程
lsof -ti :8082 | xargs kill -9  # macOS
# 或
sudo netstat -tlnp | grep :8082  # Linux，然后kill对应PID
```

```bash
# 启动项目管理RPC服务
cd explorapal/app/project-management/rpc
go run projectmanagementservice.go

# 启动AI对话RPC服务（需要先生成protobuf代码）
cd explorapal/app/ai-dialogue/rpc
go run aidialogueservice.go

# 启动语音处理RPC服务
cd explorapal/app/audio-processing/rpc
go run audioprocessingservice.go

# 启动视频处理RPC服务
cd explorapal/app/video-processing/rpc
go run videoprocessingservice.go
```

#### 3. 启动API服务
```bash
cd explorapal/app/api
go run api.go
```

#### 4. 从项目根目录启动（推荐）

**启动前检查端口占用：**
```bash
cd explorapal
chmod +x tools/check_ports.sh

# 检查所有服务端口状态
./tools/check_ports.sh

# 如有需要，清理端口占用
./tools/kill_port.sh 9001  # 项目管理RPC服务
./tools/kill_port.sh 9002  # AI对话RPC服务
./tools/kill_port.sh 9003  # API服务
./tools/kill_port.sh 9004  # 语音处理RPC服务
./tools/kill_port.sh 9005  # 视频处理RPC服务
```

```bash
# 启动项目管理RPC服务
go run app/project-management/rpc/projectmanagementservice.go

# 启动AI对话RPC服务
go run app/ai-dialogue/rpc/aidialogueservice.go

# 启动语音处理RPC服务
go run app/audio-processing/rpc/main.go

# 启动视频处理RPC服务
go run app/video-processing/rpc/main.go

# 启动API服务
go run app/api/api.go
```

**默认端口配置：**
- 项目管理RPC服务: `9001`
- AI对话RPC服务: `9002`
- API服务: `9003`

如需修改端口，请编辑对应的配置文件：
- `app/project-management/rpc/etc/project-management.yaml` (默认: 9001)
- `app/ai-dialogue/rpc/etc/ai-dialogue.yaml` (默认: 9002)
- `app/api/etc/api.yaml` (默认: 9003)

### 数据库初始化
```bash
# 方法1：使用Go脚本执行迁移（推荐）
go run migrate.go

# 方法2：使用Shell脚本执行迁移
chmod +x migrate.sh
./migrate.sh

# 方法3：直接使用MySQL命令执行迁移
mysql -hlocalhost -P3306 -uroot -p<your-db-password> explorapal < database/migrations/20241217000001_create_explorapal_tables.up.sql
```

数据库配置位于 `app/api/etc/api.yaml` 中的 `DBConfig.DataSource` 字段。

### 数据库连接验证
```bash
# 检查数据库连接和表结构
go run tools/checkdb/main.go
```

## 故障排除

### 端口占用错误

**错误信息：**
```
listen tcp 0.0.0.0:8082: bind: address already in use
```

**解决方案：**

1. **使用工具脚本（推荐）**
```bash
cd explorapal
chmod +x tools/kill_port.sh
./tools/kill_port.sh 8082  # 替换为实际占用的端口号
```

2. **手动查找并终止进程（macOS）**
```bash
# 查找占用端口的进程
lsof -i :8082

# 终止进程（替换PID为实际进程ID）
kill -9 <PID>

# 或一行命令
lsof -ti :8082 | xargs kill -9
```

3. **手动查找并终止进程（Linux）**
```bash
# 查找占用端口的进程
sudo netstat -tlnp | grep :8082
# 或
sudo ss -tlnp | grep :8082

# 终止进程（替换PID为实际进程ID）
sudo kill -9 <PID>
```

4. **修改配置文件使用其他端口**
如果无法终止占用端口的进程，可以修改配置文件使用其他端口：
- 编辑 `app/ai-dialogue/rpc/etc/ai-dialogue.yaml`，修改 `ListenOn: 0.0.0.0:8082` 为其他端口（如 `8083`）
- 编辑 `app/project-management/rpc/etc/project-management.yaml`，修改 `ListenOn: 0.0.0.0:8081` 为其他端口（如 `8084`）
- 编辑 `app/api/etc/api.yaml`，修改 `Port: 9999` 为其他端口（如 `9998`）

## 完整演示

完整的多模态AI学习演示请查看 [MVP_DEMO.md](./MVP_DEMO.md)

演示包含：
- 完整的API调用流程
- 实际的请求/响应示例
- 儿童学习探索的完整路径

## 故障排除

遇到问题？请查看 [TROUBLESHOOTING.md](./TROUBLESHOOTING.md) 获取详细的故障排除指南。

常见问题：
- **端口占用错误**：使用 `tools/kill_port.sh` 脚本终止占用端口的进程
- **Protobuf代码未生成**：运行 `app/ai-dialogue/rpc/generate_proto.sh`
- **数据库连接失败**：检查MySQL服务是否运行，确认配置文件中的密码正确

## 开发计划

### 第一阶段 ✅ 已完成
- [x] 项目架构设计
- [x] 数据库模型设计
- [x] API接口定义
- [x] 基础RPC服务实现
- [x] AI服务集成
- [x] **多模态AI功能完全实现** 🎉
- [x] 核心接口联调验证

#### 核心功能清单 ✅
- [x] **项目管理**: 创建、列表、详情查询
- [x] **AI图像识别**: 恐龙特征分析，支持自动观察记录创建
- [x] **智能问题生成**: 基于上下文生成个性化引导问题
- [x] **语音交互**: 语音转文字、文字转语音
- [x] **内容创作**: AI笔记润色、研究报告生成
- [x] **视频处理**: 视频分析、AI视频生成
- [x] **多模态演示**: 完整学习流程演示文档

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

## 项目清理

如果需要清理测试和调试文件，请运行：

```bash
# 查看将要删除的文件
cat cleanup.sh

# 执行清理（请谨慎操作）
chmod +x cleanup.sh
./cleanup.sh
```

**清理内容包括**:
- `api_examples.md` - 旧版API示例文档
- `demo_flow.sh` - 演示流程脚本
- `test_demo.sh` - 测试脚本
- `TROUBLESHOOTING.md` - 故障排除指南
- `tools/` - 调试工具目录

---

*"用AI守护和激发每个孩子与生俱来的好奇心与创造力"*

**🎯 ExploraPal v1.0.0 - 多模态AI学习平台** ✨
