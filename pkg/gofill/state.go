package gofill

import (
	"image"

	"github.com/areknoster/gofill/pkg/geom"
)




type State struct{
	Mesh      geom.Mesh
	ShowMesh  bool
	Light     LightConfig
	Texture   *image.RGBA
	NormalMap *image.RGBA
	PlaneMode PlaneMode
	RendererMode RendererMode
}

type StateStorage interface{
	Get() State
	Set(State)
}

