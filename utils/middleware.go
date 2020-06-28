package utils

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func LogRequestMetaData(next http.Handler) http.Handler {
	txID := uuid.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "txID", txID)
		r = r.WithContext(ctx)
		LogInfo.Printf("%s %s request received txID:[%s]", r.RequestURI, r.Method, txID)
		next.ServeHTTP(w, r)
	})
}
