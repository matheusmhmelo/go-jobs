package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/matheusmhmelo/go-jobs/internal/repository"
	"os"
	"time"
)

type User struct {
	Id 		int64  `json:"id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
	Pass	string `json:"password,omitempty"`
	Phone	string `json:"phone"`
	State 	string `json:"state"`
	City	string `json:"city"`
	Address string `json:"address"`
}

func (u *User) Validate() error {
	eTxt := "validation error, valid field: %v"

	if u.Name == "" {
		err := errors.New(fmt.Sprintf(eTxt, "name"))
		return err
	}

	if u.Email == "" {
		err := errors.New(fmt.Sprintf(eTxt, "email"))
		return err
	}

	return nil
}

func (u *User) Register(password string) (map[string]interface{}, error) {
	r := repository.NewUser()

	id, err := r.Create(password, u.Name, u.Email, u.Phone, u.State, u.City, u.Address)
	if err != nil {
		return nil, err
	}

	u.Id = id

	token, err := createJWT(id)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"user": u,
		"message": "User created successful!",
		"token": token,
	}

	return resp, nil
}

func createJWT(id int64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}