package utils

import (
	"net/http"
)

func LogRequestMetaData(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		LogInfo.Printf("%s %s request received", r.RequestURI, r.Method)
		next.ServeHTTP(w,r)
	})
}
