package models

import (
	"fmt"
	"github.com/JohnGeorge47/ollert/pkg/sql"
)

type User struct {
	EmailId    string `json:"email_id"`
	UserName   string `json:"user_name"`
	AccountId  int64  `json:"account_id"`
	IsAdmin    bool   `json:"is_admin"`
	Created_At string `json:"created_at"`
	updated_at string `json:"updated_at"`
}

func CreateUser(user User) error {
	query := "INSERT INTO users(email_id,user_name,account_id,is_admin,created_at) VALUES(?,?,?,?,?)"
	conn := sql.Connmanager
	lastid, err := conn.Insert(query, user.EmailId, user.UserName, user.AccountId, user.IsAdmin, user.Created_At)
	if err != nil {
		return err
	}
	fmt.Println(lastid)
	return nil
}
