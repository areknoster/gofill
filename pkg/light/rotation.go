package light

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/areknoster/gofill/pkg/geom"
	"github.com/areknoster/gofill/pkg/gofill"
)

type MovingCircle struct {
	cancel context.CancelFunc
	radius   float64
	mid      geom.Point
	radian   float64
	mx       *sync.Mutex
	interval time.Duration
	jump     float64
	position geom.Point
}

var _ gofill.LightPosition = &MovingCircle{}

func (mc *MovingCircle) Get() geom.Point {
	return mc.mid.MoveByVector(geom.AngleToVecor(mc.radian, mc.radius))
}

func (mc *MovingCircle) Cleanup() {
	mc.cancel()
}

type MovingCircleOpt func(*MovingCircle) error

func NewRotatingMovement(opts ...MovingCircleOpt) gofill.LightPosition {
	ctx, cancel := context.WithCancel(context.Background())
	mc := &MovingCircle{
		cancel: cancel,
		radius: 0.25,
		mid: geom.Point{
			X: 0.5,
			Y: 0.5,
		},
		radian:   0,
		mx:       &sync.Mutex{},
		interval: 20 * time.Millisecond,
		jump:     math.Pi / 30,
	}
	for _, opt := range opts {
		opt(mc)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(mc.interval):
				mc.mx.Lock()
				mc.radian += mc.jump
				mc.mx.Unlock()
			}
		}
	}()
	return mc
}
