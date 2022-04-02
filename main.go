package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"golang-2d-platformer/colors"
	"golang-2d-platformer/inputs"
	"golang-2d-platformer/levels"
	"golang-2d-platformer/logger"
	"golang-2d-platformer/settings"
	"log"
	"os"
	"time"
)

const FPS = 1000 / 60

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		logger.Error("sdl unable to init: %s", err.Error())
		os.Exit(1)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow(settings.WINDOW_TITLE, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, settings.SCREEN_WIDTH, settings.SCREEN_HEIGHT, sdl.WINDOW_SHOWN)

	if err != nil {
		logger.Error("sdl unable to create window: %s", err.Error())
		os.Exit(1)
	}

	defer OnDestroy(window)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		logger.Error("sdl unable to create renderer: %s", err.Error())
		os.Exit(1)
	}

	defer OnDestroy(renderer)

	level := levels.New(settings.GetLevelMap(), settings.TILE_SIZE)

	previousTime := sdl.GetTicks64()
	frameTime := sdl.GetTicks64() - previousTime

	running := true

	for running {
		running = inputs.Listen()

		err = renderer.SetDrawColor(colors.Black().RGBA())
		if err != nil {
			logger.Error("sdl unable to draw background color: %s", err.Error())
			os.Exit(1)
		}

		err = renderer.Clear()
		if err != nil {
			logger.Error("sdl unable to clear renderer: %s", err.Error())
			os.Exit(1)
		}

		if running && inputs.KeyPressed(sdl.SCANCODE_ESCAPE) {
			running = false
		}

		setWindowTitle(window)

		if level.Restart() {
			level = levels.New(settings.GetLevelMap(), settings.TILE_SIZE)
		}

		err = level.Run(renderer)
		if err != nil {
			logger.Error("error on run level: %s", err.Error())
			os.Exit(1)
		}

		frameTime = sdl.GetTicks64() - previousTime
		previousTime = sdl.GetTicks64()

		if FPS > frameTime {
			delay := uint32(FPS - frameTime)
			sdl.Delay(delay)
		}

		renderer.Present()
	}
}

type DestroyFunc interface {
	Destroy() error
}

func OnDestroy(f DestroyFunc) {
	if f != nil {
		err := f.Destroy()
		if err != nil {
			log.Printf("error on destroy %T: %s", f, err.Error())
		}
	}
}

func setWindowTitle(window *sdl.Window) {
	title := settings.WINDOW_TITLE
	title += " | "

	title += time.Now().Format("2006-01-02 15:04:05")
	title += " | "

	cursor := inputs.GetCursor()
	title += "MX=" + fmt.Sprint(cursor.X)
	title += " | "
	title += "MY=" + fmt.Sprint(cursor.Y)

	window.SetTitle(title)
}
