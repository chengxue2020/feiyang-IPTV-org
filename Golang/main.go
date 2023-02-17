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
	"github.com/forgoer/openssl"
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
		c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyinobj.GetRealurl()))
	})

	r.GET("/:path/:rid", func(c *gin.Context) {
		path := c.Param("path")
		rid := c.Param("rid")
		switch path {
		case "douyin":
			douyinobj := &liveurls.Douyin{}
			douyinobj.Rid = rid
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyinobj.GetDouYinUrl()))
		case "douyu":
			douyuobj := &liveurls.Douyu{}
			douyuobj.Rid = rid
			douyuobj.Stream_type = c.DefaultQuery("stream", "hls")
			douyuobj.Cdn_type = c.DefaultQuery("cdn", "akm-tct")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyuobj.GetRealUrl()))
		case "huya":
			huyaobj := &liveurls.Huya{}
			huyaobj.Rid = rid
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, huyaobj.GetLiveUrl()))
		case "bilibili":
			biliobj := &liveurls.BiliBili{}
			biliobj.Rid = rid
			biliobj.Platform = c.DefaultQuery("platform", "web")
			biliobj.Quality = c.DefaultQuery("quality", "10000")
			biliobj.Line = c.DefaultQuery("line", "second")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, biliobj.GetPlayUrl()))
		case "youtube":
			ytbObj := &liveurls.Youtube{}
			ytbObj.Rid = rid
			ytbObj.Quality = c.DefaultQuery("quality", "1080")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, ytbObj.GetLiveUrl()))
		}
	})
	return r
}

func main() {
	key := []byte("6354127897263145")
	defstr, _ := base64.StdEncoding.DecodeString("Mf5ZVkSUHH5xC9fH2Sao+2LgjRfydmzMgHNrVYX4AcSoI0nktkV7z1jSU6nSihf7ny+PexV73YjDoEtG7qu+Cw==")
	defurl, _ := openssl.AesECBDecrypt(defstr, key, openssl.PKCS7_PADDING)
	r := setupRouter(string(defurl))
	r.Run(":35455")
}
