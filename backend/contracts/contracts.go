package contracts

type Bindings interface {
	Greet(name string) string
}
