package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"golang_basic_rest_server/models"
)

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

type RespJson struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
  }
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *TestController) Post() {

	o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database

    profile := new(models.Profile)
    profile.Age = 30

    user := new(models.User)
    user.Profile = profile
    user.Username = "slene"

    fmt.Println(o.Insert(profile))
    fmt.Println(o.Insert(user))

    responseJSON := RespJson{
        Value1:    "dataExample",
        Value2:    "dfg",
    }
    this.Data["json"] = &responseJSON
    this.ServeJSON()
}