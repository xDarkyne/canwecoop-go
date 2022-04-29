package sync

import (
	"fmt"
	"net/http"
	"time"

	"github.com/peppage/kettle"
	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db"
	"github.com/xdarkyne/steamgo/db/models"
)

/*
 * The code in this file is extremely ugly and has to
 * be refactored soon
 */

func SyncGames() {
	httpClient := http.DefaultClient
	steamClient := kettle.NewClient(httpClient, config.App.SteamAPIKey)

	var users []models.User
	result := db.ORM.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	list := map[int64]int64{}

	for _, v := range users {
		games, _, err := steamClient.IPlayerService.GetOwnedGames(&kettle.OwnedGamesParams{SteamID: v.ID, IncludeAppInfo: 0, IncludeFree: 1})
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v2 := range games {
			gameID := v2.AppID
			if _, ok := list[gameID]; !ok {
				list[gameID] = gameID
			}
		}
	}

	count := 0
	for _, g := range list {
		count++
		details, _, err := steamClient.Store.AppDetails(g)
		if err != nil {
			fmt.Println(err)
		}
		id := fmt.Sprint(details.SteamAppID)
		result := db.ORM.Create(&models.Game{
			ID:                 fmt.Sprint(id),
			Name:               details.Name,
			IsFree:             details.IsFree,
			ShortDescription:   details.ShortDescription,
			HeaderImageUrl:     details.HeaderImage,
			Website:            details.Website,
			BackgroundImageUrl: details.Background,
			StoreUrl:           fmt.Sprintf("https://store.steampowered.com/app/%d", details.SteamAppID),
			IsHidden:           false,
		})
		if result.Error != nil {
			fmt.Println(result.Error)
			continue
		}
		var game models.Game
		db.ORM.First(&game, "id = ?", id)

		for _, c := range details.Categories {
			id := fmt.Sprint(c.ID)
			err := db.ORM.Model(&game).Association("Categories").Append(&models.Category{
				ID:        id,
				Name:      c.Description,
				Relevance: 0,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		for _, g := range details.Genres {
			id := fmt.Sprintf(g.ID)
			err := db.ORM.Model(&game).Association("Genres").Append(&models.Genre{
				ID:        id,
				Name:      g.Description,
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
