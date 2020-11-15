package state

import (
	"image"
	"image/color"
	"sync"

	"github.com/areknoster/gofill/pkg/geom"
	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/light"
	"github.com/areknoster/gofill/pkg/modes"
	"github.com/areknoster/gofill/pkg/render"
)

type StateStorage struct{
	state          gofill.State
	stateMx        *sync.Mutex
}

var _ gofill.StateStorage = &StateStorage{}

func NewStateStorage(img *image.RGBA) *StateStorage{
	ss := &StateStorage{
		state:   gofill.State{
			Mesh:      geom.NewMesh(10, 10),
			Light:     gofill.LightConfig{
				SourceMovement: light.NewRotatingMovement(),
				Color:          render.ColorToRGBA(color.White),
				Ks:             0.5,
				Kd:             0.5,
				M:              30.0,
			},
			Texture:   img,
			NormalMap: img,
		},
		stateMx: &sync.Mutex{},
	}
	ss.state.Mode = modes.NewMoveMesh(ss)
	return ss
}

func (sm StateStorage) Get() gofill.State {
	sm.stateMx.Lock()
	defer sm.stateMx.Unlock()
	return sm.state
}

func (sm StateStorage) Set(state gofill.State) {
	sm.stateMx.Lock()
	defer sm.stateMx.Unlock()
	sm.state = state

}
