package user

import (
	"github.com/gin-gonic/gin"
	"go-websocket/controllers"
	"go-websocket/lib/response"
	"go-websocket/model"
)

func List(c *gin.Context)  {
	groupsIdStr := c.Query("groupsId")//获取组id
	data := make(map[string]interface{})
	userList := model.GetUsersAll(groupsIdStr)
	data["userList"] = userList
	controllers.ShowSuccess(c,data)
}
func GroupsList(c *gin.Context)  {
	appIdStr := c.Query("appId")
	if appIdStr=="" || appIdStr=="null"{
		controllers.ShowError(c,response.AppIDNotFound,"")
		return
	}
	list,err :=model.GetGroupsAll(appIdStr)
	if err !=nil{
		controllers.ShowError(c,response.AppIDNotFound,"")
		return
	}
	controllers.ShowSuccess(c,list)
}