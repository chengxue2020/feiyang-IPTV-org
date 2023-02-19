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
	"sync"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/etherlabsio/go-m3u8/m3u8"
)

var streamCachedMap sync.Map

type Youtube struct {
	//https://www.youtube.com/watch?v=cK4LemjoFd0
	//Rid: cK4LemjoFd0
	Rid     string
	Quality string
}

func (y *Youtube) GetLiveUrl() any {
	if cached, ok := get(y.Rid); ok {
		return cached
	}
	//proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	client := &http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		//Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}
	r, _ := http.NewRequest("GET", fmt.Sprintf("https://m.youtube.com/watch?v=%v", y.Rid), nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1")
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
	quality := y.getResolution(stream)
	if quality != nil {
		return *quality
	}
	return stream
}

func (y *Youtube) getResolution(liveurl string) *string {
	client := &http.Client{Timeout: time.Second * 5}
	r, _ := http.NewRequest("GET", liveurl, nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	resp, err := client.Do(r)
	if err != nil {
		return nil
	}
	playlist, err := m3u8.Read(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil
	}

	playlists := playlist.Playlists()

	if len(playlists) < 1 {
		return nil
	}

	mapping := map[string]string{}
	for _, item := range playlists {
		mapping[strconv.Itoa(item.Resolution.Height)] = item.URI
	}

	if stream, ok := mapping[y.Quality]; ok {
		set(y.Rid, stream, 600)
		return &stream
	}

	stream := playlists[len(playlists)-1].URI
	set(y.Rid, stream, 600)
	return &stream
}

func set(key string, data interface{}, timeout int) {
	streamCachedMap.Store(key, data)
	time.AfterFunc(time.Second*time.Duration(timeout), func() {
		streamCachedMap.Delete(key)
	})
}

func get(key string) (interface{}, bool) {
	return streamCachedMap.Load(key)
}
