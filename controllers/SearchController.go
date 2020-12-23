package controllers

import "github.com/astaxie/beego"

type SearchController struct {
	beego.Controller
}

func (s *SearchController)Get()  {
	s.TplName ="bcsearch.html"
}
