package golang

// Greeter interface describes a greeters behaviour
type Greeter interface {
	SayHello() string
}

// Greet takes a greeter and prints its greeting.
func Greet(g Greeter) string {
	return g.SayHello()
}
