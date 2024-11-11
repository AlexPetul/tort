package models

type CatalogItem struct {
	ID          uint   `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Stars       uint   `json:"stars"`
	Owner       string `json:"owner"`
	GitURL      string `json:"git_url"`
	AvatarUrl   string `json:"avatar_url"`
	IsFavourite bool   `gorm:"->" json:"is_favourite"`
}

func (CatalogItem) TableName() string {
	return "catalog_items"
}
