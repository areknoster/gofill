package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/modes"
	"github.com/areknoster/gofill/pkg/plane"
	"github.com/areknoster/gofill/pkg/render"
	"github.com/areknoster/gofill/pkg/state"
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


	//cb := func(rc fyne.URIReadCloser, err error){
	//	if err != nil{
	//		logrus.Error("could not read file")
	//		return
	//	}
	//	defer rc.Close()
	//	img, err = jpeg.Decode(rc)
	//	if err != nil{
	//		logrus.Errorf("could not decode image: %s", err.Error())
	//	}
	//}

	//fd := dialog.NewFileOpen(cb, window)

	var img image.Image
	func(){
		file, err := os.OpenFile("resources/bricks.jpg", os.O_RDONLY, 0)
		if err != nil{
			logrus.Panicf("could not open file: %s", err.Error())
		}
		img, err = jpeg.Decode(file)
		if err != nil{
			logrus.Panicf("could not decode image: %s", err.Error())
		}
	}()


	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)
	ss := state.NewStateStorage(rgba)

	renderer := render.NewRenderer(ss)

	plane, setMode := plane.NewPlane(renderer, cfg.canvasSize)
	plane.Refresh()
	setMode(modes.NewMoveMesh(ss))
	//menu, setActiveEditMenu := ui.NewMenu(plane)
	//plane.HandleSelect = setActiveEditMenu

	//container := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, menu, nil), menu, plane)
	window.SetContent(plane)
	window.SetFixedSize(true)

	window.ShowAndRun()


}
