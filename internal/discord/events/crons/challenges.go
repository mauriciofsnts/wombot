package crons

import (
	"fmt"

	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/i18n"
	"code.db.cafe/wombot/internal/utils/challenges"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func Challenges(s *gocron.Scheduler, session *discordgo.Session) {

	s.Every(1).Day().Do(func() {
		var guilds []entities.Guild

		err := repos.Guild.FindAll(&guilds)

		if err != nil {
			logger.Error("Failed to load guilds", err)
			return
		}

		for _, guild := range guilds {
			t := i18n.GetLanguage(guild.Language)

			session.ChannelMessageSendEmbed(guild.ChannelID, &discordgo.MessageEmbed{
				Title:       t.Challenges.Title.Str(),
				Description: t.Challenges.Description.Str(fmt.Sprintf("#%d", guild.CurrentDay), challenges.ChallengesData[guild.CurrentDay].Name),
				Image: &discordgo.MessageEmbedImage{
					URL: challenges.ChallengesData[guild.CurrentDay].Image,
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: t.Challenges.Footer.Str(),
				},
				Color: 0x0bf6f6,
			})

		}

	})

}
