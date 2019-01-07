package message

const(
	LoginMesType ="LoginMes"
	LoginResMesType="LoginResMes"
	RegisterMesType="RegisterMes"
	RegisterResMesType="RegisterResMes"
	NotifyUserStatusMesType="NotifyUserStatusMes"
	SmsMesType ="SmsMes"
)

//user status
const (
	UserOnline =iota
	UserOffine
	UserBusyStatus
)

type RegisterMes struct {
	User User
}

type Message struct {
	Type string   `json:"type"`
	Data string    `json:"data"`
}

type LoginMes struct {
	UserId int   `json:"userId"`
	UserPwd string   `json:"userPwd"`
	UserName string   `json:"userName"`
}

type LoginResMes struct {
	Code int    `json:"code"`// 500 not register //  200 success
	UserIds []int//保存用户切片
	Error string   `json:"error"`
}

type RegisterResMes struct {
	Code int    `json:"code"`// 500 existed   200 success
	Error string   `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

