package message

const(
	LoginMesType ="LoginMes"
	LoginResMesType="LoginResMes"
	RegisterMes="RegisterMes"
)

type Register struct {
	//...
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


