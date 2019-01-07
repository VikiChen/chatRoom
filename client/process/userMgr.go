package process

import (
	"chatRoom/common/message"
	"fmt"
	"chatRoom/client/model"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User,10)
var CurUser model.CurUser


func outPutOnlineUser()  {
	fmt.Println("当前在线用户列表")
	for id,_ :=range onlineUsers{
		fmt.Println("用户Id:\t",id)
	}
}


func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes)  {
	//适当优化

	user,ok :=onlineUsers[notifyUserStatusMes.UserId]
	if !ok{
		user =&message.User{
			UserId:notifyUserStatusMes.UserId,

		}
	}
	user.UserStatus =notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId]=user
	outPutOnlineUser()
}