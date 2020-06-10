package password

import (
	"fmt"
	"github.com/JohnGeorge47/ollert/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const createdFormat = "2006-01-02 15:04:05"

func Create(email string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
		return err
	}
	passobj := models.UserPassword{
		UserEmail:  email,
		Password:   string(hash),
		Updated_at: time.Unix(1391878657, 0).Format(createdFormat),
	}
	models.PasswordCreate(passobj)
	return nil
}

func Login(email string, password string) error {
	hashpwd, err := models.GetPassword(email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(*hashpwd), []byte(password))
	if err != nil {
		return err
	}

}
