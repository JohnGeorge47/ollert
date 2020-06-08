package accountCreation

import (
	"encoding/json"
	"github.com/JohnGeorge47/ollert/internal/models"
	"io/ioutil"
	"net/http"
)

type Signup struct {
	AccountName string `json:"account_name"`
	Admin       bool   `json:"admin"`
	EmailId     string `json:"email_id"`
	UserName    string `json:"user_name"`
}

func Create() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			contentType := r.Header.Get("Content-type")
			if contentType != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				Operate(r)
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
	models.CreateAccount(acc)
	return nil
}
