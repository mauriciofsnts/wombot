package bootstrap

import (
	"code.db.cafe/wombot/internal/config"
	"code.db.cafe/wombot/internal/database"
	"code.db.cafe/wombot/internal/discord"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/Pauloo27/logger"
)

func Start() {
	logger.HandleFatal(config.LoadConfig(), "Failed to load config")

	logger.HandleFatal(i18n.Start(), "Failed to load languages snif snif")

	logger.HandleFatal(database.Start(), "Failed to start database")
	logger.HandleFatal(discord.Start(), "Failed to start discord")

}
