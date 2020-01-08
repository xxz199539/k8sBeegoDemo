package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "k8sBeegoDemo/routers"
)

func init() {
	mysql, err := beego.AppConfig.GetSection("mysql")
	if err != nil {
		fmt.Printf(err.Error())
	}
	dbname := mysql["mysqldbname"]
	dbuser := mysql["mysqldbuser"]
	dbpassword := mysql["mysqldbpassword"]
	dbhost := mysql["mysqldbhost"]
	dbport := mysql["mysqldbport"]
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RunCommand()
}

func main() {

	beego.Run("0.0.0.0:8081")
}
