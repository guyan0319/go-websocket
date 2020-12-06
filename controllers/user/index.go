package user

import (
	"github.com/gin-gonic/gin"
	"go-websocket/controllers"
	"go-websocket/lib/response"
	"go-websocket/model"
)

func List(c *gin.Context)  {
	//appIdStr := c.Query("appId")
	//appId, _ := strconv.ParseInt(appIdStr, 10, 32)
	//
	//fmt.Println("http_request 查看全部在线用户", appId)
	//
	//data := make(map[string]interface{})
	//
	//userList := websocket.UserList()
	//data["userList"] = userList
	//
	//controllers.Response(c, common.OK, "", data)
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