package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pycore/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Cho phép mọi nguồn gốc (CORS)
	},
}

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

func StreamLogsHandler(w http.ResponseWriter, r *http.Request) {
	// Nâng cấp kết nối HTTP lên WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Chỉ ghi lỗi nếu kết nối không được nâng cấp
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket request received from:", r.RemoteAddr)
	fmt.Println("Headers:", r.Header)

	// Đường dẫn file log
	logFilePath := "server.log"

	// Gửi log qua WebSocket
	err = services.StreamLogs(logFilePath, conn)
	if err != nil {
		// Gửi lỗi qua WebSocket thay vì sử dụng http.ResponseWriter
		conn.WriteMessage(websocket.TextMessage, []byte("Error streaming logs: "+err.Error()))
		return
	}
}
