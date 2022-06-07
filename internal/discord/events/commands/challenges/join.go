package challenges

import (
	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/discord/events"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/bwmarrin/discordgo"
)

func init() {
	events.RegisterSlashCommand(
		&events.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "join",
				Description: "Join a challenge",
			},
			Handler: func(ctx *events.DiscordContext, t *i18n.Language) {

				if ctx.I.GuildID == "" {
					ctx.Error(&discordgo.MessageEmbed{
						Title:       t.Errors.Title.Str(),
						Description: "Comando é apenas utilizado em guildas",
						Image: &discordgo.MessageEmbedImage{
							URL: t.Errors.GenericGif.Str(),
						},
					})

					return
				}

				var user = &entities.User{
					UserID: ctx.I.Member.User.ID,
				}

				err := repos.User.Find(user)

				if err == nil {

					ctx.Error(&discordgo.MessageEmbed{
						Title:       t.Errors.Title.Str(),
						Description: t.Errors.AlreadyRegistered.Str(),
						Image: &discordgo.MessageEmbedImage{
							URL: t.Errors.AlreadyRegisteredGif.Str(),
						},
					})

					return
				}

				err = repos.User.Create(entities.User{
					UserID:        ctx.I.Member.User.ID,
					GuildID:       ctx.I.GuildID,
					Streak:        0,
					HighestStreak: 0,
					CurrentDay:    0,
				})

				if err != nil {

					ctx.Error(&discordgo.MessageEmbed{
						Title:       t.Errors.Title.Str(),
						Description: t.Errors.Generic.Str(),
						Image: &discordgo.MessageEmbedImage{
							URL: t.Errors.GenericGif.Str(),
						},
					})

					return
				}

				ctx.Ok(&discordgo.MessageEmbed{
					Title:       t.Commands.Join.Title.Str(),
					Description: t.Commands.Join.Response.Str(),
					Image: &discordgo.MessageEmbedImage{
						URL: t.Commands.Join.Gif.Str(),
					},
				})

			},
		},
	)
}
