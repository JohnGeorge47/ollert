package handlers

import (
	"encoding/json"
	"errors"
	"github.com/JohnGeorge47/ollert/internal/core/password"
	"io/ioutil"
	"net/http"
)

type Password struct {
	UserEmail *string `json:"user_email"`
	Password  *string `json:"password"`
}

func PassWordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		contentType := r.Header.Get("Content-type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pwdbody, err := CreateCheck(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = password.Create(*pwdbody.UserEmail, *pwdbody.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case "GET":
	}
}

func CreateCheck(body []byte) (*Password, error) {
	var reqBody Password
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		return nil, err
	}
	if reqBody.Password == nil {
		return nil, errors.New("Missing password field")
	}
	if reqBody.UserEmail == nil {
		return nil, errors.New("Missing email field")
	}
	return &reqBody, nil
}
