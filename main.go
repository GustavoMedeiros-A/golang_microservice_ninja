package main

import (
	"context"
	"fmt"

	"github.com/GustavoMedeiros-A/golang_microservice_ninja/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start application", err)
	}
}
