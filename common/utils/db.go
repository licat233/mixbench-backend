package utils

import "github.com/zeromicro/go-zero/core/stores/sqlx"

func MysqlPing(sqlConn sqlx.SqlConn) error {
	db, err := sqlConn.RawDB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func NewMysql(datasource string, opts ...sqlx.SqlOption) (sqlx.SqlConn, error) {
	conn := sqlx.NewMysql(datasource, opts...)
	return conn, MysqlPing(conn)
}
