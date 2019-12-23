package pixiv

import (
	"OneeSan/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/zyx4843/gojson"
	"strconv"
	"strings"
	"time"
)

var BooksLimit int = 1000
var AllowAdd bool = true
var DiscoEveryDelay int = 100

func AddRandomIllust() {
	illust := RandomIllust()
	if gojson.Json(illust).IsValid() {
		ipid := gojson.Json(illust).Get("pid").Tostring()
		illustid, err := strconv.Atoi(ipid)
		if CheckError(err) {
			return
		}
		if !CheckError(AddPixivIllust(illustid)) {
			beego.Info("Add [" + ipid + "] to Database.")
		}
	}
}

func AddPixivIllust(pid int) error {
	if !AllowAdd {
		return errors.New("当前不允许添加插画")
	}
	illust := IllustInfo(pid)
	if IsIllust(illust) {
		if !HasBanedTags(illust) {
			books, err := strconv.Atoi(gojson.Json(illust).Get("illust").Get("total_bookmarks").Tostring())
			if CheckError(err) {
				return err
			}
			if books >= BooksLimit {
				origurl := gojson.Json(illust).Get("illust").Get("meta_single_page").Get("original_image_url").Tostring()
				morigurl := gojson.Json(illust).Get("illust").Get("meta_pages").Getindex(1).Get("image_urls").Get("original").Tostring()
				if origurl != "" || morigurl != "" {
					if origurl == "" {
						origurl = morigurl
					}
					o := orm.NewOrm()
					ipid := gojson.Json(illust).Get("illust").Get("id").Tostring()
					if !o.QueryTable("illust_info").Filter("pid", ipid).Exist() {
						ill := models.IllustInfo{}
						ill.Pid, err = strconv.Atoi(ipid)
						if CheckError(err) {
							return err
						}
						ill.Title = gojson.Json(illust).Get("illust").Get("title").Tostring()
						ill.Author = gojson.Json(illust).Get("illust").Get("user").Get("name").Tostring()
						ill.OriginalUrl = strings.Replace(origurl, "pximg.net", "pixiv.cat", 1)
						ill.MasterUrl = Orig2Mast(ill.OriginalUrl)
						if ill.Title != "" && ill.Author != "" && IsURL(ill.OriginalUrl) && IsURL(ill.MasterUrl) {
							_, err = o.Insert(&ill)
							if CheckError(err) {
								return err
							}
						} else {
							return errors.New("数据出现问题，无法添加到数据库")
						}
					} else {
						return errors.New("插画已经存在")
					}
				} else {
					return errors.New("插画链接无效")
				}
			} else {
				return errors.New("插画收藏数小于 " + strconv.Itoa(BooksLimit))
			}
		} else {
			return errors.New("插画标签不包含关键词，或是R18/R18G")
		}
	} else {
		return errors.New("不是插画或ID无效")
	}
	return nil
}

func IllustDiscovery() {
	illlist := IllustDiscoList()
	if illlist != "" {
		for i := 1; gojson.Json(illlist).Get("recommendations").Getindex(i).IsValid(); i++ {
			pids := gojson.Json(illlist).Get("recommendations").Getindex(i).Tostring()
			illustid, err := strconv.Atoi(pids)
			if !CheckError(err) {
				if !CheckError(AddPixivIllust(illustid)) {
					beego.Info("Add [" + pids + "] to Database.")
				}
			}
			time.Sleep(time.Duration(DiscoEveryDelay) * time.Millisecond)
		}
	}
}

func DBCount() int64 {
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("SELECT MAX(id) FROM illust_info;").Values(&maps, "MAX(id)")
	if CheckError(err) {
		return -1
	}
	count, err := strconv.ParseInt(maps[0]["MAX(id)"].(string), 10, 64)
	if CheckError(err) {
		return -1
	}
	return count
}

func CheckError(err error) bool {
	if err != nil {
		beego.Error(err)
		return true
	} else {
		return false
	}

}
