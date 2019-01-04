package model

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
	"chatRoom/common/message"
)


var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//factory mode
func NewUserDao(pool *redis.Pool)(userDao *UserDao)  {
	userDao =&UserDao{
		pool:pool,
	}
	return
}


func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误!
		if err == redis.ErrNil { //表示在 users 哈希中，没有找到对应id
			err = ERROR_USER_NOTEXITS
		}
		return
	}
	user = &User{}
	//这里我们需要把res 反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}


func (this *UserDao) Login(userId int,userPwd string)(user *User,err error)  {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err!=nil{
		return
	}
	if user.UserPwd!=userPwd{
		err=ERROR_USER_PWD
		return
		}
return
}

func (this *UserDao) Register(user *message.User)(err error)  {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err==nil{
		fmt.Println(err)
		err =ERROR_USER_EXITS
		return
	}
	data, err := json.Marshal(user)
	if err!=nil{
		return
	}
	//入库
	_ ,err = conn.Do("HSET", "users", user.UserId, string(data))
	if err!=nil{
		fmt.Println("保存失败 err=",err)
		return
	}

	return
}