package objects

import (
	"github.com/veandco/go-sdl2/sdl"
	"golang-2d-platformer/colors"
)

type Object struct {
	Width    int32
	Height   int32
	Position sdl.FPoint
	Color    colors.Color
}

type Size interface {
	int | int32
}

func New[size Size](w, h size, x, y float32) *Object {
	return &Object{
		Width:    int32(w),
		Height:   int32(h),
		Position: sdl.FPoint{x, y},
		Color:    colors.Grey(),
	}
}

func (o *Object) Rect() sdl.Rect {
	x, y := int32(o.Position.X), int32(o.Position.Y)
	return sdl.Rect{
		X: x,
		Y: y,
		W: o.Width,
		H: o.Height,
	}
}

func (o *Object) Draw(renderer *sdl.Renderer) error {
	err := renderer.SetDrawColor(o.Color.RGBA())
	if err != nil {
		return err
	}
	rect := o.Rect()
	return renderer.FillRect(&rect)
}
