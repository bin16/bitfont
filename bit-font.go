package bitfont

import (
	"image"

	"golang.org/x/image/font/basicfont"
)

type BitFont struct {
	Width, Height            int
	Advance, Ascent, Descent int
	Left                     int
	img                      image.Image
	bytes                    []byte
	ranges                   []basicfont.Range
}

func (u *BitFont) Mask() (mask *image.Alpha) {
	if len(u.bytes) > 0 {
		var (
			w = u.Width
			h = len(u.bytes) / w
		)
		mask = &image.Alpha{
			Pix:    u.bytes,
			Stride: u.Width,
			Rect:   image.Rect(0, 0, w, h),
		}

		return
	}

	mask = image.NewAlpha(u.img.Bounds())
	for i := 0; i < u.img.Bounds().Dx(); i++ {
		for r := 0; r < u.img.Bounds().Dy(); r++ {
			mask.Set(i, r, u.img.At(i, r))
		}
	}

	return mask
}

func (u *BitFont) Face() *basicfont.Face {
	return &basicfont.Face{
		Mask:    u.Mask(),
		Width:   u.Width,
		Height:  u.Height,
		Advance: u.Advance,
		Ascent:  u.Ascent,
		Descent: u.Descent,
		Ranges:  u.ranges,
		Left:    u.Left,
	}
}

type BitFontOpt func(u *BitFont)
type BitFontOptions struct{}

func (BitFontOptions) Width(d int) BitFontOpt {
	return func(u *BitFont) {
		u.Width = d
	}
}

func (BitFontOptions) Height(d int) BitFontOpt {
	return func(u *BitFont) {
		u.Height = d
	}
}

func (BitFontOptions) Advance(d int) BitFontOpt {
	return func(u *BitFont) {
		u.Advance = d
	}
}

func (BitFontOptions) Ascent(d int) BitFontOpt {
	return func(u *BitFont) {
		u.Ascent = d
	}
}

func (BitFontOptions) Descent(d int) BitFontOpt {
	return func(u *BitFont) {
		u.Descent = d
	}
}

func (BitFontOptions) Image(img image.Image) BitFontOpt {
	return func(u *BitFont) {
		u.img = img
	}
}

func (BitFontOptions) Bytes(data []byte) BitFontOpt {
	return func(u *BitFont) {
		u.bytes = data
	}
}

func (BitFontOptions) Left(d int) BitFontOpt {
	return func(u *BitFont) {
		u.Left = d
	}
}

func (BitFontOptions) Range(offset int, low rune, high ...rune) BitFontOpt {
	return func(u *BitFont) {
		var h = low + 1
		if len(high) > 0 {
			h = high[0]
		}

		u.ranges = append(u.ranges, basicfont.Range{
			Offset: offset,
			Low:    low,
			High:   h,
		})
	}
}

func (BitFontOptions) AddRune(offset int, r rune) BitFontOpt {
	return func(u *BitFont) {
		u.ranges = append(u.ranges, basicfont.Range{
			Offset: offset,
			Low:    r,
			High:   r + 1,
		})
	}
}

func (BitFontOptions) AddString(offset int, s string) BitFontOpt {
	return func(u *BitFont) {
		for i, r := range s {
			u.ranges = append(u.ranges, basicfont.Range{
				Offset: offset + i,
				Low:    r,
				High:   r + 1,
			})
		}
	}
}

func (BitFontOptions) Size(w, h int) BitFontOpt {
	return func(u *BitFont) {
		u.Width = w
		u.Height = h
	}
}

func (BitFontOptions) Face(advance, ascent, descent int) BitFontOpt {
	return func(u *BitFont) {
		u.Advance = advance
		u.Ascent = ascent
		u.Descent = descent
	}
}

func New(opts ...BitFontOpt) *BitFont {
	var bit = &BitFont{}
	for _, fn := range opts {
		fn(bit)
	}

	return bit
}

var Options BitFontOptions
