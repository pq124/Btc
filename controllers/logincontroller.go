package controllers

import (
	"BtWeb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController)Get()  {
	l.TplName = "login.html"
}


func (l *LoginController)Post()  {
  var user models.User

   err :=l.ParseForm(&user)
	if err!=nil {
		l.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("该获取得电话号码为:",user.Telephone)
   _,err=user.QueryData()
	if err!=nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("登陆失败")
		return
	}
	l.TplName = "bcsearch.html"
}