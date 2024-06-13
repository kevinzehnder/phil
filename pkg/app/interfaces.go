package app

import "net/http"

type Servicer[T any, I any] interface {
	GetAll() (*[]T, error)
	GetOne(id I) (*T, error)
	Create(*T) error
	Update(*T) error
	DeleteOne(id I) error
}

type Databaser[T any, I any] interface {
	GetAll() (*[]T, error)
	GetOne(*T) error
	Save(*T) (*I, error)
	Update(*T) error
	Delete(*T) error
	Shutdown() error
}

type Handler[T any] interface {
	Routes() map[string]map[string]func(http.ResponseWriter, *http.Request)
}

type ErrorHandler interface {
	Handle(err error)
}
