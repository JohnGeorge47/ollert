package accountCreation

import (
	"encoding/json"
	"github.com/JohnGeorge47/ollert/internal/models"
	"io/ioutil"
	"net/http"
	"time"
)

type Signup struct {
	AccountName string `json:"account_name"`
	Admin       bool   `json:"admin"`
	EmailId     string `json:"email_id"`
	UserName    string `json:"user_name"`
}

const createdFormat = "2006-01-02 15:04:05"

func Create() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			contentType := r.Header.Get("Content-type")
			if contentType != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				err := Operate(r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Only post method allowed"))
		}
	}
	return http.HandlerFunc(fn)
}

func Operate(r *http.Request) error {
	var reqBody Signup
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &reqBody)
	if err != nil {
		return err
	}
	acc := models.Account{
		AccountName: reqBody.AccountName,
		CreatedAt:   "",
	}
	acc_id, err := models.CreateAccount(acc)
	if err != nil {
		return err
	}
	user := models.User{
		EmailId:    reqBody.EmailId,
		UserName:   reqBody.UserName,
		AccountId:  *acc_id,
		IsAdmin:    reqBody.Admin,
		Created_At: time.Unix(1391878657, 0).Format(createdFormat),
	}
	models.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
