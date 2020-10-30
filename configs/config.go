package configs

import (
	"cloudstorageapi.com/database"
	"context"
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"path/filepath"
)

//fileSystemConfigs
const storageFolderName = "cloud-store-files-here"
const MAX_UPLOAD_SIZE_IN_BYTE = 64000000 //64MB

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

func GetRedisConnection() (redis.Conn, error) {
	ctx := context.Background()
	return redis.DialContext(ctx, "tcp", ":6379")
}
