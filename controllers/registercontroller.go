package controllers

import (
	"BtWeb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterConntroller struct {
	beego.Controller
}

func (r *RegisterConntroller)Post()  {
	var user models.User
    err := r.ParseForm(&user)
	if err!=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析失败,请重试")
		return
	}
     _,err=user.SavaData()
	if err!=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据储存失败,请重试")
	}
	r.TplName = "login.html"
}

