package state

import (
	"image/color"
	"sync"

	"github.com/areknoster/gofill/pkg/geom"
	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/light"
	"github.com/areknoster/gofill/pkg/modes"
	"github.com/areknoster/gofill/pkg/render"
)

type StateStorage struct {
	state   gofill.State
	stateMx *sync.Mutex
	Refresh func()
}

var _ gofill.StateStorage = &StateStorage{}

func NewStateStorage() *StateStorage {
	state := gofill.State{
		Mesh: geom.NewMesh(10, 10),
		Light: gofill.LightConfig{
			SourceMovement: light.NewRotation(),
			Color:          render.ColorToRGBA(color.White),
			Ks:             0.5,
			Kd:             0.5,
			M:              30.0,
		},
	}
	//li := images.ChooseNormalMap
	//li.Set(
	//	li.ListAvailable()[1], func(rgba *image.RGBA) {
	//		//state.Texture = rgba
	//		state.NormalMap = rgba
	//	})

	ss := &StateStorage{
		stateMx: &sync.Mutex{},
		Refresh: func() {
		},
	}
	state.Mode = modes.NewMoveMesh(ss)
	ss.state = state
return ss
}

func (sm *StateStorage) Get() gofill.State {
	sm.stateMx.Lock()
	defer sm.stateMx.Unlock()
	return sm.state
}

func (sm *StateStorage) Set(state gofill.State) {
	sm.stateMx.Lock()
	defer sm.stateMx.Unlock()
	sm.state = state
	sm.Refresh()

}
