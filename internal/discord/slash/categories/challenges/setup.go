package challenges

import (
	"fmt"

	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"

	"code.db.cafe/wombot/internal/database/entities"
	"code.db.cafe/wombot/internal/database/repos"
	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/i18n"
)

var (
	permission int64 = discordgo.PermissionManageServer
	min, max         = 0.0, 23.0
)

func init() {
	slash.RegisterSlashCommand(
		&slash.SlashCommand{

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
			Handler: func(ctx *slash.DiscordContext, t *i18n.Language) {

				args := ctx.I.Interaction.ApplicationCommandData().Options

				channel := args[0].ChannelValue(ctx.S)
				hour := args[1].IntValue()

				if channel.Type != discordgo.ChannelTypeGuildText {
					ctx.Error(&discordgo.MessageEmbed{
						Title:       t.Errors.Title.Str(),
						Description: t.Errors.NotATextChannel.Str(fmt.Sprintf("<#%s>", channel.ID)),
					})
					return
				}

				err := repos.Guild.Create(entities.Guild{
					GuildID:    ctx.I.GuildID,
					ChannelID:  channel.ID,
					CurrentDay: 0,
					HourOfDay:  hour,
				})

				if err != nil {
					logger.Error(err)

					ctx.Error(&discordgo.MessageEmbed{
						Title:       t.Errors.Title.Str(),
						Description: t.Errors.ToSave.Str(),
						Image: &discordgo.MessageEmbedImage{
							URL: t.Errors.ToSaveGif.Str(),
						},
					})
					return
				}

				ctx.Ok(&discordgo.MessageEmbed{
					Title:       t.Commands.Setup.Title.Str(),
					Description: t.Commands.Setup.Response.Str(fmt.Sprintf("<#%s>", channel.ID), hour),
					Image: &discordgo.MessageEmbedImage{
						URL: t.Commands.Setup.Gif.Str(),
					},
				})

			},
		},
	)
}
