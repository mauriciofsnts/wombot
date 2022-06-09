package challenges

import (
	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/i18n"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
)

func init() {

	slash.RegisterSlashCommand(
		&slash.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "send",
				Description: "Send a challenge",
			},
			Handler: func(ctx *slash.DiscordContext, t *i18n.Language) {
				if ctx.I.GuildID == "" {
					ctx.Error(&discordgo.MessageEmbed{
						Title:       t.Errors.Title.Str(),
						Description: "Comando é apenas utilizado em servidores",
						Image: &discordgo.MessageEmbedImage{
							URL: t.Errors.GenericGif.Str(),
						},
					})

					return
				}

				var guild = &entities.Guild{
					GuildID: ctx.I.GuildID,
				}

				err := repos.Guild.Find(guild)

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

				channel, err := ctx.S.UserChannelCreate(ctx.I.Member.User.ID)

				if err != nil {
					logger.Error("Error creating channel", err)
				}

				ctx.S.ChannelMessageSend(channel.ID, "*Hello :)*")

				ctx.Ok(&discordgo.MessageEmbed{
					Title:       "Hehe",
					Description: "",
				})

			},
		},
	)

}
