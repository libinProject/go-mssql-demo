package models

import (
	"database/sql"
	"strings"
	"time"

	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"
)

const (
	CONN_LIVE_TIME = 24 //连接使用时间 小时
)

var (
	db *sql.DB = nil //全局数据库连接
)

func GetSqlDb() *sql.DB {

	host := beego.AppConfig.String("db.host")
	port, err := beego.AppConfig.Int("yr_port")
	if err != nil {
		port = 1433
	}
	user := beego.AppConfig.String("db.user")
	password := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")

	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s;encrypt=disable", host, port, dbName, user, password)

	db, err = sql.Open("mssql", connString)
	if err != nil {
		return nil
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Duration(CONN_LIVE_TIME) * time.Hour)

	err = db.Ping()
	if err != nil {
		fmt.Print("PING:%s", err)
		return nil
	}
	return db
}

// 存储过程拼接sql 不带返回值
func GetProcSql(proc string, ins map[string]string) string {

	_sql := fmt.Sprintf("exec %v ", proc)

	var in string
	for key, value := range ins {
		in += key + "=" + value
		in += ","
	}

	if in != "" {
		in = strings.TrimRight(in, ",")
		_sql = fmt.Sprintf("%v %v", _sql, in)
	}
	return _sql
}

// 分页 存储过程凭借sql
func GetPageProcSql(proc string, ins map[string]string) string {

	var _sql string
	declare := "declare @TotalCount int"

	_sql = fmt.Sprintf("%v;exec %v ", declare, proc)

	var in string
	for key, value := range ins {
		in += key + "=" + value
		in += ","
	}

	if in != "" {
		in = strings.TrimRight(in, ",")
		_sql = fmt.Sprintf("%v %v", _sql, in)
	}

	outparam := "@TotalCount=@TotalCount OUTPUT"
	_sql = fmt.Sprintf("%v,%v;select 'returnValue'=@TotalCount;", _sql, outparam)
	return _sql
}
