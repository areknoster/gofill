package gofill

import (
	"image"

	"github.com/areknoster/gofill/pkg/geom"
)

type Plane interface{
	Refresh()
}

type Renderer interface{
	Render(w, h int) image.Image
}

type Mode interface {
	HandleClick(normLoc geom.Point)
	HandleDrag(start geom.Point, move geom.Vector)
	HandleDragEnd()
	Name() string
}


