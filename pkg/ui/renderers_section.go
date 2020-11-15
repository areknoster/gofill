package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/render"
)

func (m *Menu) newRenderingSection() fyne.CanvasObject{
	modes := render.NewRendererModesList()
	names := make([]string, len(modes))
	for i, mode := range modes {
		names[i] = mode.Name()
	}
	modesRadio := widget.NewRadioGroup(
		names, func(n string) {
			for _, mode := range modes {
				if n == mode.Name() {
					m.setState(
						func(state *gofill.State) {
							logrus.Debugf("renderer mode changed to %s", mode.Name())
							state.RendererMode = mode
						})
				}
			}
		})
	modesRadio.SetSelected(names[0])
	modesRadio.OnChanged(modesRadio.Selected)
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle("Rendering", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:      true,
		}),
		modesRadio,
	)
}

