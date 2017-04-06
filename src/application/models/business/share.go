package business

import (
	"../dao"
	"../error"
	"../http/user"
	//"../mysql"
	//"encoding/json"
	"git.oschina.net/pbaapp/goyaf"
	//"git.oschina.net/pbaapp/goyaf/db"
	"strconv"
	"time"
)

type Share struct {
	Base
}

//设置信息
func (this *Share) Add(params map[string]string) map[string]string {
	now := time.Now()
	_, ok := params["add_time"]
	if !ok {
		params["add_time"] = strconv.FormatInt(now.Unix(), 10)
	}
	result := dao.ShareDAO.Insert(params)
	return map[string]string{
		"icon_url": "http://appimg.pba.cn/2014/12/17/60a683df6776c2598ebd66ccf5d0c57b.jpg",
		"url":      goyaf.GetConfigByKey("url.m_pba_cn") + "/index/fatscale/share/share_id/" + strconv.FormatInt(result, 10) + "/",
		"title":    "我已经发现了体重变化的奥秘。",
		"content":  "",
	}
}

//分享的详情
func (this *Share) Content(params map[string]string) map[string]string {
	share := dao.ShareDAO.Find(params["share_id"])
	if len(share) == 0 {
		ourError.PanicError(ourError.ShareIsNotExists)
	}

	userHttp := ourHttpUser.NewUser()
	user, err := userHttp.GetUserInfo(share["user_id"], "nickname,avatar")
	if err != nil {
		ourError.PanicError(ourError.RequestUserError)
	}
	share["user_nickname"] = user["nickname"]
	share["user_avatar"] = user["avatar"]
	return share
}

var ShareBN *Share

func init() {
	goyaf.Debug("init business share")
	ShareBN = &Share{}
}
