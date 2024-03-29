package ui

import (
	"image"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/images"
	"github.com/areknoster/gofill/pkg/normde"
)

func (m *Menu) newTextureSection() fyne.CanvasObject {
	local := images.ChooseTexture
	available := local.ListAvailable()
	onSelect := func(s string) {
		err := local.Set(
			s, func(rgba *image.RGBA) {
				logrus.Debugf("selecting local image: %s", s)
				m.setState(
					func(state *gofill.State) {
						state.Texture = normde.RGBAImageToSizedVectorMap(rgba, state.Size)
					})
			})
		if err != nil {
			logrus.Errorf("could not set local image: %s", err.Error())
		}
	}
	selectLocal := widget.NewSelect(
		available, onSelect)
	selectLocal.SetSelected("moro")
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle(
			"Select texture", fyne.TextAlignCenter, fyne.TextStyle{
				Bold: true,
			}),
		selectLocal,
	)
}

func (m *Menu) newNormalMapSelection() fyne.CanvasObject {
	local := images.ChooseNormalMap
	available := local.ListAvailable()
	onSelect := func(s string) {
		err := local.Set(
			s, func(rgba *image.RGBA) {
				logrus.Debugf("selecting local image: %s", s)
				m.setState(
					func(state *gofill.State) {
						state.NormalMap = normde.RGBAToSizedNormMap(rgba, state.Size)
					})
			})
		if err != nil {
			logrus.Errorf("could not set local image: %s", err.Error())
		}
	}
	selectLocal := widget.NewSelect(
		available, onSelect)
	selectLocal.SetSelected("bricks")

	uniform := widget.NewButton(
		"uniform", func() {
			m.setState(
				func(state *gofill.State) {
					state.NormalMap = normde.NewUniform(state.Size)
				})
		})


	waveSlider := widget.NewSlider(10, 10000)
	waveSlider.Step = 1
	waveSlider.OnChanged = func(f float64) {
		m.setState(
			func(state *gofill.State) {
				state.WavesCoef = f
			})
		waveSlider.Value = f
	}
	waveSlider.SetValue(100)


	wave := widget.NewButton(
		"waves", func() {
			m.setState(
				func(state *gofill.State) {
					state.NormalMap = normde.NewWave(state.Size, state.WavesCoef )
				})
		})

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle(
			"Select normal map", fyne.TextAlignCenter, fyne.TextStyle{
				Bold: true,
			}),
		selectLocal,
		uniform,
		wave,
		waveSlider,
	)
}
