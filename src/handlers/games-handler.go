package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
	"github.com/xdarkyne/steamgo/sync"
)

func GamesHandler() http.Handler {
	r := NewDarkRouter()

	r.Get("/", getGamesHandler)
	r.Post("/", postGamesHandler)
	r.OptionsHandler("GET, POST")
	r.MethodNotAllowedHandler("GET, POST")

	return r
}

// METHOD: GET
func getGamesHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(config.App.AuthCookieName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "No cookie found", http.StatusUnauthorized)
	}
	var games []models.Game
	result := db.ORM.Find(&games, "is_hidden = ?", false)
	if result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&games)
}

// METHOD: POST
func postGamesHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(config.App.AuthCookieName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "No cookie found", http.StatusUnauthorized)
	}
	sync.SyncGames()
	w.Write([]byte(c.Value))
}
