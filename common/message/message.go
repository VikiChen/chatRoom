package message

const(
	LoginMesType ="LoginMes"
	LoginResMesType="LoginResMes"
	RegisterMesType="RegisterMes"
	RegisterResMesType="RegisterResMes"
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
	Code int    `json:"code"`// 500 not register   200 success
	Error string   `json:"error"`
}

type RegisterResMes struct {
	Code int    `json:"code"`// 500 existed   200 success
	Error string   `json:"error"`
} 