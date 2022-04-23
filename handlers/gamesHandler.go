package handlers

import (
	"fmt"
	"net/http"
)

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet: 
			getGamesHandler(w, r)
		case http.MethodPost:
			fmt.Println("Post")
		case http.MethodDelete:
			fmt.Println("Delete")
		case http.MethodPut:
			fmt.Println("Put")
		}
}

func getGamesHandler(w http.ResponseWriter, r *http.Request) {

}
