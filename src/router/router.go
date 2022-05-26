package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type DarkRouter struct {
	*chi.Mux
	allowedMethods string
}

func NewRouter() *DarkRouter {
	dr := &DarkRouter{
		chi.NewRouter(),
		"OPTIONS",
	}

	dr.Options("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", dr.allowedMethods)
		w.WriteHeader(http.StatusNoContent)
	})

	dr.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", dr.allowedMethods)
		http.Error(w, fmt.Sprintf("Method %s not allowed.", r.Method), http.StatusMethodNotAllowed)
	})

	return dr
}

func (dr *DarkRouter) allowMethod(method string) {
	if !strings.Contains(dr.allowedMethods, method) {
		methods := strings.Split(dr.allowedMethods, ", ")
		methods = append(methods, method)
		dr.allowedMethods = strings.Join(methods, ", ")
	}
}

/*** HTTP METHODS ***/

/* GET */
func (dr *DarkRouter) Get(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("GET")

	dr.Mux.Get(pattern, fn)
}

/* HEAD */
func (dr *DarkRouter) Head(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("HEAD")

	dr.Mux.Head(pattern, fn)
}

/* POST */
func (dr *DarkRouter) Post(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("POST")

	dr.Mux.Post(pattern, fn)
}

/* PUT */
func (dr *DarkRouter) Put(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("PUT")

	dr.Mux.Put(pattern, fn)
}

/* DELETE */
func (dr *DarkRouter) Delete(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("DELETE")

	dr.Mux.Delete(pattern, fn)
}

/*
 * OPTIONS METHOD NOT SUPPORTED YET
 */

/* CONNECT */
func (dr *DarkRouter) Connect(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("CONNECT")

	dr.Mux.Connect(pattern, fn)
}

/* TRACE */
func (dr *DarkRouter) Trace(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("TRACE")

	dr.Mux.Trace(pattern, fn)
}

/* PATCH */
func (dr *DarkRouter) Patch(pattern string, fn http.HandlerFunc) {
	dr.allowMethod("PATCH")

	dr.Mux.Patch(pattern, fn)
}
