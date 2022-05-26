package sync

import (
	"fmt"
	"time"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
)

/*
 * The code in this file is extremely ugly and has to
 * be refactored soon
 */

func SyncGames() {
	var users []models.User
	result := db.ORM.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	list := map[string]string{}

	for _, v := range users {
		games, err := config.App.SteamAPI.GetOwnedGames(v.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v2 := range games {
			gameID := string(v2.GameID)
			if _, ok := list[gameID]; !ok {
				list[gameID] = gameID
			}
		}
	}

	count := 0
	for _, g := range list {
		count++
		details, err := config.App.SteamAPI.GetGameDetails(g)
		if err != nil {
			fmt.Println(err)
		}
		id := string(details.Id)
		result := db.ORM.Create(&models.Game{
			ID:                 fmt.Sprint(id),
			Name:               details.Name,
			IsFree:             details.IsFree,
			ShortDescription:   details.ShortDescription,
			HeaderImageUrl:     details.HeaderImageUrl,
			Website:            details.Website,
			BackgroundImageUrl: details.BackgroundImageUrl,
			StoreUrl:           fmt.Sprintf("https://store.steampowered.com/app/%s", id),
			IsHidden:           false,
		})
		if result.Error != nil {
			fmt.Println(result.Error)
			continue
		}
		var game models.Game
		db.ORM.First(&game, "id = ?", id)

		for _, c := range details.Categories {
			id := fmt.Sprint(c.Id)
			err := db.ORM.Model(&game).Association("Categories").Append(&models.Category{
				ID:        id,
				Name:      c.Name,
				Relevance: 0,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		for _, g := range details.Genres {
			id := fmt.Sprint(g.Id)
			err := db.ORM.Model(&game).Association("Genres").Append(&models.Genre{
				ID:        id,
				Name:      g.Name,
				Relevance: 0,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		if (count % 195) == 0 {
			fmt.Println("Reached 195 Games, waiting for about 5 minutes before continuing.")
			time.Sleep(5 * time.Minute)
		}
	}
}
