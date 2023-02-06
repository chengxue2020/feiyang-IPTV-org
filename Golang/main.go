// Package Golang
// @Time:2023/02/03 02:27
// @File:main.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package main

import (
	"Golang/liveurls"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func duanyan(adurl string, realurl any) string {
	var liveurl string
	if str, ok := realurl.(string); ok {
		liveurl = str
	} else {
		liveurl = adurl
	}
	return liveurl
}

func setupRouter(adurl string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/douyin", func(c *gin.Context) {
		url := c.Query("url")
		quality := c.DefaultQuery("quality", "origin")
		douyinobj := &liveurls.Douyin{}
		douyinobj.Shorturl = url
		douyinobj.Quality = quality
		dyliveurl := duanyan(adurl, douyinobj.GetRealurl())
		c.Redirect(http.StatusMovedPermanently, dyliveurl)
	})

	r.GET("/:path/:rid", func(c *gin.Context) {
		path := c.Param("path")
		rid := c.Param("rid")
		switch path {
		case "douyin":
			douyinobj := &liveurls.Douyin{}
			douyinobj.Rid = rid
			dyliveurl := duanyan(adurl, douyinobj.GetDouYinUrl())
			c.Redirect(http.StatusMovedPermanently, dyliveurl)
		case "douyu":
			douyuobj := &liveurls.Douyu{}
			douyuobj.Rid = rid
			douyuobj.Stream_type = c.DefaultQuery("stream", "hls")
			douyuobj.Cdn_type = c.DefaultQuery("cdn", "akm-tct")
			douyuurl := duanyan(adurl, douyuobj.GetRealUrl())
			c.Redirect(http.StatusMovedPermanently, douyuurl)
		case "huya":
			huyaobj := &liveurls.Huya{}
			huyaobj.Rid = rid
			huyaurl := duanyan(adurl, huyaobj.GetLiveUrl())
			c.Redirect(http.StatusMovedPermanently, huyaurl)
		case "bilibili":
			c.String(http.StatusOK, "待重写中")
		}
	})
	return r
}

func main() {
	defurl, _ := base64.StdEncoding.DecodeString("aHR0cDovLzE1OS43NS44NS42Mzo1NjgwL2QvYWQvcm9vbWFkL3BsYXlsaXN0Lm0zdTg=")
	r := setupRouter(string(defurl))
	r.Run(":35455")
}
