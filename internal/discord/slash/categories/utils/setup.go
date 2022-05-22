package utils

import (
	"fmt"

	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/utils/reply"
	"github.com/bwmarrin/discordgo"
)

// :angry
var permission int64 = discordgo.PermissionManageServer

var Setup = &slash.SlashCommand{

	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:                     "setup",
		Description:              "Setup the bot",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "day-goal",
				Description: "Set a day goal",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "challenge-channel",
				Description: "Set the channel for challenges",
				Required:    true,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		args := i.Interaction.ApplicationCommandData().Options

		channel := args[1].ChannelValue(s)

		if channel.Type != discordgo.ChannelTypeGuildText {
			reply.Error(s, i, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("Channel <#%s> is not a text channel", channel.ID),
			})

			return
		}

		reply.Ok(s, i, &discordgo.MessageEmbed{
			Title:       "Setup complete!",
			Description: fmt.Sprintf("Now you will receive the challenges on the <#%s> :)", channel.ID),
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media4.giphy.com/media/mCIjCgs3nWQWfJZvPA/giphy.gif?cid=ecf05e47565hfcdquq8ypqog4topsoelvgbayyk0yl182um9&rid=giphy.gif&ct=g",
			},
		})

	},
}
