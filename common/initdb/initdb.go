package initdb

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatasource(user, password, host, port, dbName string) string {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

func CreateDatabase(dsn, dbName string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// 创建数据库
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}
	return nil
}

func ExecSqlFile(sqlFile string, dsn string) error {
	// 打开 SQL 文件
	file, err := os.Open(sqlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// 创建 bufio.Reader 实例
	reader := bufio.NewReader(file)

	// 逐行读取 SQL 语句并执行
	for {
		line, err := reader.ReadString(';')
		if err != nil {
			break
		}
		_, err = db.Exec(line)
		if err != nil {
			return err
		}
	}
	return nil
}
