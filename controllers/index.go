package controllers

import (
	"OneeSan/models"
	"OneeSan/pixiv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

var DB_counts int64 = 0
var SafeMode bool = false
var FreqLimit int = 0

var AddCounts int = 0
var RandSql string = "SELECT * FROM `illust_info` AS t1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(id) FROM `illust_info`)-(SELECT MIN(id) FROM `illust_info`))+(SELECT MIN(id) FROM `illust_info`)) AS id) AS t2 WHERE t1.id >= t2.id ORDER BY t1.id LIMIT 10;"

func (c *IndexController) Prepare() {
	if SafeMode {
		c.Data["db_count"] = 0
	} else {
		c.Data["db_count"] = DB_counts
	}

}

// @router / [get]
func (c *IndexController) Index() {
	SafeAbort(c)
	o := orm.NewOrm()
	rs := o.Raw(RandSql)
	var ills []models.IllustInfo
	_, err := rs.QueryRows(&ills)
	if CheckError(err) {
		c.Abort("500")
	}
	c.Data["ills"] = ills
	c.TplName = "index.html"
}

// @router /about [get]
func (c *IndexController) IndexAbout() {
	SafeAbort(c)
	c.Data["books_limit"] = pixiv.BooksLimit
	c.Data["baned_tags"] = pixiv.BanedTags
	c.Data["allow_tags"] = pixiv.AllowTags
	c.TplName = "about.html"
}

// @router /test [get]
func (c *IndexController) IndexTest() {
	c.Ctx.WriteString("唔，被你发现彩蛋啦，只是个测试页面啦，什么都没有哦qwq")

}

// @router /robots.txt [get]
func (c *IndexController) IndexRobots() {
	c.Ctx.Output.Header("Content-Type", "text/plain")
	c.TplName = "robots.html"

}

// @router /api [get]
func (c *IndexController) IndexApi() {
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw(RandSql).Values(&maps, "pid", "title", "author", "original_url", "master_url")
	if CheckError(err) {
		c.Data["json"] = map[string]interface{}{"error": 500, "illusts": nil}
	} else if SafeMode {
		c.Data["json"] = map[string]interface{}{"error": 403, "illusts": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"error": 0, "illusts": &maps}
	}
	c.ServeJSON()
}

// @router /api/dbcount [get]
func (c *IndexController) ApiDBCount() {
	if SafeMode {
		c.Ctx.WriteString("0")
	} else {
		c.Ctx.WriteString(strconv.FormatInt(DB_counts, 10))
	}
}

// @router /api/addillust [post]
func (c *IndexController) ApiAddIllust() {
	if SafeMode {
		c.Ctx.WriteString("服务器正在维护或者其他原因无法访问")
		return
	}
	AddCounts++
	if AddCounts > FreqLimit {
		c.Ctx.WriteString("系统繁忙中，请稍后重试")
		return
	}

	pids := c.GetString("pid")
	illustid, err := strconv.Atoi(pids)
	if CheckError(err) {
		c.Ctx.WriteString("插画ID必须是数字")
		return
	}
	err = pixiv.AddPixivIllust(illustid)
	if CheckError(err) {
		c.Ctx.WriteString(err.Error())
		return
	}
	beego.Info("Manual Add [" + pids + "] to Database.")
	c.Ctx.WriteString("ok")

}

func SafeAbort(c *IndexController) {
	if SafeMode {
		c.Ctx.Output.Status = 403
		c.Abort("403")
	}
}
