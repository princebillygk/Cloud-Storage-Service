package filecontroller

import (
	"cloudstorageapi.com/api/helper"
	"cloudstorageapi.com/api/response"
	"cloudstorageapi.com/configs"
	"cloudstorageapi.com/models/spaceModel"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type UploadFileResponseData struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	SizeInPath   int64  `json:"sizeInBytes"`
	FilePath     string `json:"filePath"`
	DownloadLink string `json:"downloadURL"`
}

func UploadFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fileName := r.FormValue("name")
	filePath := r.FormValue("path")

	//fileuploadlimit
	r.ParseMultipartForm(configs.MAX_UPLOAD_SIZE_IN_BYTE)

	//Attempt to get file & and file informations
	file, fileHeader, ferr := r.FormFile("file")
	if ferr != nil {
		helper.NewJsonResponse(w).SetStatus(http.StatusNoContent).
			Encode(response.NewErrorResponse(992, "Invalid or No file passed"))
		return
	}
	defer file.Close()

	//getting file extension
	fileExtension := filepath.Ext(fileHeader.Filename)

	//This will never fail cause we checked space id in middleware
	spaceIdStr := r.Header.Get("Space-ID")
	spaceId, _ := strconv.Atoi(spaceIdStr)
	space := spaceModel.Space{Id: spaceId}
	spacePath := space.GetFilePath()

	fileSavingPath := filepath.Join(spacePath, filePath)
	newFileName := fileName + fileExtension

	//creating neccesary path before writing gonna ignore any error for now
	_ = os.MkdirAll(fileSavingPath, os.ModePerm)

	//saving the uploaded file to our storage
	f, foerr := os.OpenFile(filepath.Join(fileSavingPath, newFileName), os.O_WRONLY|os.O_CREATE, 0666)
	if foerr != nil {
		log.Fatal(foerr)
		helper.NewJsonResponse(w).SetStatus(http.StatusInternalServerError).
			Encode(response.NewFailResponse("Failed to save file"))
		return
	}
	defer f.Close()
	io.Copy(f, file)

	newFileAddress := filepath.Join(filePath, newFileName)
	downloadURL := "http://" + filepath.Join(r.Host, "downloads", spaceIdStr, newFileAddress)

	//success response
	helper.NewJsonResponse(w).SetStatus(http.StatusCreated).
		Encode(response.NewSuccessResponse(
			UploadFileResponseData{
				Name:         fileName,
				Path:         filePath,
				SizeInPath:   fileHeader.Size,
				FilePath:     newFileAddress,
				DownloadLink: downloadURL},
			"File Uploaded Successfully"))
}
