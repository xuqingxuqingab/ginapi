package user

type SayHelloReq struct {
	Name string `json:"name"`
}

type SayHelloRes struct {
	Message string `json:"message"`
}