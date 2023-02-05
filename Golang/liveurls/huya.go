// Package liveurls
// @Time:2023/02/05 23:34
// @File:huya.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package liveurls

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Huya struct {
	Rid string
}

func md5huya(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func format(realstr string) string {
	i := strings.Split(realstr, "?")[0]
	b := strings.Split(realstr, "?")[1]
	r := strings.Split(i, "/")
	reg := regexp.MustCompile(".(flv|m3u8)")
	s := reg.ReplaceAllString(r[len(r)-1], "")
	c := strings.SplitN(b, "&", 4)
	cnil := c[:0]
	n := make(map[string]string)
	for _, v := range c {
		if len(v) > 0 {
			cnil = append(cnil, v)
			narr := strings.Split(v, "=")
			n[narr[0]] = narr[1]
		}
	}
	c = cnil
	fm, _ := url.QueryUnescape(n["fm"])
	ub64, _ := base64.StdEncoding.DecodeString(fm)
	u := string(ub64)
	p := strings.Split(u, "_")[0]
	f := strconv.FormatInt(time.Now().UnixNano()/100, 10)
	l := n["wsTime"]
	t := "0"
	h := p + "_" + t + "_" + s + "_" + f + "_" + l
	m := md5huya(h)
	y := c[len(c)-1]
	url := fmt.Sprintf("%s?wsSecret=%s&wsTime=%s&u=%s&seqid=%s&%s", i, m, l, t, f, y)
	return url
}

func (h *Huya) GetLiveUrl() any {
	liveurl := "https://m.huya.com/" + h.Rid
	client := &http.Client{}
	r, _ := http.NewRequest("GET", liveurl, nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	str := string(body)
	freg := regexp.MustCompile(`"(?i)liveLineUrl":"([\s\S]*?)",`)
	res := freg.FindStringSubmatch(str)
	if res == nil || res[1] == "" {
		return nil
	}
	nstr, _ := base64.StdEncoding.DecodeString(res[1])
	realstr := string(nstr)
	if strings.Contains(realstr, "replay") {
		return "https:" + realstr
	} else {
		liveurl := format(realstr)
		realurl := strings.Replace(strings.Replace(strings.Replace(liveurl, "hls", "flv", -1), "m3u8", "flv", -1), "&ctype=tars_mobile", "", -1)
		return "https:" + realurl
	}
}
