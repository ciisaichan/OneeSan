package pixiv

import (
	"strconv"
)

func IllustInfo(pid int) string {
	rep, err := HttpGet("https://api.imjad.cn/pixiv/v2/?id=" + strconv.Itoa(pid))
	if CheckError(err) {
		return ""
	}
	return string(rep)
}

func RandomIllust() string {
	rep, err := HttpGet("https://api.lolicon.app/setu/")
	if CheckError(err) {
		return ""
	}
	return string(rep)
}

func IllustDiscoList() string {
	rep, err := PixivHttpGet("https://www.pixiv.net/rpc/recommender.php?type=illust&sample_illusts=auto&num_recommendations=100&page=discovery&mode=safe")
	if CheckError(err) {
		return ""
	}
	return string(rep)
}
