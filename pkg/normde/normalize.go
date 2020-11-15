package normde

import (
	"github.com/areknoster/gofill/pkg/geom"
	"github.com/areknoster/gofill/pkg/render"
)

func NormPoint(point render.Pixel, w, h int) geom.Point {
	return geom.Point{
		X: float64(point.X)/float64(w),
		Y: 1.0 - float64(point.Y)/float64(h),
	}
}

func NormVector(vec render.Pixel, w,h int ) geom.Vector {
	return geom.Vector{
		X: float64(vec.X)/float64(w),
		Y: -float64(vec.Y)/float64(h),
	}
}

func DenormPoint(p geom.Point, w,h int) (render.Pixel){
	return render.Pixel{int(float64(w) * p.X), int(float64(w) *(1.0 - p.Y))}
}

