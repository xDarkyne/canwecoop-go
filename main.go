package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/handlers"
)

func main() {
	config.LoadConfig()
	db.Connect()

	migrate := flag.Bool("migrate", false, "Check the migration request")
	flag.Parse()

	if *migrate {
		db.Migrate()
	}

	router := http.NewServeMux()

	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/auth", handlers.AuthHandler)
	router.HandleFunc("/games", handlers.GamesHandler)

	addr := fmt.Sprintf(":%d", config.App.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
