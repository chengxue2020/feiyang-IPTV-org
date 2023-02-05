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

func setupRouter(adurl string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/douyin", func(c *gin.Context) {
		url := c.Query("url")
		quality := c.DefaultQuery("quality", "origin")
		var dyliveurl string
		douyinobj := &liveurls.Douyin{}
		douyinobj.Shorturl = url
		douyinobj.Quality = quality
		dyurl := douyinobj.GetRealurl()
		if str, ok := dyurl.(string); ok {
			dyliveurl = str
		} else {
			dyliveurl = adurl
		}
		c.Redirect(http.StatusMovedPermanently, dyliveurl)
	})

	r.GET("/:path/:rid", func(c *gin.Context) {
		path := c.Param("path")
		rid := c.Param("rid")
		switch path {
		case "douyin":
			var dyliveurl string
			douyinobj := &liveurls.Douyin{}
			douyinobj.Rid = rid
			dyurl := douyinobj.GetDouYinUrl()
			if str, ok := dyurl.(string); ok {
				dyliveurl = str
			} else {
				dyliveurl = adurl
			}
			c.Redirect(http.StatusMovedPermanently, dyliveurl)
		case "douyu":
			var douyuurl string
			douyuobj := &liveurls.Douyu{}
			douyuobj.Rid = rid
			douyuobj.Stream_type = c.DefaultQuery("stream", "flv")
			douyuobj.Cdn_type = c.DefaultQuery("cdn", "akm-tct")
			douyuliveurl := douyuobj.GetRealUrl()
			if str, ok := douyuliveurl.(string); ok {
				douyuurl = str
			} else {
				douyuurl = adurl
			}
			c.Redirect(http.StatusMovedPermanently, douyuurl)
		}
	})
	return r
}

func main() {
	defurl, _ := base64.StdEncoding.DecodeString("aHR0cDovLzE1OS43NS44NS42Mzo1NjgwL2QvYWQvcm9vbWFkL3BsYXlsaXN0Lm0zdTg=")
	r := setupRouter(string(defurl))
	r.Run(":35455")
}
