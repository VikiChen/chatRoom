package main

import (
	"fmt"
	"chatRoom/client/process"
)

//2 param ：1.userId 2.userPwd
var (
	userId int
	userPwd string
	userName string
)


func main() {
	var key int //接收用户选择
	var loop = true
	for loop {
		fmt.Println("-------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t\t 1 登陆聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 4 请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n",&userPwd)
			//complete login
			up:=&process.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("输入用户id：")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("输入用户密码：")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Println("输入用户昵称：")
			fmt.Scanf("%s\n",&userName)
			up:=&process.UserProcess{}
			up.Register(userId,userPwd,userName)
		case 3:
			fmt.Println("退出")

		default:
			fmt.Println("输入有误重新输入")
		}
	}

}