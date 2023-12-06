package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/GustavoMedeiros-A/golang_microservice_ninja/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start application", err)
	}

	cancel()
}
