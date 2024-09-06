package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/bin16/bitfont"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	pw            = 10
	ph            = 12
	letterSpacing = 2
	lineSpacing   = 2
	lineHeight    = ph + letterSpacing
)

func main() {

	var (
		x    = 20
		y    = 20
		text = []string{
			"23:59",
			"01234",
			"56789",
			":",
		}
	)

	var ft = bitfont.New(
		bitfont.Options.Size(pw, ph),
		bitfont.Options.Face(pw+letterSpacing, 8, ph-8),
		bitfont.Options.Image(img),
		bitfont.Options.Range(0, '0', '0'+11),
		bitfont.Options.Left(-2),
	)
	var dst = image.NewRGBA(image.Rect(0, 0, 100, 100))
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

	var file, _ = os.Create("image.png")
	png.Encode(file, dst)
}

//go:embed digit-8x12-sheet.png
var raw []byte

var img, _ = png.Decode(bytes.NewReader(raw))
