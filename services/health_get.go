package services

import (
	"context"

	"github.com/daesu/stoicadon/models"
	"github.com/sirupsen/logrus"
)

func GetHealth(ctx context.Context) (*models.Health, error) {
	const funcName = "services.GetHealth"
	logrus.Infof("entered %s", funcName)

	health := models.Health{}

	return &health, nil
}
