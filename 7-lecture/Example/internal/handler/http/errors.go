package handler

import (
	"errors"
	"net/http"

	sfssError "gitlab.com/lgp/http-server-example/internal/errors"
)

// errorMap is map with errors for http handle func.
type errorMap map[error]int

var (
	errorsGetMap = errorMap{
		sfssError.ErrFileNotExists: http.StatusNotFound,
		sfssError.ErrInternalError: http.StatusInternalServerError,
	}

	errorsUploadMap = errorMap{
		sfssError.ErrFileAlreadyExists: http.StatusConflict,
		sfssError.ErrInternalError:     http.StatusInternalServerError,
	}

	errorsDeleteMap = errorMap{
		sfssError.ErrFileNotExists: http.StatusNotFound,
		sfssError.ErrInternalError: http.StatusInternalServerError,
	}
)

// getMappedStatusCode return code with errorMap and err.
// If err is nil returns http.StatusOK.
// If err can't be find in errorMap returns http.StatusInternalServerError.
func getMappedStatusCode(m errorMap, err error) int {
	if err == nil {
		return http.StatusOK
	}
	for e, c := range m {
		if errors.Is(err, e) {
			return c
		}
	}
	return http.StatusInternalServerError
}
