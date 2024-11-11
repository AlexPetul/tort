package repository

import (
	"aroom/internal/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Repositorier interface {
	GetCatalogItem(ctx context.Context, id uint) (models.CatalogItem, error)
	GetFavouriteItem(ctx context.Context, id uint) (models.FavouriteItem, error)
	CreateCatalogItem(ctx context.Context, obj models.CatalogItem)
	CreateFavouriteItem(ctx context.Context, catalogItemID uint) (*models.FavouriteItem, error)
	CreateTag(ctx context.Context, obj models.Tag)
	ListTags(ctx context.Context, favouriteItemID uint) ([]models.TagResponse, error)
	ListCatalogItems(ctx context.Context) ([]models.CatalogItem, error)
	ListFavouriteItems(ctx context.Context) ([]models.FavouriteItem, error)
	DeleteFavouriteItem(ctx context.Context, obj models.FavouriteItem) error
	UpdateCurrentTag(ctx context.Context, favouriteItemID uint) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetCatalogItem(ctx context.Context, id uint) (models.CatalogItem, error) {
	var catalogItem models.CatalogItem
	r.db.Where("id = ?", id).Find(&catalogItem)
	return catalogItem, nil
}

func (r *Repository) CreateCatalogItem(ctx context.Context, obj models.CatalogItem) {
	r.db.Create(&obj)
}

func (r *Repository) ListCatalogItems(ctx context.Context) ([]models.CatalogItem, error) {
	items := []models.CatalogItem{}
	query := fmt.Sprintf(`
		SELECT *, EXISTS(
			SELECT (1)
			FROM %[1]s
			WHERE %[1]s.catalog_item_id = %[2]s.id
		) as is_favourite
		FROM %[2]s
		ORDER BY name
	`, models.FavouriteItem{}.TableName(), models.CatalogItem{}.TableName())
	r.db.Raw(query).Scan(&items)
	return items, nil
}

func (r *Repository) ListFavouriteItems(ctx context.Context) ([]models.FavouriteItem, error) {
	items := []models.FavouriteItem{}
	query := fmt.Sprintf(`
		SELECT
			fi.id,
			ci.name as CatalogItem__name,
			ci.avatar_url as CatalogItem__avatar_url,
			MAX(CASE WHEN t.current = true THEN t.name END) AS current_release,
			(SELECT t2.name
			FROM %[2]s t2
			WHERE t2.favourite_item_id = fi.id
			ORDER BY t2.id DESC
			LIMIT 1) AS latest_release
		FROM %[1]s fi
		JOIN %[3]s ci on ci.id = fi.catalog_item_id
		JOIN %[2]s t on fi.id = t.favourite_item_id
		GROUP BY fi.id, ci.name
		ORDER BY ci.name;
	`, models.FavouriteItem{}.TableName(), models.Tag{}.TableName(), models.CatalogItem{}.TableName())
	r.db.Raw(query).Scan(&items)
	return items, nil
}

func (r *Repository) CreateFavouriteItem(ctx context.Context, catalogItemID uint) (*models.FavouriteItem, error) {
	favouriteItem := models.FavouriteItem{CatalogItemID: catalogItemID}
	r.db.Create(&favouriteItem)
	return &favouriteItem, nil
}

func (r *Repository) ListTags(ctx context.Context, favouriteItemID uint) ([]models.TagResponse, error) {
	items := []models.TagResponse{}
	query := fmt.Sprintf(`
		SELECT
			MAX(CASE WHEN current = true THEN name END) AS current,
			(SELECT name FROM tags WHERE favourite_item_id = 7 ORDER BY id DESC LIMIT 1) AS latest
		FROM %[1]s
		WHERE favourite_item_id = %[2]d
	`, models.Tag{}.TableName(), favouriteItemID)
	r.db.Raw(query).Scan(&items)
	return items, nil
}

func (r *Repository) CreateTag(ctx context.Context, obj models.Tag) {
	r.db.Create(&obj)
}

func (r *Repository) GetFavouriteItem(ctx context.Context, id uint) (models.FavouriteItem, error) {
	var obj models.FavouriteItem
	r.db.Where("id = ?", id).Find(&obj)
	return obj, nil
}

func (r *Repository) DeleteFavouriteItem(ctx context.Context, obj models.FavouriteItem) error {
	r.db.Delete(&obj)
	return nil
}

func (r *Repository) UpdateCurrentTag(ctx context.Context, favouriteItemID uint) error {
	r.db.Model(&models.Tag{}).Where("favourite_item_id = ?", favouriteItemID).Update("current", false)
	r.db.Model(&models.Tag{}).Where("favourite_item_id = ?", favouriteItemID).Where("id = (?)", r.db.Model(&models.Tag{}).Select("MAX(id)")).Update("current", true)
	return nil
}
