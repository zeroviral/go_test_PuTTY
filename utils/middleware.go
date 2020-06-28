package utils

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func LogRequestMetaData(next http.Handler) http.Handler {
	txID := uuid.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "transactionID", txID)
		r = r.WithContext(ctx)
		LogInfo.Printf("%s %s request received transactionID:[%s]", r.RequestURI, r.Method, txID)
		next.ServeHTTP(w, r)
	})
}
