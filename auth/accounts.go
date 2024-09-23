package auth

import (
	"errors"
	"fmt"
	"net/mail"
	"github.com/a3ylf/flowerly/database"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {

    psw, err := bcrypt.GenerateFromPassword([]byte (password),10)
    if err != nil {
        return password,err
    }
    return string(psw), nil
}

func CheckPassword(password ,otherpassword []byte) error {
    err := bcrypt.CompareHashAndPassword(password,otherpassword)
    if err != nil {
        return err
    }
    return nil


}

func EnsureSignup(c *database.Client ) error {
  _, err := mail.ParseAddress(c.Email)

    if err != nil {
    return errors.New("Invalid Email")
  }

  if len(c.CPF) != 11 {
    return errors.New("Invalid CPF")
  }
  return nil
}
