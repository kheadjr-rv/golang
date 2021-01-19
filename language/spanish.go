package language

import "github.com/kheadjr-rv/golang"

type spanish struct{}

// Spanish returns a spanish greeter
func Spanish() golang.Greeter {
	return &spanish{}
}

func (s *spanish) SayHello() string {
	return "Hola"
}
