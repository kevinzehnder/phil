package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kevinzehnder/phil/internal/config"
	"github.com/kevinzehnder/phil/pkg/app"
	"github.com/kevinzehnder/phil/pkg/errs"
	"github.com/kevinzehnder/phil/pkg/services"

	"github.com/gorilla/schema"
)

type WhoamiHandler struct {
	Svc *services.WhoamiSvc
}

func NewWhoamiHandler(whoamiSvc *services.WhoamiSvc) *WhoamiHandler {
	whoamiHandler := WhoamiHandler{Svc: whoamiSvc}
	return &whoamiHandler
}

func (h *WhoamiHandler) Routes() map[string]map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]map[string]func(http.ResponseWriter, *http.Request){
		"/": {
			http.MethodGet: h.Ping,
		},
	}
}

func (h *WhoamiHandler) Ping(w http.ResponseWriter, r *http.Request) {
	// parse query parameters from request
	queryValues := r.URL.Query()

	// unmarshal query parameters into struct
	queryParams := app.QueryParam{}
	decoder := schema.NewDecoder()
	err := decoder.Decode(&queryParams, queryValues)
	if err != nil {
		restErr := errs.NewBadRequestError("bad query parameters")
		restErr.InternalError = err.Error()
		err = fmt.Errorf("decoder.Decode | %w", restErr)
		errorHandler(err, w)
		return
	}

	// call whoami Ping service
	err = h.Svc.Ping()
	if err != nil {
		err = fmt.Errorf("whoami.svc.Ping | %w", err)
		errorHandler(err, w)
		return
	}

	// backgroundColor := "#3dfaf0"
	settings, _ := config.GetConfig()
	version := settings.Version

	hostname, _ := os.Hostname()

	data := app.Data{
		Version:    version,
		Hostname:   hostname,
		IP:         h.Svc.GetIPs(),
		Headers:    r.Header,
		URL:        r.URL.RequestURI(),
		Host:       r.Host,
		Method:     r.Method,
		Name:       "Whoami",
		RemoteAddr: r.RemoteAddr,
	}

	// Write a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
