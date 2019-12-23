package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	IndexController
}

func (c *ErrorController) Error404() {
	c.Data["Path"] = c.Ctx.Request.RequestURI
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error500() {
	c.Data["Title"] = "500 Internal Server Error"
	c.Data["Info"] = "服务器娘你在干嘛，不要停下来啊"
	c.TplName = "error/error.html"
}

func (c *ErrorController) Error503() {
	c.Data["Title"] = "503 Service Unavailable"
	c.Data["Info"] = "唔，服务器娘现在忙不过来啦，请稍后尝试QAQ"
	c.TplName = "error/error.html"
}

func (c *ErrorController) Error403() {
	c.Data["Title"] = "403 Forbidden"
	c.Data["Info"] = "服务器娘正在维护或者其他原因无法访问QAQ"
	c.TplName = "error/error.html"
}

func CheckError(err error) bool {
	if err != nil {
		beego.Error(err)
		return true
	}else{
		return false
	}

}