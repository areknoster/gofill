package raster

import (
	"fmt"
	"image"
	"image/color"
	"testing"

	"github.com/areknoster/gofill/pkg/geom2d"
)

func TestSth(t *testing.T){
	shape := geom2d.Shape{[]geom2d.Point{
		{0.0, 0.2},
		{1.0, 0.2},
		{0.5, 1.0},
	}}
	img := image.NewRGBA(image.Rect(0,0,20,20))
	ScanLine(shape, img, func(p Pixel) {
		fmt.Printf("set(%d, %d)", p.X, p.Y)
		img.SetRGBA(p.X, p.Y, color.RGBA{255,255,255,255})
	})
	rep := "Rep:\n"
	for i := 0; i < 20; i++{
		for j := 0; j < 20; j++{
			if img.RGBAAt(j,i).R == 255{
				rep += "X"
			}else {
				rep+= "O"
			}
		}
		rep += "\n"
	}
	fmt.Print(rep)
}
