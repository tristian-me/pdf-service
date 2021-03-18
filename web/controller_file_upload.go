package web

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"pdf-service/gen"
	"pdf-service/utils"
)

// HandleHTMLUpload handles uploading files
func HandleHTMLUpload(w http.ResponseWriter, r *http.Request) {
	var err error

	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.RespBadJSON(w, http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	err = checkContentType(handler)
	if err != nil {
		utils.RespBadJSON(w, http.StatusBadRequest, err)
		return
	}

	f, err := os.OpenFile(TempDir+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.RespBadJSON(w, http.StatusConflict, err)
		return
	}

	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		utils.RespBadJSON(w, http.StatusConflict, err)
		return
	}

	// The new file
	newFilename := utils.RandomString(30) + ".html"

	// Rename the file
	src := TempDir + "/" + handler.Filename
	dst := TempDir + "/" + newFilename
	err = os.Rename(src, dst)
	if err != nil {
		utils.RespBadJSON(w, http.StatusConflict, err)
		return
	}

	newFile, err := gen.ConvertFromFile(newFilename)
	if err != nil {
		utils.RespBadJSON(w, http.StatusConflict, err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+newFile)
	w.Header().Set("Content-Type", "application/octect-stream")
	http.ServeFile(w, r, TempDir+"/"+newFile)
}

func checkContentType(handler *multipart.FileHeader) error {
	contentType := handler.Header.Get("Content-Type")
	if contentType != "text/html" {
		return errors.New("Invalid content-type")
	}

	return nil
}
