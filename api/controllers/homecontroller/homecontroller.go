package homecontroller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Hmm nothing here dude :'( but don't delete it
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) { //HomeController
	fmt.Fprintln(w, "Cloud storage api")
}
