package modes

import "github.com/areknoster/gofill/pkg/gofill"

func NewModesList(ss gofill.StateStorage) []gofill.PlaneMode {
	return []gofill.PlaneMode{
		NewMoveMesh(ss),
	}
}
