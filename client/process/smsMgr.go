package process

import (
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

func outPutGroupMes(mes *message.Message)  {
	//just show
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err!=nil{
		fmt.Println("json unmarshal err=",err)
		return
	}
	//显示消息
	info:=fmt.Sprintf("用户id:\t%d 对大家说：\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
