package gofill

import "image"

type ImageProvider interface{
	ListAvailable()[]string
	Set(name string, set func(rgba *image.RGBA)) error
}
