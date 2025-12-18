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
- **AI服务**: 阿里云Qwen系列大模型 (qwen3-vl-plus, qwen-flash, qwen3-max)
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

### 阿里云Qwen集成
- **qwen3-vl-plus** (256K): 视觉理解，支持思考模式，图像分析最优
- **qwen-flash** (1048.576K): 思考+非思考模式融合，问题生成和笔记润色
- **qwen3-max** (256K): 智能体编程优化，复杂推理和报告生成
- **qwen3-omni-flash** (48K): 多模态大模型，支持语音转文字和文字转语音

### 开发环境降级机制
当AI服务不可用时，系统会自动降级到模拟响应模式：
- ✅ 图像分析：返回合理的默认分析结果
- ✅ 问题生成：提供预设的教育性问题
- ✅ 笔记润色：保持原始内容并添加基本结构
- ✅ 报告生成：生成模板化的研究报告

这样确保即使在没有外部AI服务的情况下，MVP演示功能也能正常运行。


## 部署和运行

### 环境要求
- Go 1.22+
- MySQL 8.0+
- Redis 6.0+
- 阿里云DashScope API Key

### 内部AI服务配置 (TAL MLOps平台)
1. 获取TAL MLOps应用ID和应用密钥
2. 确保内部AI服务网络可达 (http://ai-service.tal.com)
3. 确认账户权限，支持调用qwen3-vl-plus、qwen-flash、qwen3-max、qwen3-omni-flash等模型
4. 注意：所有服务超时时间设置为60秒（包括HTTP请求、RPC调用、AI服务），确保网络连接稳定
5. 慢调用阈值：AI对话服务设置为5秒，项目管理服务设置为2秒，避免正常业务调用被误报为慢调用

### 启动服务

#### 快速启动（推荐）

**一键启动完整演示环境：**
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
```

```bash
# 启动项目管理RPC服务
go run app/project-management/rpc/projectmanagementservice.go

# 启动AI对话RPC服务
go run app/ai-dialogue/rpc/aidialogueservice.go

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

## MVP演示

完整的"发现口袋 - 恐龙篇"演示请查看 [MVP_DEMO.md](./MVP_DEMO.md)

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
