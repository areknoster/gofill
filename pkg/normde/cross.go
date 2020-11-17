package normde

import (
	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/geom3d"
)

func CrossFlat(a, b geom2d.Vector) geom3d.Vector {
	k := geom3d.Vector{a.X, a.Y, 0}
	l := geom3d.Vector{b.X, b.Y, 0}
	return k.Cross(l)
}
