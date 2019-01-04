package process

import (
	"fmt"
	"os"
	"net"
	"chatRoom/client/utils"
)

func ShowMenu()  {
	fmt.Println("------------恭喜登录成功---------------")
	fmt.Println("------------1.显示在线用户列表---------------")
	fmt.Println("------------2.发送消息---------------")
	fmt.Println("------------3.消息列表---------------")
	fmt.Println("------------4.退出系统---------------")
	fmt.Println("------------请选择（1-4）---------------")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
		os.Exit(0)
	default:
		fmt.Println("你输入有误，重新输入")
	}
}
//keep in touch with server
func  ProcessServerMes(Conn net.Conn)  {
	tf:=&utils.Transfer{
		Conn:Conn,
	}
	for{
		fmt.Println("客户端正在读取服务器端发送的消息")
		mes, err := tf.ReadPkg()
		if err!=nil{
			fmt.Println("read error")
		    return
		}
		fmt.Println("mes=",mes)
	}
}