package ui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/light"
	"github.com/areknoster/gofill/pkg/render"
)

func (m *Menu) newLightSection() fyne.CanvasObject {
	colorButton := widget.NewButtonWithIcon(
		"Pick color", theme.ColorPaletteIcon(), func() {
			colorDialog := dialog.NewColorPicker(
				"Pick light color",
				"",
				func(c color.Color) {
					m.setState(
						func(state *gofill.State) {
							state.Light.Color = render.ColorToRGBA(c)
						})
				},
				m.window,
			)
			colorDialog.Advanced = true
			colorDialog.Show()
		})


	modes := light.NewModesList()
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
							logrus.Debugf("light mode changed to %s", mode.Name())
							state.Light.SourceMovement = mode
						})
				}
			}
		})
	ksSlider := widget.NewSlider(0.0, 1.0)
	ksSlider.Step = 0.01
	ksSlider.OnChanged = func(f float64) {
		m.setState(
			func(state *gofill.State) {
				state.Light.Ks = f
			})
		ksSlider.Value = f
	}

	kdSlider := widget.NewSlider(0, 1.0)
	kdSlider.Step = 0.01
	kdSlider.OnChanged = func(f float64) {
		m.setState(
			func(state *gofill.State) {
				state.Light.Kd = f
			})
		kdSlider.Value = f
	}

	mSlider := widget.NewSlider(0, 100.0)
	mSlider.Step = 0.5
	mSlider.OnChanged = func(f float64) {
		m.setState(
			func(state *gofill.State) {
				state.Light.M = f
			})
		mSlider.Value = f
	}

	modesRadio.Horizontal = true
	modesRadio.SetSelected(modes[0].Name())
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle("Light", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:      true,
		}),
		colorButton,
		modesRadio,
		widget.NewLabel("ks[0,1]"),
		ksSlider,
		widget.NewLabel("kd[0,1]"),
		kdSlider,
		widget.NewLabel("m[1,100]"),
		mSlider,
	)
}

