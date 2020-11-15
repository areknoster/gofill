package modes

import (
	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/gofill"
)

type MoveMesh struct{
	ss gofill.StateStorage
}

func NewMoveMesh(ss gofill.StateStorage) *MoveMesh {
	return &MoveMesh{ss: ss}
}

var _ gofill.PlaneMode = &MoveMesh{}


func (m *MoveMesh) HandleClick(normLoc geom2d.Point) {

}

func (m *MoveMesh) HandleDrag(start geom2d.Point, move geom2d.Vector) {
}

func (m *MoveMesh) HandleDragEnd() {
}

func (m *MoveMesh) Name() string {
	return "Move Mesh"
}

