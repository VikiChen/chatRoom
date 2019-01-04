package process

import (
	"net"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
	"chatRoom/server/utils"
)

type UserProcess struct {
   Conn net.Conn
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
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code=200
		loginResMes.Error=""
	}else{
		loginResMes.Code=500
		loginResMes.Error="用户不存在，请注册再使用"
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
