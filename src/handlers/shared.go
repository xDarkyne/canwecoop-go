package handlers

import "net/http"

func MethodNotAllowedError(w http.ResponseWriter, allowedMethods string) {
	w.Header().Set("Allow", allowedMethods)
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func OptionMethod(w http.ResponseWriter, allowedMethods string) {
	w.Header().Set("Allow", allowedMethods)
	w.WriteHeader(http.StatusNoContent)
}
