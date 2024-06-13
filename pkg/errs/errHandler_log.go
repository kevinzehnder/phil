package errs

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type LogErrorHandler struct{}

func NewLogErrorHandler() *LogErrorHandler {
	return &LogErrorHandler{}
}

func (l *LogErrorHandler) Handle(err error) {
	var restErr *RestErr

	switch {
	// if error is a RestErr
	case errors.As(err, &restErr):
		log.Debug().
			Msgf("RestError: %s", err.Error())

	default:
		log.Error().
			Msgf("InternalServerError: %v", err.Error())
	}
}
