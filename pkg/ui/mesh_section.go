package ui

import (
	"errors"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/geom2d"
	"github.com/areknoster/gofill/pkg/gofill"
)

func (m *Menu)newMeshSection()fyne.CanvasObject{
	validator := func(s string) error {
		v, err := strconv.ParseInt(s, 10, 16)
		if v < 1 || v > 20{
			return errors.New("wrong value: v must be in [1,20]")
		}
		return err
	}

	xNum:= widget.NewEntry()
	xNum.Text= "6"
	xNum.Validator = validator
	xNum.OnChanged =  func(s string) {
		err := xNum.Validate()
		if err != nil{
			logrus.Errorf("cannot set X mesh value: %s",err.Error())
			return
		}
		v, _ := strconv.ParseInt(s, 10, 16)
		m.setState(
			func(state *gofill.State) {
				logrus.Debugf("setting mesh to %dx%d", int(v), state.Mesh.Y)
				state.Mesh = geom2d.NewMesh(int(v), state.Mesh.Y)
			})
	}


	yNum:= widget.NewEntry()
	yNum.Text="6"
	yNum.Validator = validator
	yNum.OnChanged =  func(s string) {
		err := yNum.Validate()
		if err != nil{
			logrus.Errorf("cannot set X mesh value: %s",err.Error())
			return
		}
		v, _ := strconv.ParseInt(s, 10, 16)
		m.setState(
			func(state *gofill.State) {
				logrus.Debugf("setting mesh to %dx%d", state.Mesh.X, int(v))

				state.Mesh = geom2d.NewMesh(state.Mesh.X, int(v))
			})
	}

	showCheck := widget.NewCheck("show", func(b bool) {
		m.setState(
			func(state *gofill.State) {
				state.ShowMesh = b
			})
	})
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabelWithStyle("Mesh[X,Y]", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:      true,
		}),
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
			xNum,
			yNum,
			showCheck,
		),
	)
}
