package api

import (
	"errors"
	"log"
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
	for _, mapping := range mappings {
		if errors.Is(err, mapping.Target) {
			return NewError(
				mapping.Status,
				err.Error(),
			)
		}
	}

	log.Printf("operation failed: %v", err)
	return InternalError(internalMessage + err.Error())
}
