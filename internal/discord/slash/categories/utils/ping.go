package utils

import (
	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/bwmarrin/discordgo"
)

func init() {
	slash.RegisterSlashCommand(
		&slash.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping the bot",
			},
			Handler: func(ctx *slash.DiscordContext, t *i18n.Language) {

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
