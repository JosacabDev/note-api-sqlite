package middleware

import (
	"bytes"
	"github/JosacabDev/api-sqlite/pkg/logger"
	"io"
	"net/http"
	"time"
)

type responseCacher struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rc *responseCacher) WriteHeader(statusCode int) {
	rc.statusCode = statusCode
	rc.ResponseWriter.WriteHeader(statusCode)
}

func (rc *responseCacher) Write(b []byte) (int, error) {
	rc.body.Write(b)
	return rc.ResponseWriter.Write(b)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		debug := true

		var reqBody []byte
		if debug && r.Body != nil {
			reqBody, _ = io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		rc := &responseCacher{
			ResponseWriter: w,
			body:           &bytes.Buffer{},
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rc, r)

		duration := time.Since(startTime)

		logger.Info.Printf("[%s] %s | %d | %s | %s",
			r.Method,
			r.RequestURI,
			rc.statusCode,
			r.RemoteAddr,
			duration)

		if debug {
			if r.Method == http.MethodPost || r.Method == http.MethodPut {
				logger.Info.Printf("Request Body: %s", reqBody)
			}

			logger.Info.Printf("Response Body: %s", rc.body.String())
		}

		if rc.statusCode >= 500 {
			logger.Error.Printf("Server Error: [%s] %d | %s ",
				r.Method,
				rc.statusCode,
				r.RequestURI)
		}
	})
}
