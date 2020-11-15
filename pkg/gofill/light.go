package gofill

import (
	"context"

	"github.com/areknoster/gofill/pkg/geom3d"
)

type LightConfig struct {
	SourceMovement LightMover
	ColorVector    geom3d.Vector
	Ks             float64
	Kd             float64
	M              float64
}

type LightMover interface{
	Get() geom3d.Point
	Name() string
	Move() LightMover
}

type LightPositionInit func() (LightMover, context.CancelFunc)
