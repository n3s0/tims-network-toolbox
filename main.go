package main

import (
	services "changeme/internal/services"
	"embed"
	_ "embed"
	"log"
	
	"github.com/wailsapp/wails/v3/pkg/application"
)

var assets embed.FS

func init() {
}

func main() {

	app := application.New(application.Options{
		Name:        "Tim's Network Toolbox",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&services.PingService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Tim's Network Toolbox",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
