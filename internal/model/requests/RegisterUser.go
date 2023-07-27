package requests

type RegisterUser struct {
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccountType string `json:"accountType"`
}
