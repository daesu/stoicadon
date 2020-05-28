package main

import (
	"github.com/daesu/stoicadon/api"
	"github.com/sirupsen/logrus"
)

func main() {
	application, err := api.ConfigureApplication()
	if err != nil {
		logrus.Fatalf("error: %s", err.Error())
	}

	if err := api.StartAPI(application); err != nil {
		logrus.Fatalf("could not start api: %s", err.Error())
	}
}
