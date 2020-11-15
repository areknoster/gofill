package light

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/areknoster/gofill/pkg/geom"
	"github.com/areknoster/gofill/pkg/gofill"
)

func NewModesList() []gofill.LightMode{
	return []gofill.LightMode{
		NewRotation(),
		NewStationary(),
	}
}

type Rotation struct {
	cancel   context.CancelFunc
	radius   float64
	mid      geom.Point
	radian   float64
	mx       *sync.Mutex
	interval time.Duration
	jump     float64
	position geom.Point
}

var _ gofill.LightMode = &Rotation{}

func NewRotation(opts ...MovingCircleOpt) gofill.LightMode {
	ctx, cancel := context.WithCancel(context.Background())
	mc := &Rotation{
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

func (r *Rotation) Name() string {
	return "Rotation"
}

func (r *Rotation) Get() geom.Point {
	return r.mid.MoveByVector(geom.AngleToVecor(r.radian, r.radius))
}

func (r *Rotation) Cleanup() {
	r.cancel()
}

type MovingCircleOpt func(*Rotation) error

type Stationary struct{}

func NewStationary() *Stationary {
	return &Stationary{}
}

var _ gofill.LightMode = &Stationary{}

func (s Stationary) Get() geom.Point {
	return geom.Point{
		X: 0.5,
		Y: 0.5,
	}
}

func (s Stationary) Cleanup() {
}

func (s Stationary) Name() string {
	return "Stationary"
}
