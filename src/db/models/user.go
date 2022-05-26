package models

import "time"

type User struct {
	ID           string `gorm:"primaryKey"`
	DisplayName  string `gorm:"unique"`
	AvatarUrl    string
	ProfileUrl   string
	Friends      []*User `gorm:"many2many:user_user"`
	Games        []Game  `gorm:"many2many:game_user"`
	IsTester     bool
	IsAdmin      bool
	CreatedAt    time.Time
	LastLoggedIn time.Time
}
