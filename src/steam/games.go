package steam

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
)

type Game struct {
	ID              int `json:"appid"`
	PlaytimeTotal   int `json:"playtime_forever"`
	PlaytimeWindows int `json:"playtime_windows_forever"`
	PlaytimeMac     int `json:"playtime_mac_forever"`
	PlaytimeLinux   int `json:"playtime_linux_forever"`
}

type playerGamesListJSON struct {
	Response *struct {
		Gamecount int
		Games     []Game
	}
}

func GetSteamGames(userID string) (playerGamesListJSON, error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", config.App.SteamAPIKey, userID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var gamesList playerGamesListJSON
	err = json.NewDecoder(resp.Body).Decode(&gamesList)
	if err != nil {
		fmt.Println(err)
		return gamesList, err
	}
	return gamesList, nil
}
