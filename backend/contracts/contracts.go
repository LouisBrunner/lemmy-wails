package contracts

import "context"

type App interface {
	OnStartup(ctx context.Context)
	Bindings() Bindings
}

type Bindings interface {
	Greet(name string) string
}
