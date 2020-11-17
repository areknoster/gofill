package normde

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne"

	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/geom3d"
	"github.com/areknoster/gofill/pkg/raster"
)

func RGBAToSizedNormMap(nonSized *image.RGBA, size fyne.Size) *geom3d.VectorMap {
	img := ResizeRGBA(nonSized, size)
	nm := geom3d.NewVectorMap(size.Width, size.Height)
	for i := 0; i <= size.Width; i++ {
		for j := 0; j <= size.Height; j++ {
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

func NewUniform(size fyne.Size) *geom3d.VectorMap {
	nm := geom3d.NewVectorMap(size.Width, size.Height)
	for i := 0; i <= size.Width; i++ {
		for j := 0; j <= size.Height; j++ {
			nm.SetAt(
				i, j, geom3d.Vector{
					X: 0,
					Y: 0,
					Z: 1,
				})
		}
	}
	return nm
}

func NewWave(size fyne.Size, coef float64) *geom3d.VectorMap {
	nm := geom3d.NewVectorMap(size.Width, size.Height)
	mid := geom2d.Point{0.5, 0.5}
	for i := 0; i <= size.Width; i++ {
		for j := 0; j <= size.Height; j++ {
			np := NormPoint2D(raster.Pixel{i, j}, size.Width, size.Height)
			k := geom2d.VecBetweenPoints(mid, np).Length() * math.Sqrt(coef)

			v1 := geom3d.Vector{
				X: math.Cos(k),
				Y: 0,
				Z: math.Sin(k),
			}
			v2 := geom3d.Vector{
				X: 0,
				Y: math.Cos(k),
				Z: math.Sin(k),
			}

			nm.SetAt(i, j, v1.Cross(v2).ToVersor())
		}
	}
	return nm
}
