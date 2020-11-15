package render

import (
	"image"

	"github.com/areknoster/gofill/pkg/gofill"
)

type PreciseMode struct{}

func NewRendererModesList() []gofill.RendererMode {
	return []gofill.RendererMode{
		NewPreciseMode(),
		NewInterpolationMode(),
	}
}

func NewPreciseMode() *PreciseMode {
	return &PreciseMode{}
}

var _ gofill.RendererMode = &PreciseMode{}

func (pm *PreciseMode) Render(state gofill.State) image.Image {
	w,h := state.Size.Width, state.Size.Height
	render := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			render.SetRGBA(i, j, lambertColor(state, i, j))
		}
	}
	return render
}

func (pm *PreciseMode) Name() string {
	return "Precise"
}

type InterpolationMode struct{}

func NewInterpolationMode() *InterpolationMode {
	return &InterpolationMode{}
}

var _ gofill.RendererMode = &InterpolationMode{}

func (i *InterpolationMode) Render(state gofill.State) image.Image {
	w,h := state.Size.Width, state.Size.Height
	render := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			render.SetRGBA(i, j, lambertColor(state, i, j))
		}
	}
	return render
}

func (i *InterpolationMode) Name() string {
	return "Interpolation"
}
