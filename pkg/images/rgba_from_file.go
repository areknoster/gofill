package images

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"strings"

	"fyne.io/fyne"
)

func RgbaFromFile(file fyne.URIReadCloser, set func(rgba *image.RGBA)) error {
	var img image.Image
	var err error
	defer file.Close()
	switch strings.ToLower(file.URI().Extension()){
	case  ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
		if err != nil {
			return fmt.Errorf("could not decode image: %w", err)
		}
	case ".png":
		img, err = png.Decode(file)
		if err != nil {
			return fmt.Errorf("could not decode image: %w", err)
		}
	default:
		return errors.New("not known format")
	}


	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)
	set(rgba)
	return nil
}


