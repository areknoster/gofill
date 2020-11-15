package ui

import (
	"image"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/images"
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
						state.Texture = rgba
					})
			})
		if err != nil {
			logrus.Errorf("could not set local image: %s", err.Error())
		}
	}
	selectLocal := widget.NewSelect(
		available, onSelect)
	selectLocal.SetSelected(available[0])
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle("Select texture", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:      true,
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
						state.NormalMap = rgba
					})
			})
		if err != nil {
			logrus.Errorf("could not set local image: %s", err.Error())
		}
	}
	selectLocal := widget.NewSelect(
		available, onSelect)
	selectLocal.SetSelected(available[0])
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle("Select normal map", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:      true,
		}),
		selectLocal,
	)
}