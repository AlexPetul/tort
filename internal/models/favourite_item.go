package models

type FavouriteItem struct {
	ID             uint        `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	CatalogItemID  uint        `json:"-"`
	CatalogItem    CatalogItem `json:"catalog_item"`
	CurrentRelease string      `gorm:"->" json:"current_release"`
	LatestRelease  string      `gorm:"->" json:"latest_release"`
}

func (FavouriteItem) TableName() string {
	return "favourite_items"
}
