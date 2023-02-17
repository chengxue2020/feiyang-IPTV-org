// Package liveurls
// @Time:2023/02/17 16:32
// @File:youtube.go
// @SoftWare:Goland
// @Author:Popeye
// @Contact:TG@popeyelau

package liveurls

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/etherlabsio/go-m3u8/m3u8"
)

type Youtube struct {
	//https://www.youtube.com/watch?v=cK4LemjoFd0
	//Rid: cK4LemjoFd0
	Rid     string
	Quality string
}

func (y *Youtube) GetLiveUrl() any {
	//proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	client := &http.Client{
		Timeout: time.Second * 5,
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
	stream := res.Captures[0].String()
	quality := getResolution(stream, y.Quality)
	if quality != nil {
		return *quality
	}
	return stream
}

func getResolution(liveurl string, quality string) *string {
	client := &http.Client{Timeout: time.Second * 5}
	r, _ := http.NewRequest("GET", liveurl, nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	resp, _ := client.Do(r)
	playlist, err := m3u8.Read(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil
	}

	size := playlist.ItemSize()

	if size < 1 {
		return nil
	}

	mapping := map[string]string{}
	for _, item := range playlist.Playlists() {
		mapping[strconv.Itoa(item.Resolution.Height)] = item.URI
	}

	if stream, ok := mapping[quality]; ok {
		return &stream
	}
	return &playlist.Playlists()[size-1].URI
}
