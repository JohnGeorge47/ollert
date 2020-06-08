package models

type User struct {
	EmailId    string `json:"email_id"`
	UserName   string `json:"user_name"`
	AccountId  int64  `json:"account_id"`
	IsAdmin    bool   `json:"is_admin"`
	Created_At string `json:"created_at"`
	updated_at string `json:"updated_at"`
}


func CreateUser(user User){
	
}