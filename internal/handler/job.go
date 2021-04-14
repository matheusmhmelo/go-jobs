package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/matheusmhmelo/go-jobs/internal/service/job"
	"net/http"
	"os"
	"strconv"
)

//CreateJob create a new Job for User
func CreateJob(w http.ResponseWriter, r *http.Request) {
	var j job.Job
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		commonHttpError(w, err)
		return
	}

	err = j.Validate()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ctx := r.Context()
	id := ctx.Value(os.Getenv("CONTEXT_KEY"))

	j.UserID = id.(int64)

	resp, err := j.Create()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ret, _ := json.Marshal(resp)
	_, _ = w.Write(ret)
}

//UpdateJob update a existent Job
func UpdateJob(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var j job.Job
	err = json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ctx := r.Context()
	userID := ctx.Value(os.Getenv("CONTEXT_KEY"))

	j.UserID = userID.(int64)
	j.Id = int64(id)

	err = j.ValidateUpdate()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	resp, err := j.Update()
	if err != nil {
		commonHttpError(w, err)
		return
	}

	ret, _ := json.Marshal(resp)
	_, _ = w.Write(ret)
}

//GetJobInfo return info of Job
func GetJobInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := job.GetInfo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ret, _ := json.Marshal(result)
	_, _ = w.Write(ret)
}

//RemoveJob delete Job from database
func RemoveJob(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := job.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ret, _ := json.Marshal(result)
	_, _ = w.Write(ret)
}