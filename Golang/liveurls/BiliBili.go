// Package liveurls
// @Time:2023/02/06 01:19
// @File:BiliBili.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

//B站源码来自项目：https://github.com/driveLiving/liveRedirect

package liveurls

import (
	"github.com/asmcos/requests"
	jsoniter "github.com/json-iterator/go"
)

type BiliBili struct {
	Rid string
}

func u(roomId string, pf string, cookies string) any {
	req := requests.Requests()

	fUrl := "https://api.live.bilibili.com/xlive/web-room/v1/playUrl/playUrl"
	p := requests.Params{
		"cid":           roomId,
		"qn":            "10000",
		"platform":      pf,
		"https_url_req": "1",
		"ptype":         "16",
	}
	res, err := req.Get(fUrl, p, requests.Header{"Cookie": cookies})
	if err != nil {
		return nil
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	durl := json.Get([]byte(res.Text()), "data", "durl")
	if durl == nil || durl.Size() == 0 {
		return nil
	}
	s := durl.Get(durl.Size() - 1).Get("url").ToString()
	return s
}
func (b *BiliBili) GetPlayUrl() any {
	rUrl := "https://api.live.bilibili.com/room/v1/Room/room_init?id=" + b.Rid
	resp, err := requests.Get(rUrl)
	if err != nil {
		return nil
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	pageResult := resp.Text()
	code := json.Get([]byte(pageResult), "code").ToInt()
	if code != 0 {
		return nil
	}
	cookies := resp.Cookies()
	coo := ""
	for _, c := range cookies {
		coo = coo + c.Name + "=" + c.Value + ";"
	}

	liveStatus := json.Get([]byte(pageResult), "data", "live_status").ToInt()
	if liveStatus != 1 {
		return nil
	}
	roomId := json.Get([]byte(resp.Text()), "data", "room_id").ToString()

	return u(roomId, "web", coo)
}
