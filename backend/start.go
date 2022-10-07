package backend

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/lemmy-wails/backend/api"
	"github.com/LouisBrunner/lemmy-wails/backend/internal"
)

func Start[Bindings, Config any](opts api.Options[Bindings, Config]) {
	worker := internal.NewWorker(opts)
	err := worker.Work()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
