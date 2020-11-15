package gofill

import (
	"image"

	"github.com/areknoster/gofill/pkg/geom"
)

type Plane interface {
	Refresh()
}

type Renderer interface {
	Render(w, h int) image.Image
}

type RendererMode interface {
	Render(state State, w, h int) image.Image
	Name() string
}

type PlaneMode interface {
	HandleClick(normLoc geom.Point)
	HandleDrag(start geom.Point, move geom.Vector)
	HandleDragEnd()
	Name() string
}
