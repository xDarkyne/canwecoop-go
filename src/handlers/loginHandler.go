package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/solovev/steam_go"
	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(config.App.AuthCookieName)
	if cookie != nil {
		http.Redirect(w, r, "/", 302)
		return
	}

	login(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	opId := steam_go.NewOpenId(r)
	switch opId.Mode() {
	case "":
		http.Redirect(w, r, opId.AuthUrl(), 301)
	case "cancel":
		w.Write([]byte("Authorization cancelled"))
	default:
		steamUser, err := opId.ValidateAndGetUser(config.App.SteamAPIKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var user models.User
		result := db.ORM.First(&user, steamUser.SteamId)
		if result.Error != nil {
			user = models.User{
				ID: steamUser.SteamId,
				DisplayName: steamUser.PersonaName,
				AvatarUrl: steamUser.AvatarFull,
				IsTester: false,
				IsAdmin: false,
				CreatedAt: time.Now(),
				LastLoggedIn: time.Now(),
			}
			db.ORM.Create(&user)
		}
		result = db.ORM.Model(&models.User{}).Where("id = ?", user.ID).Update("LastLoggedIn", time.Now())
		if result.Error != nil {
			fmt.Println(result.Error)
		}

		expire := time.Now().AddDate(0, 0, 1)
		cookie := &http.Cookie{Name: config.App.AuthCookieName, Value: steamUser.SteamId, Expires: expire, MaxAge: 86400, HttpOnly: true, Path: "/"}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", 302)
	}
}
