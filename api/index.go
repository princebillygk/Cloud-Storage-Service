package main

import (
	"cloudstorageapi.com/api/router"
	"log"
	"net/http"
)

func main() {
	/*/
	log.Fatal(http.ListenAndServeTLS(":8000", "localhost.crt", "localhost.key", router.Router)) //with ssl http request
	//*/
	log.Fatal(http.ListenAndServe(":8000", router.Router)) //without ssl http request
	//*/
}
