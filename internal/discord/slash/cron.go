package slash

import (
	"fmt"
	"time"

	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func StartCron(session *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(50).Second().Do(func() {
		var guilds []entities.Guild

		err := repos.Guild.FindAll(&guilds)

		if err != nil {
			logger.Error("Failed to load guilds", err)
			return
		}

		logger.Success("Guilds", guilds)

		for _, guild := range guilds {
			session.ChannelMessageSend(guild.ChannelID, fmt.Sprintf("Dia: %d", guild.CurrentDay))
		}

	})

	s.StartAsync()
}
