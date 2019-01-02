package main

import (
	"fmt"
	"net"
	"chatRoom/common/message"
	"encoding/json"
	"io"
	"encoding/binary"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8889)
	fmt.Println("读取客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则，就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据 pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}
	//把pkgLen 反序列化成 -> message.Message
	// 技术就是一层窗户纸 &mes！！
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}





func serverProcessLogin(conn net.Conn,mes *message.Message)(err error)  {
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
	err =writePkg(conn, data)
	return
}



func serverProcessMes(conn net.Conn,mes *message.Message)(err error)  {
	switch mes.Type {
		case message.LoginMesType:
			err = serverProcessLogin(conn, mes)
		case message.RegisterMes:

		default:
			fmt.Println("不支持此消息类型")
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		mes,err:=readPkg(conn)
		if err!=nil{
			if err==io.EOF{
				fmt.Println("客户端退出了，我也正常关闭")
				return
			}else {
				fmt.Println("readPkg error ",err)
				//err=errors.New("readPkg error")
				return
			}
		}
		//fmt.Println("mes=",mes)
		err=serverProcessMes(conn,&mes)
		if err!=nil{
			return
		}

	}
}

func main() {
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
