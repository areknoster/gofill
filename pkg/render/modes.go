package render

import (
	"image"

	"github.com/areknoster/gofill/pkg/gofill"
)

type PreciseMode struct{}

func NewRendererModesList()[]gofill.RendererMode{
	return []gofill.RendererMode{
		NewPreciseMode(),
		NewInterpolationMode(),
	}
}

func NewPreciseMode() *PreciseMode {
	return &PreciseMode{}
}

var _ gofill.RendererMode = &PreciseMode{}

func (pm *PreciseMode) Render(state gofill.State, w, h int) image.Image {
	return state.Texture
}

func (pm *PreciseMode) Name() string {
	return "Precise"
}

type InterpolationMode struct{}

func NewInterpolationMode() *InterpolationMode {
	return &InterpolationMode{}
}

var _ gofill.RendererMode = &InterpolationMode{}

func (i *InterpolationMode) Render(state gofill.State, w, h int) image.Image {
	return state.Texture
}

func (i *InterpolationMode) Name() string {
	return "Interpolation"
}


