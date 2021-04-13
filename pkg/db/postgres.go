package db

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgreSQL struct {
	db *sql.DB
}

func (p *PostgreSQL) Connect(c Config) {
	connectionLine := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	connectionLine = fmt.Sprintf(connectionLine,
		c.GetHost(), c.GetPort(), c.GetUser(), c.GetPassword(), c.GetDatabase(), c.GetSSL())
	drive, err := sql.Open("postgres", connectionLine)
	if err != nil {
		panic(err)
	}
	p.db = drive
}

func (p *PostgreSQL) Execute(query string, args ...interface{}) (sql.Result, error) {
	var result sql.Result

	stmtIns, err := p.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer stmtIns.Close()

	result, err = stmtIns.Exec(args...)
	if err != nil {
		log.Println(err)
		return result, err
	}

	return result, nil
}

func (p *PostgreSQL) Get(query string, args ...interface{}) (map[string]interface{}, error) {
	stmt, err := p.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	entry := make(map[string]interface{})
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry = make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
	}
	return entry, nil
}

func (p *PostgreSQL) Fetch(query string, args ...interface{}) ([]map[string]interface{}, error) {
	stmt, err := p.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}