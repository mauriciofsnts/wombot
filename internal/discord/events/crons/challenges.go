package crons

import (
	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func Challenges(s *gocron.Scheduler, session *discordgo.Session) {

	s.Every(5).Minutes().Do(func() {
		var guilds []entities.Guild

		err := repos.Guild.FindAll(&guilds)

		if err != nil {
			logger.Error("Failed to load guilds", err)
			return
		}

		for _, guild := range guilds {
			t := i18n.GetLanguage(guild.Language)

			msg, err := session.ChannelMessageSendEmbed(guild.ChannelID, &discordgo.MessageEmbed{
				Title:       t.Challenges.Title.Str(),
				Description: t.Challenges.Description.Str("#1", "Tela de login"),
				Image: &discordgo.MessageEmbedImage{
					URL: `https://cdn.dribbble.com/users/308682/screenshots/16316303/media/f9b4306971586e66bf77c5a63101e762.png?compress=1&resize=1200x900&vertical=top`,
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: t.Challenges.Footer.Str(),
				},
				Color: 0x0bf6f6,
			})

			if err != nil {
				logger.Error("Failed to create a message")
			}

			session.MessageReactionAdd(guild.ChannelID, msg.ID, "🤯")
			session.MessageReactionAdd(guild.ChannelID, msg.ID, "👍")
			session.MessageReactionAdd(guild.ChannelID, msg.ID, "👎")
		}

	})

}
