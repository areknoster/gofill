package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/areknoster/gofill/pkg/gofill"
)

type Menu struct {
	storage gofill.StateStorage
	window  fyne.Window
}

func (m *Menu) setState(mutate func(state *gofill.State)) {
	state := m.storage.Get()
	mutate(&state)
	m.storage.Set(state)
}

func NewMenuContainer(storage gofill.StateStorage, window fyne.Window) *fyne.Container {
	m := &Menu{
		storage: storage,
		window: window,
	}

	vMenu := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		m.newLightSection(),
		widget.NewSeparator(),
		m.newMeshSection(),
		widget.NewSeparator(),
		m.newTextureSection(),
		widget.NewSeparator(),
		m.newNormalMapSelection(),
		widget.NewSeparator(),
		m.newRenderingSection(),
	)

	return vMenu
}
