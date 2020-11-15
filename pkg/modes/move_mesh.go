package modes

import (
	"github.com/areknoster/gofill/pkg/geom"
	"github.com/areknoster/gofill/pkg/gofill"
)

type MoveMesh struct{
	ss gofill.StateStorage
}

func NewMoveMesh(ss gofill.StateStorage) *MoveMesh {
	return &MoveMesh{ss: ss}
}

var _ gofill.Mode = &MoveMesh{}


func (m *MoveMesh) HandleClick(normLoc geom.Point) {

}

func (m *MoveMesh) HandleDrag(start geom.Point, move geom.Vector) {
}

func (m *MoveMesh) HandleDragEnd() {
}

func (m *MoveMesh) Name() string {
	return "Move Mesh"
}

