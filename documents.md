# Cau truc du an (vi du):
golang-server/
│── main.go                 # Điểm khởi chạy chính
│── go.mod                  # File quản lý module
│── go.sum                  # File chứa checksum dependencies
│── serviceAccountKey.json
├── config/
│   ├── firebase.go         # Cấu hình kết nối Firestore
│   ├── env.go              # Load biến môi trường
│
├── controllers/
│   ├── user_controller.go  # Xử lý request liên quan đến user
│   ├── post_controller.go  # Xử lý request liên quan đến bài viết
│
├── models/
│   ├── user.go             # Định nghĩa model User
│   ├── post.go             # Định nghĩa model Post
│
├── routes/
│   ├── routes.go           # Định nghĩa các route API
│
├── services/
│   ├── user_service.go     # Xử lý logic liên quan đến User
│   ├── post_service.go     # Xử lý logic liên quan đến Post
│
└── middleware/
    ├── auth_middleware.go  # Middleware kiểm tra JWT

# Thu vien su dung:
go get github.com/gin-gonic/gin
go get firebase.google.com/go
go get google.golang.org/api/option
go get github.com/joho/godotenv

# Chay bang lenh
```
    go run main.go
```

# Test login with Google
http://localhost:8080/auth/google
https://pycore.onrender.com/auth/google