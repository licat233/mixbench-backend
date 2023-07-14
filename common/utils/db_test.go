package utils

import (
	"fmt"
	"testing"
)

func Test_db(t *testing.T) {
	dsn := "mysql://username:password@tcp(host:port)/database?param=value"
	res := extractDatabaseNameFromDSN(dsn)
	fmt.Println("結果:", res)
}
