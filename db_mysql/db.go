package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db*sql.DB


//连接数据库
func ConnectDB()  {


	config :=beego.AppConfig
	dbDriver :=config.String("db_driverName")
	dbUser :=config.String("db_user")
	dbPassword :=config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName :=config.String("db_name")
	//连接数据库，用sql.Open（）函数，该函数有两个参数
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err:=sql.Open(dbDriver,connUrl)
	if err!=nil{
		panic("数据库连接错误，请检查")
	}
	Db=db
}
