package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// GoogleDriveService struct để quản lý kết nối Drive API
type GoogleDriveService struct {
	Service *drive.Service
}

// NewGoogleDriveService khởi tạo kết nối Google Drive API
func NewGoogleDriveService(credentialsFile string) (*GoogleDriveService, error) {
	ctx := context.Background()

	// Đọc credentials JSON
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("không thể đọc file credentials: %v", err)
	}

	// Xác thực với Google Drive API
	config, err := google.JWTConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("lỗi xác thực Google Drive: %v", err)
	}

	client := config.Client(ctx)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("không thể tạo Drive service: %v", err)
	}

	return &GoogleDriveService{Service: srv}, nil
}

// UploadFileToDrive tải file lên Google Drive
func (g *GoogleDriveService) UploadFileToDrive(filename string, file io.Reader, folderID string) (string, error) {
	// Tạo metadata file
	fileMetadata := &drive.File{
		Name:    filename,
		Parents: []string{folderID}, // Lưu vào thư mục cụ thể (hoặc bỏ nếu không cần)
	}

	// Upload file
	driveFile, err := g.Service.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("upload file thất bại: %v", err)
	}

	log.Printf("Tải file lên thành công: %s", driveFile.Id)
	return driveFile.Id, nil
}
