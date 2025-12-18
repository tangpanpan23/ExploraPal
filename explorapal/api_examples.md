# API 接口使用示例

## 项目管理 API

### 1. 创建项目 - `POST /api/project/create`

**请求示例：**

```bash
curl -X POST "http://localhost:9003/api/project/create" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "title": "恐龙探索项目",
    "description": "探索恐龙世界的奥秘，从化石到生活习性",
    "category": "dinosaur",
    "tags": ["恐龙", "化石", "古生物", "探索"]
  }'
```

**请求参数说明：**

| 参数 | 类型 | 必需 | 说明 | 示例 |
|------|------|------|------|------|
| `user_id` | int64 | 是 | 用户ID | `1` |
| `title` | string | 是 | 项目标题 | `"恐龙探索项目"` |
| `description` | string | 否 | 项目描述 | `"探索恐龙世界的奥秘，从化石到生活习性"` |
| `category` | string | 是 | 项目类别 | `"dinosaur"` |
| `tags` | []string | 否 | 标签数组 | `["恐龙", "化石", "古生物", "探索"]` |

**成功响应示例：**
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

**错误响应示例：**
```json
{
  "code": 500,
  "message": "项目创建失败: RPC服务不可用"
}
```

### 2. 获取项目列表 - `POST /api/project/list`

```bash
curl -X POST "http://localhost:9003/api/project/list" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "category": "dinosaur",
    "status": "active",
    "page_size": 10,
    "page": 1
  }'
```

### 3. 获取项目详情 - `POST /api/project/detail`

```bash
curl -X POST "http://localhost:9003/api/project/detail" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1
  }'
```

### 4. 更新项目状态 - `POST /api/project/status/update`

```bash
curl -X POST "http://localhost:9003/api/project/status/update" \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "user_id": 1,
    "status": "completed"
  }'
```

## 完整请求示例

### 创建恐龙探索项目
```json
{
  "user_id": 1,
  "title": "我的恐龙王国探索",
  "description": "跟随小恐龙的脚步，一起探索古老的地球世界，了解恐龙的生活习性、栖息环境和进化历程",
  "category": "dinosaur",
  "tags": ["恐龙", "古生物", "进化", "探索", "科学"]
}
```

### 创建火箭发射项目
```json
{
  "user_id": 2,
  "title": "火箭发射挑战",
  "description": "学习物理学和工程学原理，设计并发射自己的火箭",
  "category": "rocket",
  "tags": ["火箭", "物理", "工程", "发射", "太空"]
}
```

### 创建我的世界建筑项目
```json
{
  "user_id": 3,
  "title": "我的世界奇幻城堡",
  "description": "在Minecraft中设计和建造一座宏伟的城堡，学习建筑学和创造力",
  "category": "minecraft",
  "tags": ["Minecraft", "建筑", "创造力", "游戏", "设计"]
}
```

## 支持的项目类别

- `dinosaur` - 恐龙探索
- `rocket` - 火箭发射
- `minecraft` - 我的世界
- `ocean` - 海洋探索
- `space` - 太空探索
- `robot` - 机器人制作
- `chemistry` - 化学实验
- `biology` - 生物观察

## 状态说明

- `active` - 进行中
- `completed` - 已完成
- `paused` - 已暂停
- `draft` - 草稿状态

## 注意事项

1. 所有API都运行在端口 `9003`
2. 当前JWT认证已暂时禁用，所有请求都会被处理
3. 项目管理RPC服务需要运行在端口 `9001`
4. AI对话RPC服务运行在端口 `9002`
