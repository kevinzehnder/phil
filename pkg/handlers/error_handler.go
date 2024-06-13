package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kevinzehnder/phil/pkg/errs"

	"github.com/rs/zerolog/log"
)

func errorHandler(err error, w http.ResponseWriter) {
	restErr := &errs.RestErr{}

	switch {
	// if error is a RestErr
	case errors.As(err, &restErr):
		// log error
		log.Debug().
			Str("InternalError", err.Error()).
			Msg("RestError")

		// http response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(restErr.Status)
		json.NewEncoder(w).Encode(restErr)

	default:
		// log error
		log.Error().
			Str("InternalError", err.Error()).
			Msg("InternalServerError")

		// http response
		p := errs.NewInternalServerError(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(p.Status)
		json.NewEncoder(w).Encode(p)
	}
}
