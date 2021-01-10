package request

type Register struct {
	Email        string `json:"email" binding:"required,checkEmail"`
	Username	 string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required,gte=6"`
	Verification string `json:"verification" binding:"required"`
}
