package gofill

import (
	"context"
	"image/color"

	"github.com/areknoster/gofill/pkg/geom"
)

type LightConfig struct {
	SourceMovement LightPosition
	Color          color.RGBA
	Ks             float64
	Kd             float64
	M              float64
}

type LightPosition interface{
	Get() geom.Point
	Cleanup()
}

type LightPositionInit func() (LightPosition, context.CancelFunc)
