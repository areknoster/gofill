package gofill

import (
	"context"
	"image/color"

	"github.com/areknoster/gofill/pkg/geom"
)

type LightConfig struct {
	SourceMovement LightMode
	Color          color.RGBA
	Ks             float64
	Kd             float64
	M              float64
}

type LightMode interface{
	Get() geom.Point
	Cleanup()
	Name() string
}

type LightPositionInit func() (LightMode, context.CancelFunc)
