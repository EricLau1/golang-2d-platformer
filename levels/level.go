package levels

import (
	"github.com/veandco/go-sdl2/sdl"
	"golang-2d-platformer/objects"
	"golang-2d-platformer/players"
	"golang-2d-platformer/settings"
)

const GRAVITY = 0.8

type Level interface {
	Run(renderer *sdl.Renderer) error
	Restart() bool
}

type level struct {
	tiles      []*objects.Object
	player     *players.Player
	restart    bool
	worldShift float32
}

func New(levelMap []string, tileSize int) Level {
	var tiles []*objects.Object
	var player *players.Player

	for row := 0; row < len(levelMap); row++ {
		for col := 0; col < len(levelMap[row]); col++ {

			cell := levelMap[row][col]
			x := float32(col * tileSize)
			y := float32(row * tileSize)

			if cell == 'T' {
				tiles = append(tiles, objects.New(tileSize, tileSize, x, y))
			} else if cell == 'P' {
				player = players.New(tileSize/2, tileSize/2, x, y)
			}
		}
	}

	return &level{tiles: tiles, player: player}
}

func (l *level) Run(renderer *sdl.Renderer) error {
	l.player.Update()

	l.handleHorizontalCollision()
	l.handleVerticalCollision()
	l.handleCamera()

	var err error

	for _, tile := range l.tiles {
		tile.Position.X += l.worldShift
		err = tile.Draw(renderer)
		if err != nil {
			return err
		}
	}

	return l.player.Draw(renderer)
}

func (l *level) applyGravity() {
	l.player.Direction.Y += GRAVITY
	l.player.Position.Y += l.player.Direction.Y
}

func (l *level) handleHorizontalCollision() {
	playerRect := l.player.Rect()

	for _, tile := range l.tiles {
		tileRect := tile.Rect()

		if playerRect.HasIntersection(&tileRect) {

			if l.player.Direction.X > 0 {
				l.player.Position.X = float32(tileRect.X - playerRect.W)
				l.player.Direction.X = 0
			} else if l.player.Direction.X < 0 {
				l.player.Position.X = float32(tileRect.X + tileRect.W)
				l.player.Direction.X = 0
			}
		}
	}
}

func (l *level) handleVerticalCollision() {
	l.applyGravity()

	playerRect := l.player.Rect()

	for _, tile := range l.tiles {
		tileRect := tile.Rect()

		if playerRect.HasIntersection(&tileRect) {

			if l.player.Direction.Y > 0 {
				l.player.Position.Y = float32(tileRect.Y - playerRect.H)
				l.player.Direction.Y = 0
				l.player.OnGround = true
			} else if l.player.Direction.Y < 0 {
				l.player.Position.Y = float32(tileRect.Y + tileRect.H)
				l.player.Direction.Y = 0
			}
		}
	}

	if (l.player.OnGround && l.player.Direction.Y < 0) || l.player.Direction.Y > 1 {
		l.player.OnGround = false
	}

	if l.player.Position.Y > settings.SCREEN_HEIGHT {
		l.restart = true
	}
}

func (l *level) handleCamera() {
	if l.player.Position.X < settings.SCREEN_WIDTH/4 && l.player.Direction.X < 0 {
		l.worldShift = players.SPEED
		l.player.Speed = 0
	} else if l.player.Position.X > settings.SCREEN_WIDTH-(settings.SCREEN_WIDTH/4) && l.player.Direction.X > 0 {
		l.worldShift = -players.SPEED
		l.player.Speed = 0
	} else {
		l.worldShift = 0
		l.player.Speed = players.SPEED
	}
}

func (l *level) Restart() bool {
	return l.restart
}
