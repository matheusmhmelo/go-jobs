package handler

import (
	"encoding/json"
	"net/http"
)

//Heartbeat only to check the health of the API
func Heartbeat(w http.ResponseWriter, _ *http.Request) {
	ret, _ := json.Marshal(map[string]string{
		"message": "I'm Online!",
	})
	_, _ = w.Write(ret)
}
