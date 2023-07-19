package mysqlx

import (
	"bufio"
	"database/sql"
	"fmt"
	"mixbench/common/utils"
	"os"
	"path"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Mysqlx struct {
	DB *sql.DB
}

func NewMysql(dsn string, opts ...sqlx.SqlOption) (sqlx.SqlConn, error) {
	dbName := utils.ExtractDatabaseNameFromDSN(dsn)
	mysqlx, err := New()
	if err != nil {
		return nil, err
	}
	if err := mysqlx.InitDatabase(dbName); err != nil {
		return nil, err
	}
	defer mysqlx.Close()
	conn := sqlx.NewMysql(dsn, opts...)
	db, err := conn.RawDB()
	if err != nil {
		return conn, err
	}
	return conn, db.Ping()
}

func New() (*Mysqlx, error) {
	//连接到默认数据库
	dsn := utils.NewStringByEnv("root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/?${MYSQL_CONN_OPTIONS}")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Mysqlx{
		DB: db,
	}, nil
}

func CreateDatasource(user, password, host, port, dbName string) string {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

func (m *Mysqlx) Close() {
	m.DB.Close()
}

// 创建数据库
func (m *Mysqlx) CreateDatabase(dbName string) error {
	_, err := m.DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}
	return nil
}

// 执行sql文件
func (m *Mysqlx) ExecSqlFile(sqlFile string, dsn string) error {
	fmt.Println("開始執行sql文件")
	// 打开 SQL 文件
	file, err := os.Open(sqlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建 bufio.Reader 实例
	reader := bufio.NewReader(file)

	// 逐行读取 SQL 语句并执行
	for {
		line, err := reader.ReadString(';')
		if err != nil {
			break
		}
		_, err = m.DB.Exec(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Mysqlx) InitDatabase(databaseName string) error {
	lockName := "initMysqlDB.lock"
	lockfile := path.Join(utils.ProjectRoot, lockName)
	//判断当前目录下是否存在 mysql.lock文件
	if utils.FileExist(lockfile) {
		//如果存在，则读取内容，看是否存在 databaseName
		content, err := utils.ReadFile(lockfile)
		if err != nil {
			return fmt.Errorf("读取%s文件失败: %s", lockName, err)
		}
		fileContent := strings.TrimSpace(string(content))
		// 判断是否存在 databaseName
		if strings.Contains(fileContent, databaseName) {
			//已存在数据库 databaseName, 忽略，不操作
			return nil
		}
	} else {
		//如果不存在文件
		//创建
		if err := utils.CreateFile(lockfile); err != nil {
			return fmt.Errorf("创建%s文件失败: %s", lockName, err)
		}
	}

	if err := m.ExecServiceSqlFile(databaseName); err != nil {
		return err
	}

	//追加写入 dadabaseName
	if err := utils.InsertFileContentAfter(lockfile, databaseName); err != nil {
		return fmt.Errorf("写入%s文件内容失败: %s", lockName, err)
	}

	return nil
}

func (m *Mysqlx) ExecServiceSqlFile(databaseName string) error {
	// 判断数据库是否存在
	// rows, err := m.DB.Query(fmt.Sprintf("SHOW DATABASES LIKE '%s'", databaseName))
	// if err != nil {
	// 	return fmt.Errorf("failed to check if '%s' database exists: %v", databaseName, err)
	// }
	// defer rows.Close()

	// databaseExists := rows.Next()
	// // 如果数据库存在，则不用管
	// if databaseExists {
	// 	return nil
	// }

	//创建该数据库
	_, err := m.DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", databaseName))
	if err != nil {
		return fmt.Errorf("failed to create '%s' database: %v", databaseName, err)
	}

	//切换到该数据库
	_, err = m.DB.Exec(fmt.Sprintf("USE %s", databaseName))
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
	filepath := path.Join(utils.ProjectRoot, "deploy/mysql", sqlfileName)
	// 打开 SQL 文件
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	//文件存在且可读取，则开始清空数据库
	if err := m.DropTables(""); err != nil {
		return err
	}

	// 创建 bufio.Reader 实例
	reader := bufio.NewReader(file)

	// 逐行读取 SQL 语句并执行
	for {
		line, err := reader.ReadString(';')
		if err != nil {
			break
		}
		_, err = m.DB.Exec(line)
		if err != nil {
			return err
		}
	}

	fmt.Printf("mysql database '%s' create successfully!\n", databaseName)
	return nil
}

func (m *Mysqlx) DropTables(databaseName string) error {
	db := m.DB
	if databaseName != "" {
		// 切换到该数据库
		_, err := m.DB.Exec(fmt.Sprintf("USE %s", databaseName))
		if err != nil {
			return fmt.Errorf("failed switch to '%s' database: %v", databaseName, err)
		}
	}

	// 获取所有表
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	defer rows.Close()

	// 存储表名的切片
	var tables []string

	// 遍历查询结果，将表名存储到切片中
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return err
		}
		tables = append(tables, table)
	}

	// 清空表内的数据
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE `%s`", table))
		if err != nil {
			return err
		}
	}

	// 删除存储过程、触发器、视图等对象
	objects := []string{"procedure_name", "trigger_name", "view_name"}

	for _, object := range objects {
		_, err := db.Exec(fmt.Sprintf("DROP PROCEDURE IF EXISTS %s", object))
		if err != nil {
			return err
		}
	}

	fmt.Println(databaseName + "数据库已清空")
	return nil
}
