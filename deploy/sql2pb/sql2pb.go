/*
 * @Author: licat
 * @Date: 2023-01-15 14:07:25
 * @LastEditors: licat
 * @LastEditTime: 2023-01-30 14:10:16
 * @Description: licat233@gmail.com
 */
package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"mixbench/deploy/sql2pb/core"
	"os"
	"path"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbType := flag.String("db", "mysql", "the database type")
	host := flag.String("host", "localhost", "the database host")
	port := flag.Int("port", 3306, "the database port")
	user := flag.String("user", "root", "the database user")
	password := flag.String("password", "", "the database password")
	schema := flag.String("schema", "", "the database schema")
	table := flag.String("table", "*", "the table schema，multiple tables ',' split. ")
	serviceName := flag.String("service_name", *schema, "the protobuf service name , defaults to the database schema.")
	packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database schema.")
	ignoreTableStr := flag.String("ignore_tables", "", "a comma spaced list of tables to ignore")
	ignoreColumnStr := flag.String("ignore_columns", "", "a comma spaced list of mysql columns to ignore")
	fieldStyle := flag.String("field_style", "sql_pb", "gen protobuf field style, sql_pb | sqlPb")
	multiple := flag.Bool("multiple", false, "service group")
	fileName := flag.String("filename", *schema, "the filename")

	flag.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filename := path.Join(currentDir, *fileName)

	if *schema == "" {
		fmt.Println(" - please input the database schema ")
		return
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *user, *password, *host, *port, *schema)
	db, err := sql.Open(*dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ignoreTables := strings.Split(*ignoreTableStr, ",")
	ignoreColumns := strings.Split(*ignoreColumnStr, ",")

	s, err := core.GenerateSchema(db, *table, ignoreTables, ignoreColumns, *serviceName, *goPackageName, *packageName, *fieldStyle, *multiple)

	if nil != err {
		log.Fatal(err)
	}

	if nil == s {
		log.Fatal(errors.New("s cannot is nil"))
	}
	err = core.Push(filename, fmt.Sprint(s))
	if err != nil {
		log.Fatal(err)
	}
}
