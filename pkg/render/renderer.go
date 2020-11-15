package render

import (
	"image"

	"github.com/areknoster/gofill/pkg/gofill"
)

type Renderer struct{
	stateStorage gofill.StateStorage
}

func NewRenderer(stateStorage gofill.StateStorage) *Renderer {
	return &Renderer{stateStorage: stateStorage}
}

func (r *Renderer) Render( w, h int) image.Image {
	return r.stateStorage.Get().Texture
}

var _ gofill.Renderer = &Renderer{}
