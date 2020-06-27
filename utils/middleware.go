package utils

import (
	"github.com/google/uuid"
	"net/http"
)

func LogRequestMetaData(next http.Handler) http.Handler{
	transactionID := uuid.New()
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		LogInfo.Printf("[%s]: %s %s request received", transactionID, r.RequestURI, r.Method)
		next.ServeHTTP(w,r)
	})
}
