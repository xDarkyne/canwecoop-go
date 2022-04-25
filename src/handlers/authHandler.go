package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAuthHandler(w, r)
	case http.MethodDelete:
		deleteAuthHandler(w, r)
	case http.MethodOptions:
		OptionMethod(w, "GET, DELETE, OPTIONS")
	default:
		MethodNotAllowedError(w, "GET, DELETE, OPTIONS")
	}
}

type RError struct {
	Error bool
}

// METHOD: GET
func getAuthHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie(config.App.AuthCookieName)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		rerror := RError{Error: true}
		json.NewEncoder(w).Encode(rerror)
		return
	}

	var user models.User
	db.ORM.First(&user, authCookie.Value)
	json.NewEncoder(w).Encode(user)

}

// METHOD: DELETE
func deleteAuthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(config.App.AuthCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: config.App.AuthCookieName, MaxAge: -1, Path: "/"})
	json.NewEncoder(w).Encode(RError{Error: false})
}
