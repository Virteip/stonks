package main

import (
	"fmt"
	"os"
)

const localEnvironment = "local"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = localEnvironment
	}

	var app application
	err := app.setup(env)
	if err != nil {
		return err
	}

	return app.startServer()
}
