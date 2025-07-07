package main

import (
	"github.com/Shun0212/opencode/cmd"
	"github.com/Shun0212/opencode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
