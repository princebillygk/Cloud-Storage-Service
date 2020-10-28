package configs

import (
	"cloudstorageapi.com/database"
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

//fileSystemConfigs
const storageFolderName = "cloud-store-files-here"

//db-configs
const dbname = "cloudstorageapi"
const dbuser = "cloudstorageapi"
const dbpassword = ""
const dbhost = "localhost"

var STORAGE_ROOT_PATH string
var Connection *sql.DB

func init() {
	//storage configuration
	configureStorage()
	//database configuration
	configureDatabase()

}

func configureStorage() {
	rootdir, err := os.UserHomeDir()
	if err != nil {
		rootdir = "."
	}
	STORAGE_ROOT_PATH = filepath.Join(rootdir, storageFolderName)
}

func configureDatabase() {
	db := database.Db{Name: dbname, User: dbuser, Password: dbpassword, Host: dbhost}
	var err error
	Connection, err = db.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
}
