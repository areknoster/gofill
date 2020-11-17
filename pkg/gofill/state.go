package gofill

import (
	"fyne.io/fyne"

	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/geom3d"
)

type State struct {
	Mesh         geom2d.Mesh `json:"-"`
	ShowMesh     bool
	Light        LightConfig
	Size         fyne.Size
	Texture      *geom3d.VectorMap `json:"-"`
	NormalMap    *geom3d.VectorMap `json:"-"`
	PlaneMode    PlaneMode
	RendererMode RendererMode
	WavesCoef    float64
}

type StateStorage interface {
	Get() State
	Set(State)
}
