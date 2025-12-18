# 故障排除指南

## 端口占用错误

### 错误信息
```
listen tcp 0.0.0.0:8082: bind: address already in use
panic: listen tcp 0.0.0.0:8082: bind: address already in use
```

### 快速修复

**方法1：使用工具脚本（最简单）**
```bash
cd explorapal
chmod +x tools/kill_port.sh
./tools/kill_port.sh 9002
```

**方法2：手动终止进程（macOS）**
```bash
# 查找并终止占用9002端口的进程
lsof -ti :9002 | xargs kill -9
```

**方法3：手动终止进程（Linux）**
```bash
# 查找占用端口的进程
sudo netstat -tlnp | grep :9002
# 或
sudo ss -tlnp | grep :9002

# 终止进程（替换<PID>为实际进程ID）
sudo kill -9 <PID>
```

**方法4：修改配置文件使用其他端口**

如果无法终止占用端口的进程，可以修改配置文件：

1. **修改AI对话RPC服务端口**
   ```bash
   # 编辑 app/ai-dialogue/rpc/etc/ai-dialogue.yaml
   # 将 ListenOn: 0.0.0.0:9002 改为其他端口
   ```

2. **修改项目管理RPC服务端口**
   ```bash
   # 编辑 app/project-management/rpc/etc/project-management.yaml
   # 将 ListenOn: 0.0.0.0:9001 改为其他端口
   ```

3. **修改API服务端口**
   ```bash
   # 编辑 app/api/etc/api.yaml
   # 将 Port: 9003 改为其他端口
   ```

### 默认端口列表

| 服务 | 端口 | 配置文件 |
|------|------|----------|
| 项目管理RPC | 9001 | `app/project-management/rpc/etc/project-management.yaml` |
| AI对话RPC | 9002 | `app/ai-dialogue/rpc/etc/ai-dialogue.yaml` |
| API服务 | 9003 | `app/api/etc/api.yaml` |

## 其他常见问题

### Protobuf代码未生成

**错误信息：**
```
package explorapal/app/ai-dialogue/rpc/aidialogue is not in std
```

**解决方案：**
```bash
cd explorapal/app/ai-dialogue/rpc
chmod +x generate_proto.sh
./generate_proto.sh
```

### 数据库连接失败

**错误信息：**
```
dial tcp 127.0.0.1:3306: connect: connection refused
```

**解决方案：**
1. 检查MySQL服务是否运行
2. 检查数据库配置中的密码是否正确
3. 确认数据库已创建：`CREATE DATABASE explorapal;`

### 配置文件未找到

**错误信息：**
```
config file etc/ai-dialogue.yaml not found
```

**解决方案：**
1. 从示例文件复制配置：
   ```bash
   cp app/ai-dialogue/rpc/etc/ai-dialogue.yaml.example app/ai-dialogue/rpc/etc/ai-dialogue.yaml
   ```
2. 编辑配置文件，填入实际的配置值

## 获取帮助

如果以上方法都无法解决问题，请：
1. 检查日志文件中的详细错误信息
2. 确认所有依赖服务（MySQL、Redis）都在运行
3. 查看 `tools/README.md` 了解可用工具
