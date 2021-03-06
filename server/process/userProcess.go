package process

import (
	"net"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
	"chatRoom/server/utils"
	"chatRoom/server/model"
)

type UserProcess struct {
   Conn net.Conn
   UserId int
}



func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error)  {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err!=nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}
	var resMes message.Message
	resMes.Type=message.LoginResMesType
	var loginResMes message.LoginResMes
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code=200
	//	loginResMes.Error=""
	//}else{
	//	loginResMes.Code=500
	//	loginResMes.Error="用户不存在，请注册再使用"
	//}

	user, e := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if e!=nil{
		if err==model.ERROR_USER_NOTEXITS{
			loginResMes.Code=500
			loginResMes.Error=err.Error()
		}else if err==model.ERROR_USER_PWD{
			loginResMes.Code=403
			loginResMes.Error=err.Error()
		}else {
			loginResMes.Code=505
			loginResMes.Error="服务器内部错误"
		}
	}else {
		loginResMes.Code=200
		//将登录成功的用户userId给 this
		this.UserId=loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他用户该用户上线
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//当前用户的id 放入loginResMes ids
		//遍历userids
		for id,_:=range userMgr.onlineUsers{
			loginResMes.UserIds=append(loginResMes.UserIds,id)
		}
		fmt.Println(user,"登录成功")
	}


	//seralize
	data,err:=json.Marshal(loginResMes)
	if err!=nil{
		fmt.Println("json marshal fail",err)
		return
	}

	resMes.Data=string(data)
	data, err= json.Marshal(resMes)
	if err!=nil{
		fmt.Println("json marshal fail",err)
		return
	}
	tf :=&utils.Transfer{
		Conn:this.Conn,

	}
	err =tf.WritePkg(data)
	return
}


//编写通知推送
func (this *UserProcess) NotifyOthersOnlineUser(userId int)  {
	//遍历onlineUsers
	for id,up :=range userMgr.onlineUsers{
		//过滤自己
		if id ==userId{
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess)NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type =message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId=userId
	notifyUserStatusMes.Status=message.UserOnline
	data,err :=json.Marshal(notifyUserStatusMes)
	if err!=nil{
		fmt.Println("json.Marshal error=",err)
		return
	}
	mes.Data=string(data)

	data,err =json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal error=",err)
		return
	}
	tf :=&utils.Transfer{
		Conn:this.Conn,
	}
	err =tf.WritePkg(data)
	if err!=nil{
		fmt.Println("NotifyMeOnline err=",err)
		return
	}

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err!=nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}
	var resMes message.Message
	resMes.Type=message.LoginResMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	print(err)
	if err!=nil{
		if err==model.ERROR_USER_EXITS{
			registerResMes.Code=505
			registerResMes.Error=model.ERROR_USER_EXITS.Error()
		}else {
			registerResMes.Code=506
			registerResMes.Error="注册发送未知错误"
		}
	}else {
		registerResMes.Code=200
	}

	data,err:=json.Marshal(registerResMes)
	if err!=nil{
		fmt.Println("json.Marshal fail" ,err)
		return
	}
	resMes.Data=string(data)
	data,err= json.Marshal(resMes)
	if err!=nil{
		fmt.Println("json.Marshal fail",err)
		return
	}
	tf :=&utils.Transfer{
		Conn:this.Conn,
	}
	err=tf.WritePkg(data)
	return

}
