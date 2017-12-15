package resource

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler func(http.ResponseWriter, *http.Request)

type Model struct {
	Index  handler
	Show   handler
	Create handler
	Update handler
	Delete handler
}

// FIXME: Remove mux.Route reference and use Route interface
type Router interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) *mux.Route
}

type Route interface {
	Methods(string) Route
}

func Generate(path string, res Model, r Router) {
	pathWithId := path + "/{id:[0-9]+}"
	generateRoutes(path, pathWithId, res, r)
}

func GenerateWithStringId(path string, res Model, r Router) {
	pathWithId := path + "/{id}"
	generateRoutes(path, pathWithId, res, r)
}

func generateRoutes(path string, pathWithId string, res Model, r Router) {
	if res.Index != nil {
		r.HandleFunc(path, res.Index).Methods("GET")
	}
	if res.Show != nil {
		r.HandleFunc(pathWithId, res.Show).Methods("GET")
	}
	if res.Create != nil {
		r.HandleFunc(path, res.Create).Methods("POST")
	}
	if res.Update != nil {
		r.HandleFunc(pathWithId, res.Update).Methods("PUT")
	}
	if res.Delete != nil {
		r.HandleFunc(pathWithId, res.Delete).Methods("DELETE")
	}
}
