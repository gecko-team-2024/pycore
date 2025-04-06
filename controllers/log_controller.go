package controllers

import (
	"encoding/json"
	"net/http"
	"pycore/services"
)

func GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	logFilePath := "server.log"

	logs, err := services.ReadLogs(logFilePath)
	if err != nil {
		http.Error(w, "Error reading logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"logs": logs,
	})
}
