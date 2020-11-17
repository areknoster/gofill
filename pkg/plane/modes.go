package plane

import (
	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/gofill"
)


func NewModesList(ss gofill.StateStorage) []gofill.PlaneMode {
	return []gofill.PlaneMode{
		NewMoveMesh(ss),
	}
}


type MoveMesh struct {
	storage gofill.StateStorage
	curVert int
}

func NewMoveMesh(ss gofill.StateStorage) *MoveMesh {
	return &MoveMesh{storage: ss, curVert: -1}
}

var _ gofill.PlaneMode = &MoveMesh{}

func (m *MoveMesh) HandleClick(normLoc geom2d.Point) {
}

func (m *MoveMesh) HandleDrag(start geom2d.Point, move geom2d.Vector) {
	state := m.storage.Get()
	points := state.Mesh.Points
	if m.curVert == -1 {
		m.curVert = points.ClosestToVertex(start)
	}

	newVert := points[m.curVert].MoveByVector(move)
	points[m.curVert] = newVert
	state.Mesh.Points = points
	m.storage.Set(state)
}

func (m *MoveMesh) HandleDragEnd() {
	m.curVert = -1
}

func (m *MoveMesh) Name() string {
	return "Move Mesh"
}
