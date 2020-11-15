package normde

import (
	"image"

	"fyne.io/fyne"

	"github.com/areknoster/gofill/pkg/geom3d"
)

func RGBAImageToSizedVectorMap(nonSized *image.RGBA, size fyne.Size) *geom3d.VectorMap{
	img := ResizeRGBA(nonSized, size)
	vm := geom3d.NewVectorMap(size.Width, size.Height)
	for i := 0; i < size.Width; i++ {
		for j := 0; j < size.Height; j++ {
			vm.SetAt(i, j, NormRGBAToVec(img.RGBAAt(i, j)))
		}
	}
	return vm
}
