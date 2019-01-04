package main

import (
	"fmt"
	"net"
	"time"
	"chatRoom/server/model"
)


func process(conn net.Conn) {
	defer conn.Close()
    processor :=&Processor{
		Conn:conn,
	}
	err := processor.process2()
	if err!=nil{
		fmt.Println("error")
		return
	}
}

func initUserDao()  {
	model.MyUserDao =model.NewUserDao(pool)
}


func main() {
	initPool("localhost:6379",16,2,300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听。。。")
	listener, e := net.Listen("tcp", "0.0.0.0:8889")
	defer listener.Close()
	if e != e {
		fmt.Println("net listener error", e)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, i := listener.Accept()
		if i != nil {
			fmt.Println("accept error", i)
		}
		//连接成功
		go process(conn)

	}

}
