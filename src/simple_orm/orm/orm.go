package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

func Connect(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Minute)
	return conn, conn.Ping()
}

type Dba interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(string) (*sql.Stmt, error)
}

type Query struct {
	DB     Dba
	table  string
	wheres []string
	only   []string
	limit  string
	offset string
	order  string
	errs   []string
}

