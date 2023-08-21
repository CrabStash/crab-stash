package db

type Ownership struct {
	Owns []string `json:"owns"`
}
type UserCrucial struct {
	Email  string `json:"email"`
	Id     string `json:"id"`
	Passwd string `json:"passwd"`
}
