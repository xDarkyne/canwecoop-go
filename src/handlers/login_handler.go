package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/solovev/steam_go"
	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
	"github.com/xdarkyne/steamgo/router"
)

func LoginHandler() http.Handler {
	r := router.NewRouter()

	r.Get("/", getLoginHandler)

	return r
}

// METHOD: GET
func getLoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(config.App.AuthCookieName)
	if cookie != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	/* Set the origin cookie here if the
	 * Referer header is set. The cookie
	 * is used to redirect the user to
	 * the page he came from
	 */
	ref := r.Referer()
	if ref != "" {
		http.SetCookie(w, &http.Cookie{
			Name:  "cwc-origin",
			Value: ref,
			Path:  "/",
		})
	}

	opId := steam_go.NewOpenId(r)
	switch opId.Mode() {
	case "":
		http.Redirect(w, r, opId.AuthUrl(), http.StatusMovedPermanently)
	case "cancel":
		w.Write([]byte("Authorization cancelled"))
	default:

		/* THIS CODE IS ONLY REACHED ON SUCCESSFULL AUTHENTICATION */
		steamUser, err := opId.ValidateAndGetUser(config.App.SteamAPIKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var user models.User
		result := db.ORM.First(&user, steamUser.SteamId)
		if result.Error != nil {
			user = models.User{
				ID:           steamUser.SteamId,
				DisplayName:  steamUser.PersonaName,
				AvatarUrl:    steamUser.AvatarFull,
				ProfileUrl:   steamUser.ProfileUrl,
				IsTester:     false,
				IsAdmin:      false,
				Friends:      []*models.User{},
				Games:        []models.Game{},
				CreatedAt:    time.Now(),
				LastLoggedIn: time.Now(),
			}
			db.ORM.Create(&user)
		}

		// update Username/Avatar in the case they changed
		if steamUser.AvatarFull != user.AvatarUrl || steamUser.PersonaName != user.DisplayName {
			result = db.ORM.Model(&user).Updates(models.User{AvatarUrl: steamUser.AvatarFull, DisplayName: steamUser.PersonaName})
			if result.Error != nil {
				fmt.Println(result.Error)
			}
		}

		// update last login date-time
		result = db.ORM.Model(&user).Updates(models.User{LastLoggedIn: time.Now()})
		if result.Error != nil {
			fmt.Println(result.Error)
		}

		// update friendslist
		friends, err := config.App.SteamAPI.GetFriendsList(steamUser.SteamId)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(friends); i++ {
			var friend models.User
			result := db.ORM.First(&friend, friends[i].UserID)
			if result.Error != nil {
				continue
			}
			err := db.ORM.Model(&user).Association("Friends").Append(&models.User{ID: friend.ID})
			if err != nil {
				fmt.Println(err)
			}
		}

		destination := "/"
		c, _ := r.Cookie("cwc-origin")
		if c != nil {
			destination = c.Value
			// remove the origin cookie
			http.SetCookie(w, &http.Cookie{
				Name:   "cwc-origin",
				Path:   "/",
				MaxAge: -1,
			})
		}

		expire := time.Now().AddDate(0, 0, 1)
		cookie := &http.Cookie{Name: config.App.AuthCookieName, Value: steamUser.SteamId, Expires: expire, MaxAge: 86400, HttpOnly: true, Path: "/"}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, destination, http.StatusMovedPermanently)
	}
}
