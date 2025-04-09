# Cau truc du an (vi du):
pycore/
│── main.go                     # Điểm khởi chạy chính
│── go.mod                      # Quản lý module Go
│── go.sum                      # Checksum của các dependencies
│── index.html                  # Giao diện HTML tĩnh (nếu có)
│── documents.md                # Tài liệu ghi chú
│── README.md                   # Hướng dẫn sử dụng
│── .env                        # File chứa biến môi trường
│── .gitignore                  # Bỏ qua các file không cần commit
│── .firebaserc                 # Cấu hình Firebase CLI
│── firebase.json               # Cấu hình Firebase
│── credentials.json            # Thông tin xác thực Firebase
│── serviceAccountKey.json      # Khóa dịch vụ Firebase
│── server.log                  # Log server
│── storage.rules               # Quy tắc truy cập storage Firebase
│
├── config/
│   ├── cors_config.go          # Cấu hình CORS
│   ├── firebase.go             # Cấu hình kết nối Firebase
│   ├── oauth.go                # Cấu hình OAuth
│
├── controllers/
│   ├── auth_controller.go      # Xử lý request xác thực
│   ├── file_controller.go      # Xử lý upload/download file
│   ├── log_controller.go       # Xử lý liên quan đến log
│   ├── user_controller.go      # Xử lý request người dùng
│
├── middleware/
│   ├── auth_middleware.go      # Middleware xác thực JWT
│   ├── logger_middleware.go    # Middleware ghi log
│
├── models/
│   ├── oauth.go                # Định nghĩa model OAuth
│   ├── user.go                 # Định nghĩa model User
│
├── routes/
│   ├── routes.go               # Khai báo các route của API
│
├── services/
│   ├── file_service.go         # Xử lý logic file
│   ├── log_service.go          # Xử lý logic log
│   ├── user_service.go         # Xử lý logic người dùng
│
└── data/
    └── pyfarm.zip              # File dữ liệu đính kèm

# Thu vien su dung:
# Thư viện cho routing:
go get github.com/gorilla/mux

# Quản lý biến môi trường:
go get github.com/joho/godotenv

# Xác thực và bảo mật:
go get github.com/golang-jwt/jwt/v5
go get github.com/rs/cors

# Firebase và Google Cloud:
go get firebase.google.com/go
go get google.golang.org/api

# UUID:
go get github.com/google/uuid

# OAuth2 và mã hóa:
go get golang.org/x/crypto
go get golang.org/x/oauth2