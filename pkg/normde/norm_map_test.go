package normde

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/areknoster/gofill/pkg/geom3d"
)

func TestRGBAToNormVersor(t *testing.T) {
	type args struct {
		rgba color.RGBA
	}
	tests := []struct {
		name string
		args args
		want geom3d.Vector
	}{
		{
			name: "example from webpage",
			args: args{
				rgba: color.RGBA{127, 127, 255, 0},
			},
			want: geom3d.Vector{0, 0, 1.0},

		},
		{
			name: "full red half green should be 1,0,0",
			args: args{
				rgba: color.RGBA{255,127,0,0},
			},
			want: geom3d.Vector{1,0,0},
		},
		{
			name: "non-red and the others nullified should be -1,0,0",
			args: args{
				rgba: color.RGBA{0,127,0,0},
			},
			want: geom3d.Vector{-1,0,0},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := RGBAToNormVersor(tt.args.rgba); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RGBAToNormVersor() = %v, want %v", got, tt.want)
				}
			})
	}
}
