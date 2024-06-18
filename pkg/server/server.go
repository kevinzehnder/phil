package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kevinzehnder/phil/pkg/app"
	"github.com/kevinzehnder/phil/pkg/server/middlewares"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/kevinzehnder/phil/docs"
)

type Server struct {
	router   *chi.Mux
	handlers []app.Handler[any]
	srv      http.Server
}

func NewServer(handlers ...app.Handler[any]) *Server {
	router := chi.NewRouter()
	AddGlobalMiddlewares(router)
	AddSwagger(router)
	AddRoutes(router, handlers...)
	return &Server{router: router, handlers: handlers}
}

func (s *Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	s.srv = http.Server{Addr: address, Handler: s.router}
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.TODO())
}

func AddRoutes(router chi.Router, handlers ...app.Handler[any]) {
	// create a sub-router
	apiRouter := chi.NewRouter()

	// add API middlewares
	// apiRouter.Use(middlewares.JwtAuthMiddleware)

	// add routes from handlers
	for _, handler := range handlers {
		for path, methods := range handler.Routes() {
			for method, fn := range methods {
				apiRouter.MethodFunc(method, path, fn)
			}
		}
	}
	router.Mount("/", apiRouter)
}

func AddGlobalMiddlewares(router chi.Router) {
	router.Use(middlewares.ZeroLogGPTLogger)
}

func AddSwagger(router chi.Router) {
	router.Mount("/swagger", httpSwagger.WrapHandler)
}
