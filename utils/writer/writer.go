package writer

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mailru/easyjson"

	myerrors "github.com/masharpik/TransactionalSystem/utils/errors"
	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

func WriteErrorMessageRespond(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	responseErr := myerrors.New(message)

	w.WriteHeader(statusCode)
	started, _, err := easyjson.MarshalToHTTPResponseWriter(responseErr, w)

	if !started {
		errorMsg := fmt.Errorf(literals.LogErrorOccurredBeforeResponseWriterMethods, err)
		logger.LogRequestError(r, http.StatusInternalServerError, errorMsg)
		return
	} else if err != nil {
		logger.LogRequestError(r, http.StatusInternalServerError, err)
		return
	}

	logger.LogRequestError(r, statusCode, errors.New(message))
}

func WriteErrorJSONRespond(w http.ResponseWriter, r *http.Request, statusCode int, responseJSON easyjson.Marshaler, err error) {
	w.WriteHeader(statusCode)
	started, _, err := easyjson.MarshalToHTTPResponseWriter(responseJSON, w)

	if !started {
		errorMsg := fmt.Errorf(literals.LogErrorOccurredBeforeResponseWriterMethods, err)
		logger.LogRequestError(r, http.StatusInternalServerError, errorMsg)
		return
	} else if err != nil {
		logger.LogRequestError(r, http.StatusInternalServerError, err)
		return
	}

	logger.LogRequestError(r, statusCode, err)
}

func WriteSuccessJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, responseJSON easyjson.Marshaler) {
	w.WriteHeader(statusCode)
	started, _, err := easyjson.MarshalToHTTPResponseWriter(responseJSON, w)
	if !logger.DebugOn {
		return
	}

	if !started {
		errorMsg := fmt.Errorf(literals.LogErrorOccurredBeforeResponseWriterMethods, err)
		logger.LogRequestError(r, http.StatusInternalServerError, errorMsg)
		return
	} else if err != nil {
		logger.LogRequestError(r, http.StatusInternalServerError, err)
		return
	}

	logger.LogRequestSuccess(r, statusCode)
}
