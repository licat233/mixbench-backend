package utils

import (
	"os"
	"strings"
)

// 读取SQL文件内容并提取所有SQL语句
func ReadSQLFile(filename string) ([]string, error) {
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

// 从DSN字符串中提取数据库名称
func ExtractDatabaseNameFromDSN(dsn string) string {
	// 分割DSN字符串，以"/"作为分隔符
	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return ""
	}

	// 获取最后一个分割部分，并去除可能出现的查询参数
	dbPart := strings.Split(parts[len(parts)-1], "?")[0]

	return strings.TrimSpace(dbPart)
}
