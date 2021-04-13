package request

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var ContextKey = os.Getenv("CONTEXT_KEY")

type AccessDetails struct {
	UserId	int64
}

type controller func(http.ResponseWriter, *http.Request)

func PreRequest(h controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := TokenValid(r); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		access, err := ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey, access.UserId)

		w.Header().Set("Content-Type", "application/json")
		h(w, r.WithContext(ctx))
	}
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			UserId:   int64(userId),
		}, nil
	}
	return nil, err
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}