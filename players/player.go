package players

import (
	"github.com/veandco/go-sdl2/sdl"
	"golang-2d-platformer/colors"
	"golang-2d-platformer/inputs"
	"golang-2d-platformer/objects"
)

const (
	SPEED      = 8.0
	JUMP_FORCE = -16.0
)

type Player struct {
	*objects.Object
	Direction sdl.FPoint
	Speed     float32
	OnGround  bool
}

func New[size objects.Size](w, h size, x, y float32) *Player {
	player := new(Player)
	player.Object = objects.New(w, h, x, y)
	player.Speed = SPEED
	player.Color = colors.Red()
	return player
}

func (p *Player) Update() {
	p.Direction.X = 0

	if inputs.KeyPressed(sdl.SCANCODE_RIGHT) {
		p.Direction.X = 1
	} else if inputs.KeyPressed(sdl.SCANCODE_LEFT) {
		p.Direction.X = -1
	}

	if inputs.KeyPressed(sdl.SCANCODE_SPACE) && p.OnGround {
		p.Direction.Y = JUMP_FORCE
	}

	p.Position.X += p.Direction.X * p.Speed
}
