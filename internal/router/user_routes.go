package router

import (
	"net/http"
)

type userHandler interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}
type Route struct {
	handler userHandler
}

func NewRoute(h userHandler) *Route {
	return &Route{handler: h}
}
func (ro *Route) RegisterRoute() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", ro.handler.Index)
	router.HandleFunc("GET /{id}", ro.handler.Show)
	router.HandleFunc("POST /", ro.handler.Store)
	router.HandleFunc("PUT /{id}", ro.handler.Update)
	router.HandleFunc("DELETE /{id}", ro.handler.Destroy)
	return router
}
