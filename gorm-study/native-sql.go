package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// datasourcename
	// 用户名:密码@tcp(地址（localhost或者127.0.0.1:3306）)/数据库库名
	dsn := "root:@你的密码@tcp(127.0.0.1:3306)/gorm_db_new"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("dsn格式错误%s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败%s", err)
	}
	fmt.Println(db)
}
