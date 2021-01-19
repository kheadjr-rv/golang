package language

import "testing"

func Test_english_SayHello(t *testing.T) {
	tests := []struct {
		name string
		e    *english
		want string
	}{
		{name: "english", e: &english{}, want: "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &english{}
			if got := e.SayHello(); got != tt.want {
				t.Errorf("english.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
