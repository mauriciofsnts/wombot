package challenges

import (
	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/i18n"
	"code.db.cafe/wombot/internal/utils/reply"
	"github.com/bwmarrin/discordgo"
)

func init() {
	slash.RegisterSlashCommand(
		&slash.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "join",
				Description: "Join a challenge",
			},
			Handler: func(ctx *slash.DiscordContext, t *i18n.Language) {

				reply.Error(ctx.S, ctx.I, &discordgo.MessageEmbed{
					Title:       t.Errors.Generic.Str(),
					Description: "Example",
				})

			},
		},
	)
}
