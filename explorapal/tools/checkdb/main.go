package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	DBConfig struct {
		DataSource string `json:",optional"`
	} `json:",optional"`
}

func main() {
	// 读取配置文件
	var c Config
	conf.MustLoad("../../app/api/etc/api.yaml", &c)

	if c.DBConfig.DataSource == "" {
		log.Fatal("数据库配置未找到")
	}

	fmt.Printf("🔧 数据库配置: %s\n", c.DBConfig.DataSource)

	// 连接数据库
	db, err := sql.Open("mysql", c.DBConfig.DataSource)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("✅ 数据库连接成功")

	// 检查表是否存在
	tables := []string{
		"users", "projects", "observations",
		"questions", "expressions", "achievements", "project_activities",
	}

	fmt.Println("📋 检查数据库表...")

	for _, table := range tables {
		var count int
		query := fmt.Sprintf("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'explorapal' AND table_name = '%s'", table)
		err := db.QueryRow(query).Scan(&count)
		if err != nil {
			log.Printf("检查表 %s 失败: %v", table, err)
			continue
		}

		if count > 0 {
			fmt.Printf("  ✅ %s 表存在\n", table)
		} else {
			fmt.Printf("  ❌ %s 表不存在\n", table)
		}
	}

	// 显示表结构示例
	fmt.Println("\n📊 用户表结构示例:")
	rows, err := db.Query("DESCRIBE users LIMIT 5")
	if err != nil {
		log.Printf("获取表结构失败: %v", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var field, typ, null, key, def, extra string
			rows.Scan(&field, &typ, &null, &key, &def, &extra)
			fmt.Printf("  %s %s\n", field, typ)
		}
	}
}
