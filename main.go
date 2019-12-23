package main

import (
	"OneeSan/controllers"
	_ "OneeSan/models"
	"OneeSan/pixiv"
	_ "OneeSan/routers"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

func main() {
	getdelay, err := beego.AppConfig.Int("get_illust_delay")
	if controllers.CheckError(err) {
		return
	}
	pixiv.BooksLimit, err = beego.AppConfig.Int("add_books_limit")
	if controllers.CheckError(err) {
		return
	}
	controllers.SafeMode, err = beego.AppConfig.Bool("safe_mode")
	if controllers.CheckError(err) {
		return
	}
	pixiv.AllowAdd, err = beego.AppConfig.Bool("allow_add_illust")
	if controllers.CheckError(err) {
		return
	}
	controllers.FreqLimit, err = beego.AppConfig.Int("add_freq_limit")
	if controllers.CheckError(err) {
		return
	}
	pixiv.DiscoEveryDelay, err = beego.AppConfig.Int("illust_every_delay")
	if controllers.CheckError(err) {
		return
	}
	discodelay, err := beego.AppConfig.Int("illust_disco_delay")
	if controllers.CheckError(err) {
		return
	}
	banedtags := beego.AppConfig.String("baned_tags")
	if len(banedtags) != 0 {
		pixiv.BanedTags = strings.Split(banedtags, ",")
	}
	allowtags := beego.AppConfig.String("allow_tags")
	if len(allowtags) != 0 {
		pixiv.AllowTags = strings.Split(allowtags, ",")
	}
	pixiv.PixivCookie = beego.AppConfig.String("pixiv_cookie")

	go WhileGet(getdelay)
	go WhileCount()
	go WhileResetFreq()
	go WhileDiscoIllust(discodelay)
	beego.Run()
}

func WhileGet(delay int) {
	for delay > 0 {
		pixiv.AddRandomIllust()
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func WhileCount() {
	for {
		count := pixiv.DBCount()
		if count != -1 {
			controllers.DB_counts = pixiv.DBCount()
		}
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

func WhileResetFreq() {
	for {
		controllers.AddCounts = 0
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

func WhileDiscoIllust(delay int) {
	for delay > 0 {
		pixiv.IllustDiscovery()
		time.Sleep(time.Duration(delay) * time.Minute)
	}
}
