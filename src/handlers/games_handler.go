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

// METHOD: GET
func getGamesHandler(w http.ResponseWriter, r *http.Request) {
	cursor := r.URL.Query().Get("cursor")
	limit := 24

	var allGames []models.Game

	stmt := db.ORM.
		Preload("Categories").
		Preload("Genres")

	p := paginator.New(&paginator.Config{
		Limit: limit,
	})

	if len(cursor) != 0 {
		p.SetAfterCursor(cursor)
	}

	result, pCursor, err := p.Paginate(stmt, &allGames)
	// paginator error
	if err != nil {
		panic(err.Error())
	}
	// gorm error
	// https://gorm.io/docs/error_handling.html
	if result.Error != nil {
		panic(result.Error.Error())
	}

	type response struct {
		Games      []models.Game
		NextCursor string
	}

	aRes := response{
		Games:      allGames,
		NextCursor: "",
	}

	fmt.Println(*pCursor.After)

	if len(*pCursor.After) != 0 {
		aRes.NextCursor = *pCursor.After
	}

	json.NewEncoder(w).Encode(aRes)
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