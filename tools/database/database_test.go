package database

import (
	"context"
	"database/sql"
	"testing"
	// 注册 mysql 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	UserID int64
}

func Test_sql(t *testing.T) {
	// 创建 db 实例
	db, err := sql.Open("mysql", "platform_group:rEfXWkH\"s,?q.IAF@tcp(120.92.49.234:3306)/aigc_gateway_test?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s")
	if err != nil {
		t.Error(err)
		return
	}

	// 执行 sql
	ctx := context.Background()
	row := db.QueryRowContext(ctx, "SELECT id FROM aksks")
	if row.Err() != nil {
		t.Error(err)
		return
	}
	row = db.QueryRowContext(ctx, "SELECT id FROM aksks")
	// 解析结果
	var u user
	if err = row.Scan(&u.UserID); err != nil {
		t.Error(err)
		return
	}
	t.Log(u.UserID)
}
