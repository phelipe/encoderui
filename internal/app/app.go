package app

import "github.com/phelipe/gcode/internal/encoders"

type App struct {
	encoderList    map[string]Encoder
	currentEncoder Encoder
}

func New() *App {
	return &App{
		encoderList: map[string]Encoder{
			"base64": encoders.Base64{},
		},
		currentEncoder: encoders.Base64{},
	}
}

func (app *App) SetEncoder(name string) bool {
	if data, exist := app.encoderList[name]; exist {
		app.currentEncoder = data
		return true
	}
	return false
}
