package models

type Tag struct {
	ID              uint          `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Name            string        `json:"name"`
	Current         bool          `json:"current"`
	FavouriteItemID uint          `json:"-"`
	FavouriteItem   FavouriteItem `gorm:"constraint:OnDelete:CASCADE" json:"favourite_item"`
}

type TagResponse struct {
	Current string `json:"current"`
	Latest  string `json:"latest"`
}

func (Tag) TableName() string {
	return "tags"
}
