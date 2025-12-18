package middleware

import (
	"fmt"
	"net/http"
	"time"
)

const (
	green  = "\033[32m"
	white  = "\033[37m"
	yellow = "\033[33m"
	red    = "\033[31m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		statusColor := colorByStatus(rw.status)
		methodColor := colorByMethod(r.Method)

		fmt.Printf(
			"%s[API]%s %s | %s%3d%s | %8s | %s%-6s%s %s\n",
			blue, reset,
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusColor, rw.status, reset,
			duration,
			methodColor, r.Method, reset,
			r.URL.Path,
		)
	})
}

func colorByStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorByMethod(method string) string {
	switch method {
	case "GET":
		return green
	case "POST":
		return blue
	case "PUT":
		return yellow
	case "DELETE":
		return red
	default:
		return white
	}
}
