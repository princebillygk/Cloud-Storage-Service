package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Db struct {
	Name     string
	User     string
	Host     string
	Password string
}

func (db *Db) getConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable", db.Name, db.User, db.Password, db.Host)
}

func (db *Db) Connect() (*sql.DB, error) {
	return sql.Open("postgres", db.getConnectionString())
}
