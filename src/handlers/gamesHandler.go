package handlers

import (
	"fmt"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/steam"
)

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getGamesHandler(w, r)
	case http.MethodOptions:
		OptionMethod(w, "GET, OPTIONS")
	default:
		MethodNotAllowedError(w, "GET, OPTIONS")
	}
}

func getGamesHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(config.App.AuthCookieName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "No cookie found", http.StatusUnauthorized)
	}
	games, err := steam.GetSteamGames(c.Value)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, v := range games.Response.Games {
		fmt.Println(v.ID)
	}
	w.Write([]byte("Success"))
}
