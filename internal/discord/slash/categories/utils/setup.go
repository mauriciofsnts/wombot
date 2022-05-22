package utils

import (
	"strconv"

	"code.db.cafe/wombot/internal/discord/slash"
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

		args = args

		// save the day goal on bd

		// save the challenge channel on bd

		s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Setup complete!",
							Description: "O setup está quente!",
							Color:       0x42f54b,
							Image: &discordgo.MessageEmbedImage{
								URL: "https://acegif.com/wp-content/gifs/race-car-11.gif",
							},
							Thumbnail: &discordgo.MessageEmbedThumbnail{
								URL: "https://brandlogos.net/wp-content/uploads/2013/06/uefa-champions-league-eps-vector-logo.png",
							},
							Author: &discordgo.MessageEmbedAuthor{
								Name:    "Wombot",
								IconURL: "https://sm.ign.com/ign_br/cover/j/john-wick-/john-wick-chapter-4_129x.jpg",
							},
							Footer: &discordgo.MessageEmbedFooter{
								Text:    "que pezinho uwu",
								IconURL: "https://i.pinimg.com/originals/d3/d1/75/d3d175e560ae133f1ed5cd4ec173751a.png",
							},
							Fields: []*discordgo.MessageEmbedField{
								{
									Name:  "Day goal",
									Value: strconv.Itoa(int(args[0].IntValue())),
								},
								{
									Name:  "Channel id :v",
									Value: args[1].ChannelValue(s).Name,
								},
							},
						},
					},
				},
			},
		)

	},
}
