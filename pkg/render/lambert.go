package render

import (
	"image/color"
	"math"

	"github.com/areknoster/gofill/pkg/geom3d"
	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/normde"
	"github.com/areknoster/gofill/pkg/raster"
)

func lambertColor(state *gofill.State, x, y int) geom3d.Vector{
	kd := state.Light.Kd
	ks := state.Light.Ks
	m := state.Light.M
	ll := state.Light.ColorVector
	lo := state.Texture.At(x, y)
	normP := normde.NormPoint2D(raster.Pixel{x, y}, state.Size.Width, state.Size.Height)
	L := normde.XYWithZ(normP, 0.0).
		VectorTo(state.Light.SourceMovement.Get()).
		ToVersor()
	N := state.NormalMap.At(x, y)
	V := geom3d.Vector{0, 0, 1.0}
	R := N.TimesScalar(2 * N.Dot(L)).Substract(L)
	if R.Z < 0 {
		R = geom3d.Vector{0,0,0}
	}
	cosNL := N.Dot(L)
	if cosNL <=0{
		return geom3d.Vector{0,0,0}
	}
	cosVRpowm := math.Pow(V.Dot(R), m)

	crossLlLo := ll.MultiplyByElem(lo)
	dPart := crossLlLo.TimesScalar(kd * cosNL)
	sPart := crossLlLo.TimesScalar(ks * cosVRpowm)
	return sPart.Add(dPart)
}

func lambertPixel(state *gofill.State, x, y int) color.RGBA {
	return normde.DenormVecToRGBA(lambertColor(state,x,y))
}
