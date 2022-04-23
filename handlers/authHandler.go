package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xdarkyne/steamgo/config"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet: 
			getAuthHandler(w, r)
		case http.MethodPost:
			fmt.Println("Post")
		case http.MethodDelete:
			deleteAuthHandler(w, r)
		case http.MethodPut:
			fmt.Println("Put")
		}
}

func getAuthHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteAuthHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie(config.App.AuthCookieName)
	if err != nil {
		fmt.Println(err)
		return;
	}

	expire := time.Now().AddDate(0, 0, 1)
	cookie := &http.Cookie{Name: config.App.AuthCookieName, Value: authCookie.Value, Expires: expire, MaxAge: -1, HttpOnly: true}
	http.SetCookie(w, cookie)

	fmt.Println(authCookie.Value)
}
