package normde

import (
	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/geom3d"
)

func JustXY(v geom3d.Point) geom2d.Point{
	return geom2d.Point{
		X: v.X,
		Y: v.Y,
	}
}
func XYWithZ(xy geom2d.Point, z geom3d.Fl) geom3d.Point{
	return geom3d.Point{
		X: xy.X,
		Y: xy.Y,
		Z: z,
	}
}
