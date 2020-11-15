package images

import (
	"errors"
	"fmt"
	"image"

	"fyne.io/fyne/storage"

	"github.com/areknoster/gofill/pkg/gofill"
)

type LocalImages map[string] string

func (l LocalImages) ListAvailable() []string {
	list := make([]string, len(l))
	i := 0
	for name, _ := range l {
		list[i] = name
		i++
	}
	return list
}


func (l LocalImages) Set(name string, set func(rgba *image.RGBA)) error {
	path, ok := l[name]
	if !ok{
		return errors.New("resource not found")
	}
	uri := storage.NewFileURI(path)
	urc, err := storage.OpenFileFromURI(uri)
	if err != nil{
		return fmt.Errorf("could not open file %v, %w", uri, err )
	}
	err = RgbaFromFile(urc, set)
	if err != nil{
		return fmt.Errorf("could not convert file to bitmap: %w", err)
	}

	return nil
}

var _ gofill.ImageProvider = ChooseNormalMap

var ChooseNormalMap LocalImages = map[string]string{
	"bricks" : "resources/normal_bricks.jpg",
	"shapes" : "resources/normal_shapes.png",
	"all colors": "resources/normal_allcolors.jpeg",
}

var ChooseTexture LocalImages = map[string]string{
	"moro" : "resources/moro_texture.png",
	"red" : "resources/red_texture.jpeg",
	"kasia" : "resources/kasia.JPG",
}


