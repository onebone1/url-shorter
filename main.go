package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
  URL string
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "reurl.html", nil)
}

func reurl(c *gin.Context) {
  var url string
	if in, isExist := c.GetPostForm("url"); isExist && in != "" {
		url = in
	} else {
		c.HTML(http.StatusBadRequest, "reurl.html", gin.H{
			"error": errors.New("Please fill a url"),
		})
		return
	}
  c.HTML(http.StatusOK, "reurl.html", gin.H{
    "short_url": url,
  })
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("front_end/html/*")
	server.Static("/assets", "front_end/assets")
	server.GET("/", home)
  server.GET("/reurl", home)
  server.POST("/reurl", reurl)
	server.Run(":8888")
}
