package normde

import (
	"image/color"
	"math"

	"github.com/areknoster/gofill/pkg/geom3d"
)

func NormRGBAToVec(rgba color.RGBA) geom3d.Vector {
	return geom3d.Vector{
		X: geom3d.Fl(rgba.R) / 255.0,
		Y: geom3d.Fl(rgba.G) / 255.0,
		Z: geom3d.Fl(rgba.B) / 255.0,
	}
}


func DenormVecToRGBA(v geom3d.Vector) color.RGBA{

	return color.RGBA{
		R: uint8(math.Min(v.X, 1.0) * 255.0),
		G: uint8(math.Min(v.Y, 1.0) * 255),
		B: uint8(math.Min(v.Z, 1.0) * 255),
		A: 255,
	}
}
