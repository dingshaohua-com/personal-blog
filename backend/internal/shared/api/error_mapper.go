package api

import (
	"backend/internal/shared/validation"
	"errors"
	"log"
	"net/http"
)

type ErrorMapping struct {
	Target error
	Status int
}

// MapError 领域错误 → HTTP 状态码
func MapError(
	internalMessage string,
	err error,
	mappings ...ErrorMapping,
) error {
	if err == nil {
		return nil
	}
	if validationErr, ok := errors.AsType[*validation.Error](err); ok {
		return NewError(http.StatusUnprocessableEntity, validationErr.Message)
	}
	for _, mapping := range mappings {
		if errors.Is(err, mapping.Target) {
			return NewError(
				mapping.Status,
				err.Error(),
			)
		}
	}

	log.Printf("%s: %v", internalMessage, err)
	return InternalError()
}
