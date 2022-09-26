package bindings

import (
	"fmt"

	"github.com/LouisBrunner/lemmy/backend/contracts"
)

type bindings struct{}

func New() contracts.Bindings {
	return &bindings{}
}

func (me *bindings) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
