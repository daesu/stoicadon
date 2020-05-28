package main

import (
	"github.com/daesu/stoicadon/api"
	"github.com/sirupsen/logrus"
)

func main() {
	application, err := api.ConfigureApplication()
	if err != nil {
		logrus.Fatalf("error: %w", err)
	}

	if err := api.StartAPI(application); err != nil {
		logrus.Fatalf("could not start api: %w", err)
	}
}
