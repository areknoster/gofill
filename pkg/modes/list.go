package modes

import "github.com/areknoster/gofill/pkg/gofill"

func NewModesList(ss gofill.StateStorage) []gofill.Mode{
	return []gofill.Mode{
		NewMoveMesh(ss),
	}
}
