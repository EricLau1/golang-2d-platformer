package logger

import "github.com/veandco/go-sdl2/sdl"

func Info(format string, values ...interface{}) {
	sdl.LogInfo(sdl.LOG_CATEGORY_APPLICATION, format, values...)
}

func Error(format string, values ...interface{}) {
	sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, format, values...)
}
