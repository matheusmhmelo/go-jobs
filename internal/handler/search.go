package handler

import (
	"encoding/json"
	"github.com/matheusmhmelo/go-jobs/internal/service/search"
	"net/http"
)

//SearchJobs search for Jobs by the title
func SearchJobs(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	page := r.URL.Query().Get("page")

	resp, err := search.Jobs(title, page)
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ret, _ := json.Marshal(resp)
	_, _ = w.Write(ret)
}
