package main

import (
	"embed"
	"log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/mstiles-grs/wails_demo/playbook/sql_database/sqlDB"

)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
	db *sql.DB
}

// NewApp creates a new App instance
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup() {
	// Initialize your database here
	db := sql_database.AppStartUp()


func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "playbook",
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
