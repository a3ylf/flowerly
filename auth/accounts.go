package auth

import "golang.org/x/crypto/bcrypt"

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
