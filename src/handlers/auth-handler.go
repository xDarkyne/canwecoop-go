package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
)

func AuthHandler() http.Handler {
	r := NewDarkRouter()

	r.Get("/", getAuthHandler)
	r.Delete("/", deleteAuthHandler)
	r.OptionsHandler("GET, DELETE")
	r.MethodNotAllowedHandler("GET, DELETE")

	return r
}

// METHOD: GET
func getAuthHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie(config.App.AuthCookieName)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "User is not authenticated", http.StatusUnauthorized)
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
		http.Error(w, "User is not authenticated", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: config.App.AuthCookieName, MaxAge: -1, Path: "/"})
	w.Write([]byte("Logged out successfully!"))
}
