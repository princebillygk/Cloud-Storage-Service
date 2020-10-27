package config

import (
	"cloudstorageapi.com/database"
	"database/sql"
	"log"
)

const dbname = "cloudstorageapi"
const dbuser = "cloudstorageapi"
const dbpassword = ""
const dbhost = "localhost"

var Connection *sql.DB

func init() {
	log.Println("Establishing database connection")
	db := database.Db{Name: dbname, User: dbuser, Password: dbpassword, Host: dbhost}
	var err error
	Connection, err = db.Connect()
	if err != nil {

		log.Fatal(err)
		return
	}
	log.Println("Established database connection")
}
