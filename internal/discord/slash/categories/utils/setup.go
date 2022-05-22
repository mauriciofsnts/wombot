package utils

import (
	"fmt"

	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/utils/reply"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
)

var (
	// :angry
	permission int64 = discordgo.PermissionManageServer
	// :angry again
	min, max = 0.0, 23.0
)

var Setup = &slash.SlashCommand{

	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:                     "setup",
		Description:              "Setup the bot",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "challenge-channel",
				Description: "Set the channel for challenges",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "hour",
				Description: "Set a day hour",
				Required:    true,
				MinValue:    &min,
				MaxValue:    max,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		args := i.Interaction.ApplicationCommandData().Options

		channel := args[0].ChannelValue(s)
		hour := args[1].IntValue()

		if channel.Type != discordgo.ChannelTypeGuildText {
			reply.Error(s, i, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("Channel <#%s> is not a text channel", channel.ID),
			})
			return
		}

		err := repos.Guild.Create(entities.Guild{
			GuildID:    i.GuildID,
			ChannelID:  channel.ID,
			CurrentDay: 0,
			HourOfDay:  hour,
		})

		if err != nil {
			logger.Error(err)

			reply.Error(s, i, &discordgo.MessageEmbed{
				Title:       "An error occurred when saving the data",
				Description: "Robots are crazy",
				Image: &discordgo.MessageEmbedImage{
					URL: "https://media4.giphy.com/media/l46CwEYnbFtFfjZNS/giphy.gif?cid=ecf05e47gkz2ncyxh0bcdfeezwr50ppd8f7wapxs0qd4c4xj&rid=giphy.gif&ct=g",
				},
			})
			return
		}

		reply.Ok(s, i, &discordgo.MessageEmbed{
			Title:       "Setup complete!",
			Description: fmt.Sprintf("Now you will receive the challenges on the <#%s> at %d:00 :)", channel.ID, hour),
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media4.giphy.com/media/mCIjCgs3nWQWfJZvPA/giphy.gif?cid=ecf05e47565hfcdquq8ypqog4topsoelvgbayyk0yl182um9&rid=giphy.gif&ct=g",
			},
		})

	},
}
