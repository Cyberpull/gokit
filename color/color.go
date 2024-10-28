package color

import "github.com/fatih/color"

type Color struct {
	color.Color
}

func New(value ...Attribute) *Color {
	attrs := a(value...)
	return &Color{Color: *color.New(attrs...)}
}

func NewRGB(r int, g int, b int) *Color {
	return &Color{Color: *color.RGB(r, g, b)}
}

func NewBgRGB(r int, g int, b int) *Color {
	return &Color{Color: *color.BgRGB(r, g, b)}
}
