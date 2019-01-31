package models

import (
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id          int
    Username        string
    Password        string
    Token        string
    Email        string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
    Id          int
    Age         int16
    User        *User   `orm:"reverse(one)"` // Reverse relationship (optional)
}

type Test struct {
    Id          int
	Age         int16
}
func init() {
    // Need to register model in init
	orm.RegisterModel(new(User), new(Profile), new(Test))
	
}
