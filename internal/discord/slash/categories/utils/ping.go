package utils

import (
	"code.db.cafe/wombot/internal/discord/slash"
	"github.com/bwmarrin/discordgo"
)

var Ping = &slash.SlashCommand{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
}
