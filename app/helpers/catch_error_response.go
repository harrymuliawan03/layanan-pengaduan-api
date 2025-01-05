package helpers

import (
	"net/http"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/schemas"
)

func CatchErrorResponseApi(r *schemas.ResponseApiError) *schemas.SetResponseApiError {
	switch r.Status {
	case schemas.ApiErrorBadRequest:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusBadRequest,
			Message:    r.Message,
		}
	case schemas.ApiErrorForbidden:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusForbidden,
			Message:    r.Message,
		}
	case schemas.ApiErrorNotFound:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusNotFound,
			Message:    r.Message,
		}
	case schemas.ApiErrorUnprocessAble:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    r.Message,
		}
	default:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    r.Message,
		}
	}
}
