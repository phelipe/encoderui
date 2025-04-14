package app

import "github.com/phelipe/encoderui/internal/encoders"

type Aplication struct {
	EncoderList    map[string]Encoder
	CurrentEncoder Encoder
}

func New() *Aplication {
	return &Aplication{
		EncoderList: map[string]Encoder{
			"base64": encoders.Base64{},
		},
		CurrentEncoder: encoders.Base64{},
	}
}

func (app *Aplication) SetEncoder(name string) bool {
	if data, exist := app.EncoderList[name]; exist {
		app.CurrentEncoder = data
		return true
	}
	return false
}
