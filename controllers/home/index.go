package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
		data := gin.H{
			"title":        "一对一聊天首页",
		}
		c.HTML(http.StatusOK, "index.html", data)
}
func Room(c *gin.Context)  {
		data := gin.H{
			"title":        "聊天室",
		}
		c.HTML(http.StatusOK, "room.html", data)
}
