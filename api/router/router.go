package router

import (
	"cloudstorageapi.com/api/controllers/filecontroller"
	"cloudstorageapi.com/api/controllers/homecontroller"
	"cloudstorageapi.com/api/middlewares"
	"cloudstorageapi.com/configs"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var Router = httprouter.New()

func init() {
	Router.GET("/", homecontroller.GetIndex)
	//file routers
	Router.POST("/files", middlewares.AuthenticateUser(filecontroller.UploadFile))
	Router.DELETE("/files", middlewares.AuthenticateUser(filecontroller.DeleteFile))
	Router.ServeFiles("/downloads/*filepath", http.Dir(configs.STORAGE_ROOT_PATH))
}
