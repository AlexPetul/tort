package main

import (
	"aroom/internal/git"
	"aroom/internal/models"
	"aroom/internal/repository"
	"embed"
	"os"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var assets embed.FS

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.CatalogItem{},
		&models.FavouriteItem{},
		&models.Tag{},
	)

	repo := repository.NewRepository(db)
	gitClient := git.NewClient(os.Getenv("GITHUB_TOKEN"))

	app := NewApp(repo, gitClient)

	err = wails.Run(&options.App{
		Title:  "aroom",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
