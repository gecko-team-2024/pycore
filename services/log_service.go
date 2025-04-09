package services

import (
	"bufio"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func ReadLogs(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func StreamLogs(filePath string, conn *websocket.Conn) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Đọc file từ cuối (tail)
	file.Seek(0, os.SEEK_END)
	reader := bufio.NewReader(file)

	for {
		// Đọc từng dòng mới từ file
		line, err := reader.ReadString('\n')
		if err != nil {
			// Nếu không có dòng mới, chờ 1 giây và thử lại
			time.Sleep(1 * time.Second)
			continue
		}

		// Gửi dòng log qua WebSocket
		err = conn.WriteMessage(websocket.TextMessage, []byte(line))
		if err != nil {
			// Nếu kết nối WebSocket bị đóng, thoát khỏi vòng lặp
			return err
		}
	}
}
