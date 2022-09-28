package contracts

import "github.com/LouisBrunner/lemmy/library/backend/api"

type ConfigInternal struct {
	LastWidth  int
	LastHeight int
	LastX      int
	LastY      int
}

type ConfigManager[User any] interface {
	api.ConfigManager[User]
	Internal() ConfigInternal
	SaveInternal(config ConfigInternal) error
}
