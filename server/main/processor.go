package main

import (
	"net"
	"chatRoom/common/message"
	"fmt"
	process2 "chatRoom/server/process"
	"io"
	"chatRoom/server/utils"
)

type Processor struct {
	Conn  net.Conn
}


func (this *Processor)serverProcessMes(mes *message.Message)(err error)  {

	switch mes.Type {
	case message.LoginMesType:
		up :=&process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMes:

	default:
		fmt.Println("不支持此消息类型")
	}
	return
}

func (this *Processor)process2()(err error) {
	for {
		tf :=&utils.Transfer{
			Conn:this.Conn,
		}
		mes,err:=tf.ReadPkg()
		if err!=nil{
			if err==io.EOF{
				fmt.Println("客户端退出了，我也正常关闭")
				return err
			}else {
				fmt.Println("readPkg error ",err)
				//err=errors.New("readPkg error")
				return err
			}
		}
		//fmt.Println("mes=",mes)
		err=this.serverProcessMes(&mes)
		if err!=nil{
			return err
		}

	}
}