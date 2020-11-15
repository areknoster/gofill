package render

import (
	"fmt"
	"image/color"
	"math"
	"testing"
)

func TestExample(t *testing.T){
	rgba := color.RGBA{}
	rgba.R = uint8(math.Pow(2, 10))
	fmt.Println(rgba)
}
