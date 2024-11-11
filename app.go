package main

import (
	"aroom/internal/git"
	"aroom/internal/models"
	"aroom/internal/preload"
	"aroom/internal/repository"
	"context"
	"fmt"
)

type App struct {
	ctx  context.Context
	git  git.Git
	repo repository.Repositorier
}

func NewApp(repo *repository.Repository, git *git.GitClient) *App {
	return &App{
		git:  git,
		repo: repo,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	items, err := a.repo.ListCatalogItems(a.ctx)
	if err != nil {
		panic(err)
	}
	if len(items) == 0 {
		projects, err := preload.ReadInitial()
		if err != nil {
			panic(err)
		}

		for _, v := range projects {
			repo := a.git.GetRepo(v.Owner, v.Name)

			a.repo.CreateCatalogItem(a.ctx, models.CatalogItem{
				Name:      *repo.Name,
				Stars:     uint(*repo.StargazersCount),
				Owner:     v.Owner,
				AvatarUrl: *repo.Owner.AvatarURL,
				GitURL:    fmt.Sprintf("https://github.com/%s", *repo.FullName),
			})
		}
	}
}

func (a *App) LoadCatalog() []models.CatalogItem {
	items, err := a.repo.ListCatalogItems(a.ctx)
	if err != nil {
		panic(err)
	}
	return items
}

func (a *App) LoadFavourites() []models.FavouriteItem {
	items, err := a.repo.ListFavouriteItems(a.ctx)
	if err != nil {
		panic(err)
	}
	return items
}

func (a *App) ListFavouriteItemTags(id uint) []models.TagResponse {
	items, err := a.repo.ListTags(a.ctx, id)
	if err != nil {
		panic(err)
	}
	return items
}

func (a *App) AddToFavourites(id uint) {
	catalogItem, err := a.repo.GetCatalogItem(a.ctx, id)
	if err != nil {
		panic(err)
	}
	favouriteItem, err := a.repo.CreateFavouriteItem(a.ctx, catalogItem.ID)
	if err != nil {
		panic(err)
	}

	tags := a.git.GetRepoTags(catalogItem.Owner, catalogItem.Name)
	tag := tags[0]
	a.repo.CreateTag(a.ctx, models.Tag{Name: *tag.Name, FavouriteItemID: favouriteItem.ID, Current: true})
}

func (a *App) DeleteFavouriteItem(id uint) {
	obj, err := a.repo.GetFavouriteItem(a.ctx, id)
	if err != nil {
		panic(err)
	}
	a.repo.DeleteFavouriteItem(a.ctx, obj)
}

func (a *App) AlignTags(favouriteItemID uint) {
	a.repo.UpdateCurrentTag(a.ctx, favouriteItemID)
}
