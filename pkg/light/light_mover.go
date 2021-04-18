package light

import (
	"math"

	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/geom3d"
	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/normde"
)

func NewModesList() []gofill.LightMover {
	return []gofill.LightMover{
		NewRotation(),
		NewStationary(),
	}
}

type Rotation struct {
	mid        geom3d.Point
	radius     float64
	radianXY   float64
	radianZ    float64
	angleJump  float64
	zAmplitude float64
	zJump      float64
}

var _ gofill.LightMover = Rotation{}

func NewRotation(opts ...MovingCircleOpt) Rotation {
	r := Rotation{
		radius: 0.25,
		mid: geom3d.Point{
			X: 0.5,
			Y: 0.5,
			Z: 0.5,
		},
		radianXY:   0,
		angleJump:  math.Pi / 120,
		zAmplitude: 0.2,
		zJump:      0.03,

	}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

func (r Rotation) Move() gofill.LightMover {
	r.radianXY += r.angleJump
	r.radianZ += r.zJump
	return r
}

func (r Rotation) Name() string {
	return "Rotation"
}

func (r Rotation) Get() geom3d.Point {
	xy := normde.JustXY(r.mid).MoveByVector(geom2d.AngleToVecor(r.radianXY, r.radius))
	z := r.mid.Z + r.zAmplitude * math.Sin(r.radianZ)
	return normde.XYWithZ(xy, z)
}


type MovingCircleOpt func(*Rotation)

type Stationary struct{}

func (s Stationary) Move() gofill.LightMover {
	return s
}

func NewStationary() Stationary {
	return Stationary{}
}

var _ gofill.LightMover = Stationary{}

func (s Stationary) Get() geom3d.Point {
	return geom3d.Point{
		X: 0.5,
		Y: 0.5,
		Z: 10.0,
	}
}

func (s Stationary) Cleanup() {
}

func (s Stationary) Name() string {
	return "Stationary"
}
