// Package Golang
// @Time:2023/02/03 02:27
// @File:main.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package main

import (
	"Golang/liveurls"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
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
			}
			c.Redirect(http.StatusMovedPermanently, dyliveurl)
		}
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":35455")
}
