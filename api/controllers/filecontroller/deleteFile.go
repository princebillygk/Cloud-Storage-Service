package filecontroller

import (
	"cloudstorageapi.com/api/helper"
	"cloudstorageapi.com/api/response"
	"cloudstorageapi.com/models/spaceModel"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type DeleteFileRequestData struct {
	FilePath string `json:"path"`
}

type DeleteFileResponseData struct {
	FilePath string `json:"path"`
}

func DeleteFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//geting file path from url parameter
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		helper.NewJsonResponse(w).SetStatus(http.StatusBadRequest).
			Encode(response.NewErrorResponse(
				991, "Request with invalid path"))
		return
	}

	//This will success at any cost because we have checked in middleware so I am skiping it
	spaceId, _ := strconv.Atoi(r.Header.Get("Space-ID"))
	space := spaceModel.Space{Id: spaceId}
	spacePath := space.GetFilePath()

	//Attempt to remove the file
	rerr := os.RemoveAll(filepath.Join(spacePath, filePath))
	if rerr != nil {
		helper.NewJsonResponse(w).SetStatus(http.StatusInternalServerError).
			Encode(response.NewFailResponse("No such file or directory"))
		return
	}

	//success response
	helper.NewJsonResponse(w).SetStatus(http.StatusOK).
		Encode(response.NewSuccessResponse(
			DeleteFileResponseData{filePath},
			"Deleted successfully",
		))
}
