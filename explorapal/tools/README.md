# 工具脚本说明

## 端口管理工具

### kill_port.sh

强制终止占用指定端口的进程。

**用法：**
```bash
chmod +x tools/kill_port.sh
./tools/kill_port.sh <端口号>
```

**示例：**
```bash
# 终止占用8082端口的进程
./tools/kill_port.sh 8082

# 终止占用8081端口的进程
./tools/kill_port.sh 8081
```

**说明：**
- 自动查找占用指定端口的所有进程
- 强制终止这些进程
- 验证端口是否已释放
- 如果无法终止，会提示需要sudo权限

### check_port.sh

检查端口占用情况，并提供交互式终止选项。

**用法：**
```bash
chmod +x tools/check_port.sh
./tools/check_port.sh <端口号>
```

**示例：**
```bash
# 检查8082端口占用情况
./tools/check_port.sh 8082
```

**说明：**
- 检查指定端口是否被占用
- 如果被占用，显示占用进程信息
- 询问是否终止占用进程
- 需要用户确认才会终止进程

## 数据库工具

### checkdb/main.go

检查数据库连接和表结构。

**用法：**
```bash
go run tools/checkdb/main.go
```

**说明：**
- 从 `app/api/etc/api.yaml` 读取数据库配置
- 测试数据库连接
- 检查所有必需的表是否存在
- 显示表结构示例
