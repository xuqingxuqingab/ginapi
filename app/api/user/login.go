package user

type RegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRes struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
