package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type IndexData struct {
	URL string
  }

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "reurl.html", nil)
}

func Reurl(c *gin.Context) {
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