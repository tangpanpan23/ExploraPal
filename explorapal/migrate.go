package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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
	conf.MustLoad("app/api/etc/api.yaml", &c)

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

	// 执行所有迁移文件
	fmt.Println("🚀 开始执行数据库迁移...")

	// 定义迁移文件列表（按时间戳顺序）
	migrations := []struct {
		file string
		desc string
	}{
		{
			file: "database/migrations/20241217000001_create_explorapal_tables.up.sql",
			desc: "创建基础表结构",
		},
		{
			file: "database/migrations/20241220000001_alter_observations_image_url_length.up.sql",
			desc: "增加observations表image_url字段长度",
		},
	}

	for _, migration := range migrations {
		fmt.Printf("📄 执行迁移: %s\n", migration.desc)

		content, err := ioutil.ReadFile(migration.file)
		if err != nil {
			log.Fatalf("读取迁移文件失败 %s: %v", migration.file, err)
		}

		// 分割SQL语句并执行
		sqlStatements := strings.Split(string(content), ";")

		for i, statement := range sqlStatements {
			statement = strings.TrimSpace(statement)
			if statement == "" {
				continue
			}

			fmt.Printf("  执行第 %d 个SQL语句...\n", i+1)
			if _, err := db.Exec(statement); err != nil {
				log.Fatalf("执行SQL失败: %v\nSQL: %s", err, statement)
			}
		}

		fmt.Printf("✅ %s 迁移完成\n\n", migration.desc)
	}

	fmt.Println("✅ 数据库迁移执行完成！")
	fmt.Println("📋 已创建/更新的内容：")
	fmt.Println("  ✅ users (用户表)")
	fmt.Println("  ✅ projects (项目表)")
	fmt.Println("  ✅ observations (观察记录表) - image_url字段已扩展为LONGTEXT")
	fmt.Println("  ✅ questions (问题记录表)")
	fmt.Println("  ✅ expressions (表达记录表)")
	fmt.Println("  ✅ achievements (成果记录表)")
	fmt.Println("  ✅ project_activities (项目活动记录表)")
	fmt.Println("")
	fmt.Println("📝 数据库字段更新说明：")
	fmt.Println("  - observations.image_url: VARCHAR(500) → LONGTEXT")
	fmt.Println("    原因：支持存储base64编码的图片数据")
}
