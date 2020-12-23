package main

import (
	"BtWeb/bccommand"
	"BtWeb/db_mysql"
	_ "BtWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()
    address:=bccommand.GetAddress()
    fmt.Println(address)
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

    //result,err:=GetBlockCount()
	//if err!=nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(result)
	beego.Run()
}

