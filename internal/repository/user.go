package repository

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/go-jobs/pkg/db"
	"log"
)

type User struct {
	db        db.Database
}

func NewUser() *User {
	return &User{
		db: Db,
	}
}

func (u *User) Create(password, name, email, phone, state, city, address string) (int64, error) {
	if password == "" {
		err := errors.New("password can not be empty")
		return 0, err
	}

	q := fmt.Sprintf("INSERT INTO users (name, email, password, phone, state, city, address) VALUES ('%v', '%v', MD5('%v'), '%v', '%v', '%v', '%v') RETURNING id",
		name, email, password, phone, state, city, address)

	r, err := u.db.Get(q)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id := r["id"].(int64)

	return id, nil
}

func (u *User) Login(email, password string) (map[string]interface{}, error) {
	q := fmt.Sprintf("SELECT id, name, email, phone, state, city, address FROM users WHERE email = '%v' AND password = MD5('%v')",
		email, password)

	r, err := u.db.Get(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return r, nil
}
