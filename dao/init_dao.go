package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Db ① GoプログラムからMySQLへ接続
var Db *sql.DB

func init() {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER") //"test_user"
	mysqlPwd := os.Getenv("MYSQL_PWD")   // "password"
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE") //"test_database"

	mysqlUser = "test_user"
	mysqlPwd = "password"
	mysqlDatabase = "test_database"

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	// _db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost,mysqlDatabase))
	// _db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s?parseTime=true", mysqlUser, mysqlPwd, mysqlDatabase))
	// ①-2
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	Db = _db
}
