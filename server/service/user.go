package service

type User struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
