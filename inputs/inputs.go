package inputs

import "github.com/veandco/go-sdl2/sdl"

type Input interface {
	Listen() bool
	KeyPressed(key sdl.Scancode) bool
	GetCursor() sdl.Point
}

type input struct {
	keyboard []uint8
	cursor   sdl.Point
}

func (i *input) Listen() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			return false

		case sdl.MOUSEMOTION:
			motion := event.(*sdl.MouseMotionEvent)
			i.cursor.X = motion.X
			i.cursor.Y = motion.Y

		case sdl.KEYDOWN:
			fallthrough
		case sdl.KEYUP:
			i.keyboard = sdl.GetKeyboardState()
		}
	}
	return true
}

func (i *input) KeyPressed(key sdl.Scancode) bool {
	return i.keyboard[key] == 1
}

func (i *input) GetCursor() sdl.Point {
	return i.cursor
}

func New() Input {
	return &input{
		keyboard: sdl.GetKeyboardState(),
		cursor:   sdl.Point{},
	}
}

var in Input

func init() {
	in = New()
}

func Listen() bool {
	return in.Listen()
}

func KeyPressed(key sdl.Scancode) bool {
	return in.KeyPressed(key)
}

func GetCursor() sdl.Point {
	return in.GetCursor()
}
