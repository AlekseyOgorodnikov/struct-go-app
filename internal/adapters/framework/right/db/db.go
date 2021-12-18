package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect db
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("DB connecting is failur: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) AddHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").
		Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).
		ToSql()
	if err != nil {
		return err
	}

	result, err := da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	log.Printf("Result insert row in db: %v", result)

	return nil
}
