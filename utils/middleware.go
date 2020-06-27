package utils

import (
	"net/http"
	"time"
)

func LogRequestMetaData(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		LogInfo.Printf("%s %s request received at %s", r.RequestURI, r.Method, time.Now().Format(time.RFC3339))
		next.ServeHTTP(w,r)
	})
}
