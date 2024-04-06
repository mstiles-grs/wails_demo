package main

import (
    "embed"
    "log"

    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
  "github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

    app := &App{}

    err := wails.Run(&options.App{
        Title:  "Basic Demo",
        Width:  1024,
        Height: 768,
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        OnStartup:  app.startup,
        OnShutdown: app.shutdown,
        Bind: []interface{}{
            app,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
