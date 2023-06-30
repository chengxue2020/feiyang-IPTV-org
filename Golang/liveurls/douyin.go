// Package liveurls
// @Time:2023/02/03 01:59
// @File:douyin.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package liveurls

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type Douyin struct {
	Quality  string
	Shorturl string
	Rid      string
}

func GetRoomId(url string) any {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	r.Header.Add("authority", "v.douyin.com")
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	reurl := resp.Header.Get("Location")
	reg := regexp.MustCompile(`\d{19}`)
	res := reg.FindAllStringSubmatch(reurl, -1)
	if res == nil {
		return nil
	}
	return res[0][0]
}

func (d *Douyin) GetRealurl() any {
	var mediamap map[string]map[string]map[string]map[string]map[string]any
	var roomid string
	if str, ok := GetRoomId(d.Shorturl).(string); ok {
		roomid = str
	} else {
		return nil
	}
	client := &http.Client{}
	params := map[string]string{
		"aid":              "6383",
		"live_id":          "1",
		"device_platform":  "web",
		"language":         "zh-CN",
		"enter_from":       "web_search",
		"cookie_enabled":   "true",
		"screen_width":     "1920",
		"screen_height":    "1080",
		"browser_language": "zh-CN",
		"browser_name":     "Chrome",
		"room_id":          roomid,
		"scene":            "pc_stream_4k",
	}

	r, _ := http.NewRequest("GET", "https://live.douyin.com/webcast/room/info_by_scene/?", nil)
	q := r.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	str, _ := url.QueryUnescape(string(body))
	json.Unmarshal([]byte(str), &mediamap)
	var realurl any
	if mediaslice, ok := mediamap["data"]["stream_url"]["live_core_sdk_data"]["pull_data"]["Hls"].([]any); ok {
		for _, v := range mediaslice {
			if newmediamap, ok := v.(map[string]any); ok {
				for k := range newmediamap {
					switch d.Quality {
					case "uhd":
						{
							if newmediamap[k] == "uhd" {
								realurl = newmediamap["url"]
							}
						}
					case "origin":
						{
							if newmediamap[k] == "origin" {
								realurl = newmediamap["url"]
							}
						}
					case "hd":
						{
							if newmediamap[k] == "hd" {
								realurl = newmediamap["url"]
							}
						}
					case "sd":
						{
							if newmediamap[k] == "sd" {
								realurl = newmediamap["url"]
							}
						}
					case "ld":
						{
							if newmediamap[k] == "ld" {
								realurl = newmediamap["url"]
							}
						}
					}
				}
			}

		}
	}
	return realurl
}

func (d *Douyin) GetDouYinUrl() any {
	liveurl := "https://live.douyin.com/" + d.Rid
	client := &http.Client{}
	var mediamap map[string]map[string]any
	re, _ := http.NewRequest("GET", liveurl, nil)
	re.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	re.Header.Add("upgrade-insecure-requests", "1")
	oresp, _ := client.Do(re)
	defer oresp.Body.Close()
	oreg := regexp.MustCompile(`(?i)__ac_nonce=(.*?);`)
	ores := oreg.FindStringSubmatch(oresp.Header["Set-Cookie"][0])
	r, _ := http.NewRequest("GET", liveurl, nil)
	cookie1 := &http.Cookie{Name: "__ac_nonce", Value: ores[1]}
	r.AddCookie(cookie1)
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	r.Header.Add("upgrade-insecure-requests", "1")
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	str, _ := url.QueryUnescape(string(body))
	reg := regexp.MustCompile(`(?i)\"roomid\"\:\"([0-9]+)\"`)
	res := reg.FindAllStringSubmatch(str, -1)
	if res == nil {
		return nil
	}
	nnreg := regexp.MustCompile(`(?i)\"id_str\":\"` + res[0][1] + `(?i)\"[\s\S]*?\"hls_pull_url\"`)
	nnres := nnreg.FindAllStringSubmatch(str, -1)
	nnnreg := regexp.MustCompile(`(?i)\"hls_pull_url_map\"[\s\S]*?}`)
	nnnres := nnnreg.FindAllStringSubmatch(nnres[0][0], -1)
	if nnnres == nil {
		return nil
	}
	json.Unmarshal([]byte(`{`+nnnres[0][0]+`}`), &mediamap)
	return mediamap["hls_pull_url_map"]["FULL_HD1"]
}
