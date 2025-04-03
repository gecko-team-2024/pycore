package controllers

import (
	"net/http"
	"pycore/services"
)

// UploadFileHandler xử lý request upload file
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Nhận file từ form-data
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Không thể đọc file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Khởi tạo Google Drive Service
	driveService, err := services.NewGoogleDriveService("credentials.json")
	if err != nil {
		http.Error(w, "Lỗi kết nối Google Drive", http.StatusInternalServerError)
		return
	}

	// Upload file lên Google Drive
	driveFileID, err := driveService.UploadFileToDrive("fileName", file, "FOLDER_ID")
	if err != nil {
		http.Error(w, "Upload file thất bại", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Upload thành công", "drive_file_id": "` + driveFileID + `"}`))
}
