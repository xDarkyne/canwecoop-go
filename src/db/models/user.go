package models

import "time"

type User struct {
	ID           string `gorm:"primaryKey"`
	DisplayName  string `gorm:"unique"`
	AvatarUrl    string
	ProfileUrl   string
	IsTester     bool
	IsAdmin      bool
	CreatedAt    time.Time
	LastLoggedIn time.Time
}
