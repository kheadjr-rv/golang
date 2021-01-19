package language

import "testing"

func Test_spanish_SayHello(t *testing.T) {
	tests := []struct {
		name string
		s    *spanish
		want string
	}{
		{name: "spanish", s: &spanish{}, want: "Hola"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &spanish{}
			if got := s.SayHello(); got != tt.want {
				t.Errorf("spanish.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
