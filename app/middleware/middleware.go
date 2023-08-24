package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
	"github.com/masharpik/TransactionalSystem/utils/writer"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			logger.LogOperationSuccess(fmt.Sprintf("Function %s %s execution took %s\n", r.Method, r.URL.Path, time.Since(start).String()))
		}()

		defer func() {
			if err := recover(); err != nil {
				var actualErr error
				if e, ok := err.(error); ok {
					actualErr = e
				} else {
					actualErr = fmt.Errorf("%v", err)
				}
				writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, fmt.Errorf(literals.LogPanicOccured, actualErr).Error())
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
