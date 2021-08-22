package route

import (
	"github.com/gin-gonic/gin"

	"url-shorter/api"
)

func Route() {
	server := gin.Default()
	server.LoadHTMLGlob("views/*")
	server.Static("/assets", "public/assets")
	server.GET("/", api.Home)
  	server.POST("/", api.Reurl);
  //server.GET("/reurl", home) // convert into API
  //server.POST("/reurl", reurl)
	server.Run(":8888")
}