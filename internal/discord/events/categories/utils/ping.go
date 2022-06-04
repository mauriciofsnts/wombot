package utils

import (
	"code.db.cafe/wombot/internal/discord/events"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/bwmarrin/discordgo"
)

func init() {
	events.RegisterSlashCommand(
		&events.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping the bot",
			},
			Handler: func(ctx *events.DiscordContext, t *i18n.Language) {

				ctx.Ok(&discordgo.MessageEmbed{
					Title:       t.Commands.Ping.Title.Str(),
					Description: t.Commands.Ping.Response.Str(),
					Image: &discordgo.MessageEmbedImage{
						URL: t.Commands.Ping.Gif.Str(),
					},
				})

			},
		})
}
