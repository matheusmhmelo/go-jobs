package user

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/go-jobs/internal/repository"
)

type Login struct {
	Email 	string `json:"email"`
	Pass	string `json:"password,omitempty"`
}

func (l *Login) Validate() error {
	eTxt := "validation error, valid field: %v"

	if l.Email == "" {
		err := errors.New(fmt.Sprintf(eTxt, "name"))
		return err
	}

	if l.Pass == "" {
		err := errors.New(fmt.Sprintf(eTxt, "email"))
		return err
	}

	return nil
}

func (l *Login) Do() (map[string]interface{}, error) {
	r := repository.NewUser()

	user, err := r.Login(l.Email, l.Pass)
	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		err = errors.New("user not found")
		return nil, err
	}

	token, err := createJWT(user["id"].(int64))
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"user": user,
		"message": "User logged in!",
		"token": token,
	}

	return resp, nil
}
