package render

import (
	"image"

	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/normde"
	"github.com/areknoster/gofill/pkg/raster"
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

func (pm *PreciseMode) Render(state gofill.State) *image.RGBA {
	w,h := state.Size.Width, state.Size.Height
	render := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			render.SetRGBA(i, j, lambertPixel(&state, i, j))
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

func (i *InterpolationMode) Render(state gofill.State) *image.RGBA {
	w,h := state.Size.Width, state.Size.Height
	render := image.NewRGBA(image.Rect(0, 0, w, h))
	triangles := state.Mesh.GetTriangles()
	for _, triangle := range triangles {
		raster.ScanLine(
			triangle,
			render,
			i.interpolateTriangle(render, state, triangle))
	}

	return render
}

func (i *InterpolationMode) Name() string {
	return "Interpolation"
}

func (i *InterpolationMode) interpolateTriangle(render *image.RGBA, state gofill.State, triangle geom2d.Shape) func(p raster.Pixel) {
	w, h := state.Size.Width - 1, state.Size.Height - 1
	A := triangle.PointsSet[0]
	Apx := normde.DenormPoint2D(triangle.PointsSet[0], w,h)
	colA := lambertColor(&state, Apx.X, Apx.Y)
	B := triangle.PointsSet[1]
	Bpx := normde.DenormPoint2D(triangle.PointsSet[1], w,h)
	colB := lambertColor(&state, Bpx.X, Bpx.Y)
	C := triangle.PointsSet[2]
	Cnorm := normde.DenormPoint2D(triangle.PointsSet[2], w,h)
	colC := lambertColor(&state, Cnorm.X, Cnorm.Y)
	AB := A.VectorTo(B)
	AC := A.VectorTo(C)
	N := normde.CrossFlat(AB, AC).Length()

	return func(p raster.Pixel){
		P := normde.NormPoint2D(p, w,h)
		w := normde.CrossFlat(AB, A.VectorTo(P)).Length() / N
		v := normde.CrossFlat(AC, A.VectorTo(P)).Length() / N
		u := 1 - w - v
		col := colA.TimesScalar(u).Add(colB.TimesScalar(v)).Add(colC.TimesScalar(w))
		if u < 0 {
			return
		}
		render.SetRGBA(p.X, p.Y, normde.DenormVecToRGBA(col))
	}
}
