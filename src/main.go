package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/handlers"
	"github.com/xdarkyne/steamgo/sync"
)

func main() {
	config.LoadConfig()
	db.Connect()

	migrate := flag.Bool("migrate", false, "Check the migration request")
	syncFlag := flag.Bool("sync", false, "Check the migration request")
	flag.Parse()

	if *migrate {
		db.Migrate()
		return
	}

	if *syncFlag {
		sync.SyncGames()
		return
	}

	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Mount("/login", handlers.LoginHandler())
		r.Mount("/auth", handlers.AuthHandler())
		r.Mount("/games", handlers.GamesHandler())
	})
	router.Handle("/*", handlers.FileHandler(handlers.HTMLDir{D: http.Dir("public/")}))

	addr := fmt.Sprintf(":%d", config.App.Port)
	fmt.Println("Running on port ", config.App.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
