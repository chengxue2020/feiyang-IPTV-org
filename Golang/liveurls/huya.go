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
	"html"
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

func format(e string) string {
	i := strings.Split(e, "?")[0]
	b := strings.Split(e, "?")[1]
	r := strings.Split(i, "/")
	re := regexp.MustCompile(".(flv|m3u8)")
	s := re.ReplaceAllString(r[len(r)-1], "")
	srcAntiCode := html.UnescapeString(b)
	c := strings.Split(srcAntiCode, "&")
	cc := c[:0]
	n := make(map[string]string)
	for _, x := range c {
		if len(x) > 0 {
			cc = append(cc, x)
			ss := strings.Split(x, "=")
			n[ss[0]] = ss[1]
		}
	}
	c = cc
	fm, _ := url.QueryUnescape(n["fm"])
	uu, _ := base64.StdEncoding.DecodeString(fm)
	u := string(uu)
	p := strings.Split(u, "_")[0]
	f := strconv.FormatInt(time.Now().UnixNano()/100, 10)
	l := n["wsTime"]
	t := "0"
	h := p + "_" + t + "_" + s + "_" + f + "_" + l
	m := md5huya(h)
	url := fmt.Sprintf("%s?wsSecret=%s&wsTime=%s&u=%s&seqid=%s&txyp=%s&fs=%s&sphdcdn=%s&sphdDC=%s&sphd=%s&u=0&t=100&sv=", i, m, l, t, f, n["txyp"], n["fs"], n["sphdcdn"], n["sphdDC"], n["sphd"])
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
	if res == nil {
		return nil
	}
	nstr, _ := base64.StdEncoding.DecodeString(res[1])
	realstr := string(nstr)
	if strings.Contains(realstr, "replay") {
		return "https:" + realstr
	} else {
		liveurl := format(realstr)
		liveurl = strings.Replace(liveurl, "hls", "flv", -1)
		liveurl = strings.Replace(liveurl, "m3u8", "flv", -1)
		return "https:" + liveurl
	}
}
