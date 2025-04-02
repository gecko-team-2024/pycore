package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logFile, _ = os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
var logger = log.New(logFile, "", log.LstdFlags)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		logger.Printf("[%s] %s %s - %s", r.Method, r.RequestURI, r.RemoteAddr, duration)
	})
}

func LogEvent(event string) {
	logger.Println(event)
}
