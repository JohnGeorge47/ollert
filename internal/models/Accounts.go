package models

import (
	"github.com/JohnGeorge47/ollert/pkg/sql"
	"time"
)

type Account struct {
	AccountName string
	CreatedAt   string
}

var Table = "Accounts"

func CreateAccount(acc Account) (*int64, error) {
	acc.CreatedAt = time.Now().Format(time.RFC3339)
	query := "INSERT INTO Accounts(account_name) values(?)"
	c := sql.Connmanager
	lastid, err := c.Insert(query, acc.AccountName)
	if err != nil {
		return nil, err
	}
	return lastid, err
}
