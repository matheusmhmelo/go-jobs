package handler

import (
	"encoding/json"
	"net/http"
)

func commonHttpError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	ret, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	_, _ = w.Write(ret)
}
