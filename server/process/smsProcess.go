package process

import (
	"chatRoom/common/message"
	"net"
	"encoding/json"
	"fmt"
	"chatRoom/server/utils"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err!=nil{
		fmt.Println("json unmarhsal err=",err)
		return
	}

	data,err:=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
	}

	for id ,up:=range userMgr.onlineUsers{
		if id ==smsMes.UserId{
		   continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn){
	tf :=utils.Transfer{
		Conn:conn,
	}
	err := tf.WritePkg(data)
	if err!=nil{
		fmt.Println("send error =",err)
	}
}