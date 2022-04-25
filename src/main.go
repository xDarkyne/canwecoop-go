package main

import (
	"encoding/json"
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

	router.HandleFunc("/api/login", handlers.LoginHandler)
	router.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie(config.App.AuthCookieName)
			if err != nil {
				fmt.Println(err)
				return
			}
		
			http.SetCookie(w, &http.Cookie{Name: config.App.AuthCookieName, MaxAge: -1})
			json.NewEncoder(w).Encode("eerr")
	})
	router.HandleFunc("/api/auth", handlers.AuthHandler)
	router.HandleFunc("/api/games", handlers.GamesHandler)
	router.Handle("/", handlers.FileHandler("public/"))

	addr := fmt.Sprintf(":%d", config.App.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
