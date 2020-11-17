package render

import (
	"image"

	"golang.org/x/image/colornames"

	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/normde"
	"github.com/areknoster/gofill/pkg/raster"
)

func putMesh(state gofill.State, img *image.RGBA) *image.RGBA{
	if !state.ShowMesh{
		return img
	}
	segments := state.Mesh.GetSegments()
	brush := raster.NewSquareBrush(1, img, raster.ColorToRGBA(colornames.Blueviolet))
	for _, segment := range segments {
		a:= normde.DenormPoint2D(segment.A, state.Size.Width, state.Size.Height)
		b:= normde.DenormPoint2D(segment.B, state.Size.Width, state.Size.Height)
		raster.BresenhamLine(a,b,brush)
	}
	return img
}
