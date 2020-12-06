package controllers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/lib/response"
	"net/http"
)

func ShowError(c *gin.Context,code uint32, msg string){
	var data interface{}
	message :=response.Response(response.OK, msg, data)
	// 设置返回格式是json
	c.JSON(http.StatusOK, message)
	return
}

func ShowSuccess(c *gin.Context,  data interface{}){
	message := response.Response(response.OK, "", data)
	// 设置返回格式是json
	c.JSON(http.StatusOK, message)
	return
}

