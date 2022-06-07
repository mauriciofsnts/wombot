package challenges

import (
	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/discord/events"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
)

func init() {
	events.RegisterSlashCommand(
		&events.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "submission",
				Description: "send a challenge",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionAttachment,
						Name:        "image",
						Description: "Send the development of your challenge as an image!",
						Required:    true,
					},
				},
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

				args := ctx.I.Interaction.ApplicationCommandData().Options

				if args[0] == nil {
					logger.Debug("File not sent/found")
					return
				}

				logger.Debug("Content: ", ctx.I.Interaction.Message)
				logger.Debug("Type:", ctx.I.Data)

				var guild = &entities.Guild{
					GuildID: ctx.I.GuildID,
				}

				err := repos.Guild.Find(guild)

				logger.Debug(ctx.I)

				if err != nil {
					logger.Error("Guild not found", err)
					return
				}

				var user = entities.User{
					UserID: ctx.I.Member.User.ID,
				}

				err = repos.User.Find(&user)

				if err != nil {
					logger.Error("User not found/registered ", err)
				}

				var submission = entities.Submission{
					UserId:     user.UserID,
					GuildId:    guild.GuildID,
					CurrentDay: 1,
					MessageId:  "example",
				}

				err = repos.Submission.Create(submission)

				if err != nil {
					logger.Error("A problem occurred to the registration for submission", err)
					return
				}

				ctx.Ok(&discordgo.MessageEmbed{
					Title: "Parabéns pelo envio!",
				})

			},
		},
	)
}
