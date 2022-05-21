package bootstrap

import (
	"code.db.cafe/wombot/internal/config"
	"code.db.cafe/wombot/internal/discord"
	"github.com/Pauloo27/logger"
)

func Start() {
	err := config.LoadConfig()
	logger.HandleFatal(err, "Failed to load config")

	logger.HandleFatal(discord.Start(), "Failed to start discord")

}
