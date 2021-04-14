package repository

import (
	"fmt"
	"github.com/matheusmhmelo/go-jobs/pkg/db"
	"log"
	"os"
	"strconv"
)

type Search struct {
	db        db.Database
}

func NewSearch() *Search {
	return &Search{
		db: Db,
	}
}

func (s *Search) AllJobs(page int) ([]map[string]interface{}, int64, error) {
	limit, offset, err := getLimitAndOffset(page)
	if err != nil {
		return nil, 0, err
	}

	q := "SELECT " +
			"	j.id, " +
			"	j.title as job_title, " +
			"	j.description as job_description, " +
			"	j.created as job_created, " +
			"	j.user_id, " +
			"	u.name as user_name, " +
			"	u.email as user_email, " +
			"	u.phone as user_phone, " +
			"	u.state as user_state, " +
			"	u.city as user_city, " +
			"	u.address as user_address" +
			"	FROM jobs j " +
			"	INNER JOIN users u " +
			"	ON j.user_id = u.id " +
			"	ORDER BY j.id ASC " +
			"	LIMIT " + limit +
			"	OFFSET " + offset

	r, err := s.db.Fetch(q)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	pQuery := fmt.Sprintf(
		"SELECT " +
			"	count(j.id) as total " +
			"	FROM jobs j " +
			"	INNER JOIN users u " +
			"	ON j.user_id = u.id")

	pRes, err := s.db.Get(pQuery)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	total := pRes["total"].(int64)

	return r, total, nil
}

func (s *Search) Jobs(title string, page int) ([]map[string]interface{}, int64, error) {
	limit, offset, err := getLimitAndOffset(page)
	if err != nil {
		return nil, 0, err
	}

	q := "SELECT " +
			"	j.id, " +
			"	j.title as job_title, " +
			"	j.description as job_description, " +
			"	j.created as job_created, " +
			"	j.user_id, " +
			"	u.name as user_name, " +
			"	u.email as user_email, " +
			"	u.phone as user_phone, " +
			"	u.state as user_state, " +
			"	u.city as user_city, " +
			"	u.address as user_address" +
			"	FROM jobs j " +
			"	INNER JOIN users u " +
			"	ON j.user_id = u.id " +
			"	WHERE LOWER(j.title) LIKE LOWER('%" + title + "%')" +
			"	ORDER BY j.id ASC " +
			"	LIMIT " + limit +
			"	OFFSET " + offset

	r, err := s.db.Fetch(q)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	pQuery := "SELECT " +
			"	count(j.id) as total " +
			"	FROM jobs j " +
			"	INNER JOIN users u " +
			"	ON j.user_id = u.id " +
		"	WHERE LOWER(j.title) LIKE LOWER('%" + title + "%')"

	pRes, err := s.db.Get(pQuery)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	total := pRes["total"].(int64)

	return r, total, err
}

func getLimitAndOffset(page int) (string, string, error) {
	limit := os.Getenv("RESULTS_PER_PAGE")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "", "", err
	}

	offset := (page - 1) * limitInt

	return limit, strconv.Itoa(offset), nil
}