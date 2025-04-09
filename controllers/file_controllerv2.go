package controllers

import (
	"fmt"
	"net/http"
	"pycore/middleware"
	"pycore/services"
)

func DownloadFileHandlerV2(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fileName := "pyfarm.zip"
	filePath, err := services.GetFilePath(fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	fmt.Printf("User %s is downloading %s\n", username, fileName)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}
