package job

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/go-jobs/internal/repository"
	"time"
)

type Job struct {
	Id          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
}

func (j *Job) Validate() error {
	eTxt := "validation error, valid field: %v"

	if j.Title == "" {
		err := errors.New(fmt.Sprintf(eTxt, "title"))
		return err
	}

	if j.Description == "" {
		err := errors.New(fmt.Sprintf(eTxt, "description"))
		return err
	}

	return nil
}

func (j *Job) ValidateUpdate() error {
	err := j.Validate()
	if err != nil {
		return err
	}

	r := repository.NewJob()

	job, err := r.GetInfo(int(j.Id))
	if err != nil {
		return err
	}

	if len(job) == 0 {
		err = errors.New("job not found")
		return err
	}

	if job["user_id"].(int64) != j.UserID {
		err = errors.New("wrong user")
		return err
	}

	return nil
}

func (j *Job) Create() (map[string]interface{}, error) {
	r := repository.NewJob()

	date := timeNowInSP()
	j.Created = date.Format(time.RFC3339)

	id, err := r.Create(j.UserID, j.Title, j.Description, j.Created)
	if err != nil {
		return nil, err
	}

	j.Id = id

	resp := map[string]interface{}{
		"job": j,
		"message": "Job created successful!",
	}

	return resp, nil
}

func (j *Job) Update() (map[string]interface{}, error) {
	r := repository.NewJob()

	err := r.Update(j.Title, j.Description, j.Id)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"job": j,
		"message": "Job updated successful!",
	}

	return resp, nil
}

func GetInfo(id int) (map[string]interface{}, error) {
	r := repository.NewJob()

	j, err := r.GetInfo(id)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"job": j,
	}

	return resp, nil
}

func Delete(id int) (map[string]interface{}, error) {
	r := repository.NewJob()

	err := r.Delete(id)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"message": "Job deleted successful!",
	}

	return resp, nil
}

func timeNowInSP() time.Time {
	date := time.Now()

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return date
	}

	return date.In(loc)
}