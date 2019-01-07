package process

import (
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
	"chatRoom/client/utils"
)

type SmsProcess struct {

}

//发送群聊的消息

func (this *SmsProcess) SendGroupMes(content string)(err error)  {
		var mes message.Message
		mes.Type=message.SmsMesType
		var smsMes message.SmsMes
		smsMes.Content=content
		smsMes.UserId=CurUser.UserId
		smsMes.UserStatus=CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err!=nil{
		fmt.Println("send group mes err=",err)
		return
	}
	mes.Data=string(data)

	data, err = json.Marshal(mes)
	if err!=nil{
		fmt.Println("send group mes err=",err)
		return
	}

	tf :=utils.Transfer{
		Conn:CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err!=nil{
		fmt.Println("send mess err=",err.Error())
		return
	}
	return
}