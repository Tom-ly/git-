package controllers

import "github.com/astaxie/beego"

type RegisterController struct {
	beego.Controller
}

func (r*RegisterController) Post(){
	r.TplName = "login.html"

}