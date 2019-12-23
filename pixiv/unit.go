package pixiv

import (
	"github.com/zyx4843/gojson"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

var AllowTags []string
var BanedTags []string
var PixivCookie string = ""

func HasBanedTags(jsons string) bool {
	_, taga := gojson.Json(jsons).Get("illust").Get("tags").ToArray()
	tags := Arr2String(taga)
	if strings.Contains(tags, "R-18") {
		return true
	}
	if strings.Contains(tags, "R-18G") {
		return true
	}
	for x := 0; x < len(BanedTags); x++ {
		if strings.Contains(tags, BanedTags[x]) {
			return true
		}
	}
	for i := 0; i < len(AllowTags); i++ {
		if strings.Contains(tags, AllowTags[i]) {
			return false
		}
	}
	return true
}

func Arr2String(arrs []string) string {
	arrlen := len(arrs)
	var strs string
	for i := 0; i < arrlen; i++ {
		strs += arrs[i]
		if i != arrlen-1 {
			strs += ","
		}
	}
	return strs
}

func IsIllust(jsons string) bool {
	if strings.EqualFold(gojson.Json(jsons).Get("illust").Get("type").Tostring(), "illust") {
		return true
	}
	return false
}

func Orig2Mast(originalurl string) string {
	paths, _ := filepath.Split(originalurl)
	filenameWithSuffix := path.Base(originalurl)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return strings.Replace(paths, "img-original", "img-master", 1) + filenameOnly + "_master1200.jpg"

}

func IsURL(urls string) bool {
	u, err := url.Parse(urls)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	return true

}

func HttpGet(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PixivHttpGet(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header.Add("Referer", "https://www.pixiv.net/")
	req.Header.Add("Cookie", PixivCookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
