package plane

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"github.com/areknoster/gofill/pkg/gofill"
	"github.com/areknoster/gofill/pkg/normde"
	"github.com/areknoster/gofill/pkg/raster"
)

type Plane struct {
	widget.BaseWidget
	renderer gofill.Renderer
	mode gofill.PlaneMode
	size          fyne.Size
}

func NewPlane(renderer gofill.Renderer, size fyne.Size) (*Plane, func(mode gofill.PlaneMode)) {
	p := &Plane{
		size:          size,
		renderer: renderer,
	}
	p.ExtendBaseWidget(p)
	return p, func(mode gofill.PlaneMode){
		p.mode = mode
	}
}

var (
	_ fyne.Widget    = &Plane{}
	_ gofill.Plane   = &Plane{}
	_ fyne.Draggable = &Plane{}
	_ fyne.Tappable  = &Plane{}
)

func (p *Plane) Size() fyne.Size{
	return p.size
}

func (p *Plane) MinSize() fyne.Size{
	return p.size
}

func (p *Plane) Tapped(event *fyne.PointEvent) {
	normPt := normde.NormPoint2D(
		raster.Pixel{X: event.Position.X, Y: event.Position.Y},
		p.size.Width, p.size.Height)
	p.mode.HandleClick(normPt)
	logrus.Debugf("Tapped: %v", normPt)
}

func (p *Plane) Dragged(event *fyne.DragEvent) {
	start := normde.NormPoint2D(
		raster.Pixel{
			X: event.Position.X-event.DraggedX,
			Y: event.Position.Y-event.DraggedY,
		},
		p.size.Width,
		p.size.Height)

	vec := normde.NormVector2D(
		raster.Pixel{X: event.DraggedX, Y: event.DraggedY},
		p.size.Width, p.size.Height)
	p.mode.HandleDrag(start, vec)
	logrus.Debugf("Drag: start: %v, vec: %v", start, vec)
}

func (p *Plane) DragEnd() {
	logrus.Debugf("Drag finished")
	p.mode.HandleDragEnd()
}


func (p *Plane) Refresh(){
	p.BaseWidget.Refresh()
}


func (p *Plane) CreateRenderer() fyne.WidgetRenderer {
	return NewPlaneRenderer(p)
}



