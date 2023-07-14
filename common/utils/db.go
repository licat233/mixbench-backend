package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func MysqlPing(sqlConn sqlx.SqlConn) error {
	db, err := sqlConn.RawDB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func NewMysql(dsn string, opts ...sqlx.SqlOption) (sqlx.SqlConn, error) {
	dbName := extractDatabaseNameFromDSN(dsn)
	if err := InitDatabase(dbName); err != nil {
		return nil, err
	}
	conn := sqlx.NewMysql(dsn, opts...)
	return conn, MysqlPing(conn)
}

func InitDatabase(databaseName string) error {
	if databaseName == "" {
		return nil
	}
	dsn := NewStringByEnv("root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE_PREFIX}?${MYSQL_CONN_OPTIONS}")

	// MySQL数据库连接信息
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL(%s): %v", dsn, err)
	}
	defer db.Close()

	// 判断数据库是否存在
	rows, err := db.Query(fmt.Sprintf("SHOW DATABASES LIKE '%s'", databaseName))
	if err != nil {
		return fmt.Errorf("failed to check if '%s' database exists: %v", databaseName, err)
	}
	defer rows.Close()

	databaseExists := rows.Next()
	// 如果数据库存在，则不用管
	if databaseExists {
		return nil
	}

	//创建该数据库
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", databaseName))
	if err != nil {
		return fmt.Errorf("failed to create '%s' database: %v", databaseName, err)
	}

	//切换到该数据库
	_, err = db.Exec(fmt.Sprintf("USE %s", databaseName))
	if err != nil {
		return fmt.Errorf("failed switch to '%s' database: %v", databaseName, err)
	}

	//获取sql文件名称，即服务名
	parts := strings.SplitN(databaseName, "_", 2)
	serviceName := parts[0]
	if len(parts) > 1 {
		serviceName = parts[1]
	}

	//开始执行sql文件
	sqlfileName := fmt.Sprintf("%s.sql", serviceName)
	filepath := path.Join(CurrentProjectRootPath(), "deploy/mysql", sqlfileName)
	// 读取SQL文件内容
	queries, err := readSQLFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read '%s' file:%v", filepath, err)
	}
	// 执行SQL语句
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to execute '%s' SQL queries: %v", sqlfileName, err)
		}
	}

	fmt.Printf("mysql database '%s' create successfully!\n", databaseName)
	return nil
}

// 从DSN字符串中提取数据库名称
func extractDatabaseNameFromDSN(dsn string) string {
	// 分割DSN字符串，以"/"作为分隔符
	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return ""
	}

	// 获取最后一个分割部分，并去除可能出现的查询参数
	dbPart := strings.Split(parts[len(parts)-1], "?")[0]

	return strings.TrimSpace(dbPart)
}

// 读取SQL文件内容并提取所有SQL语句
func readSQLFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 按分号分隔SQL语句
	queries := strings.Split(string(content), ";")
	var cleanQueries []string

	// 清理并提取每个SQL语句
	for _, query := range queries {
		cleanQuery := strings.TrimSpace(query)
		if cleanQuery != "" {
			cleanQueries = append(cleanQueries, cleanQuery)
		}
	}

	return cleanQueries, nil
}
