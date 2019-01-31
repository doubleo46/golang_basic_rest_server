package main

import (
	_ "golang_basic_rest_server/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)


func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RunSyncdb("default", false, true)
	orm.RunCommand()

}


func main() {
	orm.Debug = true

	beego.Run()
}

