package handler

import (
	"encoding/json"
	"github.com/matheusmhmelo/go-jobs/internal/service/user"
	"net/http"
	"os"
	"strconv"
)

//Register create a new User on system
func Register(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		commonHttpError(w, err)
		return
	}

	pass := u.Pass
	u.Pass = ""

	err = u.Validate()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	resp, err := u.Register(pass)
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ret, _ := json.Marshal(resp)
	_, _ = w.Write(ret)
}

//Login authenticate a User on system
func Login(w http.ResponseWriter, r *http.Request) {
	var l user.Login
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		commonHttpError(w, err)
		return
	}

	err = l.Validate()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	resp, err := l.Do()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ret, _ := json.Marshal(resp)
	_, _ = w.Write(ret)
}

//ValidSession valid if session is OK
func ValidSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := ctx.Value(os.Getenv("CONTEXT_KEY"))
	ret, _ := json.Marshal(map[string]string{
		"message": "User ID: " + strconv.Itoa(int(id.(int64))),
	})
	_, _ = w.Write(ret)
}