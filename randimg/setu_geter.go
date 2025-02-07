package randimg

import (
	"strings"

	"github.com/Yiwen-Chan/ZeroBot-Plugin/api/msgext"
	zero "github.com/wdvxdr1123/ZeroBot"
)

var RANDOM_API_URL = "https://api.pixivweb.com/anime18r.php?return=img"
var BLOCK_REQUEST = false

func init() { // 插件主体
	zero.OnRegex(`^设置随机图片网址(.*)$`, zero.SuperUserPermission).SetBlock(true).SetPriority(20).
		Handle(func(ctx *zero.Ctx) {
			url := ctx.State["regex_matched"].([]string)[1]
			if !strings.HasPrefix(url, "http") {
				ctx.Send("URL非法!")
			} else {
				RANDOM_API_URL = url
			}
			return
		})
	// 随机图片
	zero.OnFullMatchGroup([]string{"随机图片"}).SetBlock(true).SetPriority(24).
		Handle(func(ctx *zero.Ctx) {
			if BLOCK_REQUEST {
				ctx.Send("请稍后再试哦")
			} else {
				BLOCK_REQUEST = true
				ctx.Send(msgext.ImageNoCache(RANDOM_API_URL))
				BLOCK_REQUEST = false
			}
			return
		})
}
