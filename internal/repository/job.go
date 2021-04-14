package repository

import (
	"fmt"
	"github.com/matheusmhmelo/go-jobs/pkg/db"
	"log"
	"strconv"
)

type Job struct {
	db        db.Database
}

func NewJob() *Job {
	return &Job{
		db: Db,
	}
}

func (j *Job) Create(user int64, title, description, date string) (int64, error) {
	q := fmt.Sprintf("INSERT INTO jobs (user_id, title, description, created) VALUES (%v, '%v', '%v', '%v') RETURNING id",
		strconv.Itoa(int(user)), title, description, date)

	r, err := j.db.Get(q)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id := r["id"].(int64)

	return id, nil
}

func (j *Job) Update(title, description string, id int64) error {
	q := fmt.Sprintf("UPDATE jobs SET title = '%v', description = '%v' WHERE id = %v",
		title, description, strconv.Itoa(int(id)))

	_, err := j.db.Execute(q)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (j *Job) GetInfo(id int) (map[string]interface{}, error) {
	q := fmt.Sprintf("SELECT id, title, description, created, user_id FROM jobs WHERE id = %v", strconv.Itoa(id))

	r, err := j.db.Get(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return r, nil
}

func (j *Job) Delete(id int) error {
	q := fmt.Sprintf("DELETE FROM jobs WHERE id = %v", strconv.Itoa(id))

	_, err := j.db.Execute(q)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}