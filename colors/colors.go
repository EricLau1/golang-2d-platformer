package colors

import "github.com/veandco/go-sdl2/sdl"

type Color interface {
	RGBA() (uint8, uint8, uint8, uint8)
}

type color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func New(red, green, blue uint8) Color {
	return &color{
		r: red,
		g: green,
		b: blue,
		a: sdl.ALPHA_TRANSPARENT,
	}
}

func (c *color) RGBA() (uint8, uint8, uint8, uint8) {
	return c.r, c.g, c.b, c.a
}

func Red() Color {
	return New(255, 0, 0)
}

func Green() Color {
	return New(0, 255, 0)
}

func Blue() Color {
	return New(0, 0, 255)
}

func Black() Color {
	return New(0, 0, 0)
}

func White() Color {
	return New(255, 255, 255)
}

func Grey() Color {
	return New(189, 189, 189)
}
