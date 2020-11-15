package gofill

import (
	"image"

	"github.com/areknoster/gofill/pkg/geom"
)




type State struct{
	Mesh      geom.Mesh
	Light     LightConfig
	Texture   *image.RGBA
	NormalMap *image.RGBA
	Mode Mode
}

type StateStorage interface{
	Get() State
	Set(State)
}

