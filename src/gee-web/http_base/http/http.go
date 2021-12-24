package http

import "net/http"

type Handler interface {
	ServeHttp(w http.ResponseWriter, r *http.Request)
}

//func ListenAndServe(address string, h Handler) error
