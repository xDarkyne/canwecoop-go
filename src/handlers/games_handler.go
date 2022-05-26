package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
	"github.com/xdarkyne/steamgo/router"
	"github.com/xdarkyne/steamgo/sync"
)

func GamesHandler() http.Handler {
	r := router.NewRouter()

	r.Get("/", getGamesHandler)
	r.Post("/", postGamesHandler)

	return r
}

func GetNextCursor(c paginator.Cursor) *string {
	if c.After != nil {
		return nil
	}

	return c.After
}

// METHOD: GET
func getGamesHandler(w http.ResponseWriter, r *http.Request) {
	cursorParam := r.URL.Query().Get("cursor")
	limit := 24

	var allGames []models.Game

	stmt := db.ORM.
		Preload("Categories").
		Preload("Genres")

	p := paginator.New(&paginator.Config{
		Limit: limit,
	})

	if len(cursorParam) != 0 {
		p.SetAfterCursor(cursorParam)
	}

	result, cursor, err := p.Paginate(stmt, &allGames)
	// paginator error
	if err != nil {
		panic(err.Error())
	}
	// gorm error
	// https://gorm.io/docs/error_handling.html
	if result.Error != nil {
		panic(result.Error.Error())
	}

	response := struct {
		Games      []models.Game
		NextCursor *string
	}{
		Games:      allGames,
		NextCursor: GetNextCursor(cursor),
	}

	json.NewEncoder(w).Encode(response)
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
