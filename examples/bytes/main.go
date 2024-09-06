package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/bin16/bitfont"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	pw            = 3
	ph            = 5
	letterSpacing = 1
	lineSpacing   = 1
	lineHeight    = ph + 1
)

func main() {

	var (
		x    = 10
		y    = 10
		text = []string{
			"23:59",
			"01234",
			"56789",
			":😊",
		}
	)

	var ft = bitfont.New(
		bitfont.Options.Size(pw, ph),
		bitfont.Options.Face(pw+letterSpacing, 0, ph),
		bitfont.Options.Bytes(digits),
		bitfont.Options.Range(0, '0', '0'+11),
		bitfont.Options.Range(11, '😊'),
	)
	var dst = image.NewRGBA(image.Rect(0, 0, 50, 50))
	var dr = font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.RGBA{255, 0, 0, 255}),
		Face: ft.Face(),
	}

	for i, s := range text {
		dr.Dot = fixed.P(x, y+lineHeight*i)
		dr.DrawString(s)
		dst.Set(x, y+lineHeight*i, color.RGBA{0, 255, 255, 255})
	}

	var file, _ = os.Create("image-b.png")
	png.Encode(file, dst)
}

var (
	F byte = 0xFF
)

var digits = []byte{

	// 0
	F, F, F,
	F, 0, F,
	F, 0, F,
	F, 0, F,
	F, F, F,

	// 1
	0, F, 0,
	F, F, 0,
	0, F, 0,
	0, F, 0,
	F, F, F,

	// 2
	F, F, 0,
	0, 0, F,
	0, F, 0,
	F, 0, 0,
	F, F, F,

	// 3
	F, F, 0,
	0, 0, F,
	0, F, 0,
	0, 0, F,
	F, F, 0,

	// 4
	F, 0, F,
	F, 0, F,
	F, F, F,
	0, 0, F,
	0, 0, F,

	// 5
	F, F, F,
	F, 0, 0,
	F, F, F,
	0, 0, F,
	F, F, F,

	// 6
	0, F, F,
	F, 0, 0,
	F, F, F,
	F, 0, F,
	F, F, F,

	// 7
	F, F, F,
	0, 0, F,
	0, F, 0,
	0, F, 0,
	0, F, 0,

	// 8
	F, F, F,
	F, 0, F,
	F, F, F,

	F, 0, F,
	F, F, F,

	// 9
	F, F, F,
	F, 0, F,
	F, F, F,
	0, 0, F,
	F, F, 0,

	// :
	0, 0, 0,
	0, F, 0,
	0, 0, 0,
	0, F, 0,
	0, 0, 0,

	// emoji 😊
	0, 0, 0,
	F, 0, F,
	0, 0, 0,
	F, 0, F,
	F, F, F,
}
