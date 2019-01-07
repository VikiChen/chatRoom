package process

import "fmt"

//UserMgr 在server 有且只有一个，声明为全局变量

var (
	userMgr *UserMgr
)


type UserMgr struct {
	onlineUsers map[int]*UserProcess
}


//init

func init()  {
	userMgr =&UserMgr{
		onlineUsers:make(map[int]*UserProcess,1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess)  {
	this.onlineUsers[up.UserId]=up
}


func (this *UserMgr) DelOnlineUser(userId int)  {
	delete(this.onlineUsers,userId)
}

func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess  {
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUserById(userId int) ( up *UserProcess,err error)  {
	up ,ok :=  this.onlineUsers[userId]
	if ! ok{
		err=fmt.Errorf("用户不存在")
		return
		}
		return
}
