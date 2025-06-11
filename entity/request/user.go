package request

type Register struct {
	Username    string `json:"username" binding:"required,max=20"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	RepeatedPwd string `json:"repeated_pwd" binding:"required,eqfield=Password"`
	Email       string `json:"email" binding:"required,email"`
}

type Login struct {
	Username string `json:"username" binding:"required,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
