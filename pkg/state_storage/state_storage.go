package state_storage

import (
	"sync"

	"fyne.io/fyne"

	"github.com/areknoster/gofill/pkg/geom3d"
	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/plane"
)

type StateStorage struct {
	state   gofill.State
	stateMx *sync.Mutex
	Refresh func()
}

var _ gofill.StateStorage = &StateStorage{}

func NewStateStorage(size fyne.Size) *StateStorage {
	state := gofill.State{
		Light: gofill.LightConfig{
			ColorVector: geom3d.Vector{1.0, 1.0, 1.0},
			Ks:          0.5,
			Kd:          0.5,
			M:           10.0,
		},
		Size: size,
	}

	ss := &StateStorage{
		stateMx: &sync.Mutex{},
		Refresh: func() {},
	}
	state.PlaneMode = plane.NewMoveMesh(ss)
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
	sm.state = state
	sm.stateMx.Unlock()
	sm.Refresh()
}
