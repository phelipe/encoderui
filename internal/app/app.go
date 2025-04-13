package app

import "github.com/phelipe/encoderui/internal/encoders"

type Aplication struct {
	encoderList    map[string]Encoder
	currentEncoder Encoder
}

func New() *Aplication {
	return &Aplication{
		encoderList: map[string]Encoder{
			"base64": encoders.Base64{},
		},
		currentEncoder: encoders.Base64{},
	}
}

func (app *Aplication) SetEncoder(name string) bool {
	if data, exist := app.encoderList[name]; exist {
		app.currentEncoder = data
		return true
	}
	return false
}
