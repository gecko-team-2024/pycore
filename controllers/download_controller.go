package controllers

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"pycore/config"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func LoadFileEnv() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No ENV file found")
		}
	}
}

func DownloadFolderHandler(w http.ResponseWriter, r *http.Request) {
	// Lấy token từ header Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return
	}

	// Loại bỏ tiền tố "Bearer "
	token := authHeader[len("Bearer "):]

	// Tạo client Google Drive
	client := config.GoogleOAuthConfig.Client(context.Background(), &oauth2.Token{
		AccessToken: token,
	})

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		http.Error(w, "Unable to create Drive client", http.StatusInternalServerError)
		return
	}

	// ID của thư mục Google Drive
	folderID := "1KMpYbynIp8wFrSpmsziPNuYtdEYOYAHb"

	// Liệt kê file trong thư mục
	query := fmt.Sprintf("'%s' in parents and trashed=false", folderID)
	fileList, err := srv.Files.List().Q(query).Fields("files(id, name)").Do()
	if err != nil {
		http.Error(w, "Unable to retrieve files", http.StatusInternalServerError)
		return
	}

	// Tạo file zip để lưu trữ các file
	zipFileName := "folder.zip"
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		http.Error(w, "Unable to create zip file", http.StatusInternalServerError)
		return
	}
	defer func() {
		zipFile.Close()
		os.Remove(zipFileName) // Xóa file zip sau khi gửi về client
	}()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Tải từng file và thêm vào file zip
	for _, file := range fileList.Files {
		downloadURL := fmt.Sprintf("https://www.googleapis.com/drive/v3/files/%s?alt=media", file.Id)
		resp, err := client.Get(downloadURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to download file: %s", file.Name), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Tạo entry trong file zip
		zipEntry, err := zipWriter.Create(file.Name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to create zip entry for file: %s", file.Name), http.StatusInternalServerError)
			return
		}

		// Ghi nội dung file vào zip
		_, err = io.Copy(zipEntry, resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write file to zip: %s", file.Name), http.StatusInternalServerError)
			return
		}
	}

	// Trả file zip về client
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", zipFileName))
	w.Header().Set("Content-Type", "application/zip")
	http.ServeFile(w, r, zipFileName)
}
