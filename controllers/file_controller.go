package controllers

import (
	"net/http"
	"pycore/services"
)

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "pyfarm.zip"

	filePath, err := services.GetFilePath(fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}
