package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/modes"
	"github.com/areknoster/gofill/pkg/plane"
	"github.com/areknoster/gofill/pkg/render"
	"github.com/areknoster/gofill/pkg/state"
	"github.com/areknoster/gofill/pkg/ui"
)

type Config struct{
	title string
	canvasSize fyne.Size
}

func main(){
	logrus.SetLevel(logrus.DebugLevel)
	cfg := Config{
		title:      "GoFill",
		canvasSize: fyne.Size{600, 600},
	}

	fyneApp := app.New()
	window :=fyneApp.NewWindow(cfg.title)



	ss := state.NewStateStorage()
	menu := ui.NewMenuContainer(ss, window)

	renderer := render.NewRenderer(ss)

	plane, setMode := plane.NewPlane(renderer, cfg.canvasSize)
	ss.Refresh = plane.Refresh
	setMode(modes.NewMoveMesh(ss))

	//plane.HandleSelect = setActiveEditMenu

	container := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, menu, nil), menu, plane)
	window.SetContent(container)
	window.SetFixedSize(true)

	window.ShowAndRun()


}
