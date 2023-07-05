package main

type User struct {
	Id     string `json:"id",omitempty`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

type Token struct {
	Token string `json:"token"`
}
