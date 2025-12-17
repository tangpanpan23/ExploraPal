#!/bin/bash

# 数据库迁移脚本
echo "🚀 开始执行数据库迁移..."

# 从api.yaml配置文件读取数据库连接信息
# DBConfig.DataSource: root:tangpanpan314@tcp(localhost:3306)/explorapal?charset=utf8mb4&parseTime=true&loc=Local

DB_HOST="localhost"
DB_PORT="3306"
DB_USER="root"
DB_PASS="tangpanpan314"
DB_NAME="explorapal"

# 迁移文件路径
MIGRATION_FILE="database/migrations/20241217000001_create_explorapal_tables.up.sql"

# 检查迁移文件是否存在
if [ ! -f "$MIGRATION_FILE" ]; then
    echo "❌ 迁移文件不存在: $MIGRATION_FILE"
    exit 1
fi

echo "📁 迁移文件: $MIGRATION_FILE"
echo "🗄️  目标数据库: $DB_NAME"

# 执行迁移
mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$MIGRATION_FILE"

if [ $? -eq 0 ]; then
    echo "✅ 数据库迁移执行成功！"
    echo "📋 已创建的表："
    echo "  - users (用户表)"
    echo "  - projects (项目表)"
    echo "  - observations (观察记录表)"
    echo "  - questions (问题记录表)"
    echo "  - expressions (表达记录表)"
    echo "  - achievements (成果记录表)"
    echo "  - project_activities (项目活动记录表)"
else
    echo "❌ 数据库迁移执行失败！"
    exit 1
fi
