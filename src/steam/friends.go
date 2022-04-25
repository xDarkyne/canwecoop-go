package steam

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xdarkyne/steamgo/config"
)

// Relationship is a type of relationship
type Relationship string

const (
	// All is a type of relationship
	All Relationship = "all"
	// Friend is a type of relationship
	Friend Relationship = "friend"
)

// SteamFriend is a relationship between two steam users
type SteamFriend struct {
	SteamID      uint64 `json:",string"`
	Relationship Relationship
	FriendSince  int64 `json:"friend_since"`
}

type playerFriendsListJSON struct {
	Friendslist *struct {
		Friends []SteamFriend
	}
}

func GetFriendsList(userID string) (playerFriendsListJSON, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend", config.App.SteamAPIKey, userID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var friends playerFriendsListJSON
	err = json.NewDecoder(resp.Body).Decode(&friends)
	if err != nil {
		fmt.Println(err)
		return friends, err
	}
	return friends, nil
}
