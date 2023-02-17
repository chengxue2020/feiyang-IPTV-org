// Package liveurls
// @Time:2023/02/17 16:32
// @File:youtube.go
// @SoftWare:Goland
// @Author:Popeye
// @Contact:TG@popeyelau

package liveurls

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"io"
	"net/http"
)

type Youtube struct {
	//https://www.youtube.com/watch?v=cK4LemjoFd0
	//Rid: cK4LemjoFd0
	Rid string
}

func (y *Youtube) GetLiveUrl() any {
	//proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		//Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}
	r, _ := http.NewRequest("GET", fmt.Sprintf("https://www.youtube.com/watch?v=%v", y.Rid), nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	str := string(body)

	reg := regexp2.MustCompile(`(?<=hlsManifestUrl":").*\.m3u8`, regexp2.RE2)
	res, _ := reg.FindStringMatch(str)
	if res == nil {
		return nil
	}
	return res.Captures[0].String()
}
