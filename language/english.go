package language

import (
	"github.com/kheadjr-rv/golang"
)

type english struct{}

// English returns a english greeter
func English() golang.Greeter {
	return &english{}
}

func (e *english) SayHello() string {
	return "Hello"
}
