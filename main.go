package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/plane"
	"github.com/areknoster/gofill/pkg/render"
	"github.com/areknoster/gofill/pkg/state_storage"
	"github.com/areknoster/gofill/pkg/ui"
)

type Config struct{
	title string
	canvasSize fyne.Size
}

func main(){
	logrus.SetLevel(logrus.ErrorLevel)
	cfg := Config{
		title:      "GoFill",
		canvasSize: fyne.Size{700, 700},
	}

	fyneApp := app.New()
	window :=fyneApp.NewWindow(cfg.title)

	ss := state_storage.NewStateStorage(cfg.canvasSize)

	renderer := render.NewRenderer(ss)

	nPlane, setMode := plane.NewPlane(renderer, cfg.canvasSize)
	ss.Refresh = nPlane.Refresh
	setMode(plane.NewMoveMesh(ss))

	menu := ui.NewMenuContainer(ss, window)
	container := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, menu, nil), menu, nPlane)
	window.SetContent(container)
	window.SetFixedSize(true)
	go func(){
		for{
			time.Sleep(70 * time.Millisecond)
			state := ss.Get()
			state.Light.SourceMovement = state.Light.SourceMovement.Move()
			ss.Set(state)
		}
	}()
	window.ShowAndRun()


}
