package render

import (
	"image"

	"github.com/areknoster/gofill/pkg/gofill"
)

type Renderer struct {
	stateStorage gofill.StateStorage
}

func NewRenderer(stateStorage gofill.StateStorage) *Renderer {
	state := stateStorage.Get()
	state.RendererMode = NewPreciseMode()
	stateStorage.Set(state)
	return &Renderer{stateStorage: stateStorage}
}

func (r *Renderer) Render(w, h int) image.Image {
	state := r.stateStorage.Get()
	return state.RendererMode.Render(state)
}

var _ gofill.Renderer = &Renderer{}




