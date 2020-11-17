package gofill

import (
	"image"

	"github.com/areknoster/gofill/pkg/geom2d"
)

type Plane interface {
	Refresh()
}

type Renderer interface {
	Render(w, h int) image.Image
}

type RendererMode interface {
	Render(state State) *image.RGBA
	Name() string
}

type PlaneMode interface {
	HandleClick(normLoc geom2d.Point)
	HandleDrag(start geom2d.Point, move geom2d.Vector)
	HandleDragEnd()
	Name() string
}
