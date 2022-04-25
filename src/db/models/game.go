package models

type Game struct {
	ID                 string `gorm:"primaryKey"`
	Name               string
	IsFree             bool
	ShortDescription   string
	HeaderImageUrl     string
	Website            string
	Categories         []Category `gorm:"many2many:game_category"`
	Genres             []Genre    `gorm:"many2many:game_genre"`
	BackgroundImageUrl string
	StoreUrl           string
	IsHidden           bool
}

type Category struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Relevance int
}

type Genre struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Relevance int
}
