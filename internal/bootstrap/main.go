package bootstrap

import (
	"code.db.cafe/wombot/internal/config"
	"github.com/Pauloo27/logger"
)

func Start() {
	logger.Info("Hello world :)")

	err := config.LoadConfig()

	logger.HandleFatal(err, "Failed to load config")

	logger.Success(config.Wombot.Token)
}
