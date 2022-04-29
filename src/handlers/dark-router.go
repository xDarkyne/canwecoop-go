package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
 * "extemsion" of the chi router to "automate" some
 * boring tasks like implementing the options method.
 */
type DarkRouter struct {
	*chi.Mux
}

func (dr DarkRouter) OptionsHandler(allowedMethods string) {
	dr.Options("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", allowedMethods)
		w.WriteHeader(http.StatusNoContent)
	})
}

func (dr DarkRouter) MethodNotAllowedHandler(allowedMethods string) {
	dr.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", allowedMethods)
		http.Error(w, fmt.Sprintf("Method %s not allowed.", r.Method), http.StatusMethodNotAllowed)
	})
}

func NewDarkRouter() *DarkRouter {
	r := &DarkRouter{chi.NewRouter()}
	return r
}
