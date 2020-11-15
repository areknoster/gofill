package normde

import (
	"image"
	"image/color"

	"fyne.io/fyne"

	"github.com/areknoster/gofill/pkg/geom3d"
)

func RGBAToSizedNormMap(nonSized *image.RGBA, size fyne.Size) *geom3d.VectorMap {
	img := ResizeRGBA(nonSized, size)
	nm := geom3d.NewVectorMap(size.Width, size.Height)
	for i := 0; i < size.Width; i++ {
		for j := 0; j < size.Height; j++ {
			pixel := img.RGBAAt(i, j)
			nm.SetAt(i, j, RGBAToNormVersor(pixel))
		}
	}
	return nm
}

func RGBAToNormVersor(rgba color.RGBA) geom3d.Vector {
	v := geom3d.Vector{
		X: geom3d.Fl(int(rgba.R)-127)/128.0,
		Y: geom3d.Fl(int(rgba.G)-127)/128.0,
		Z: geom3d.Fl(rgba.B)/255.0,
	}
	return v.ToVersor()
}
