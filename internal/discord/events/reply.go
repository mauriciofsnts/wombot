package events

import "github.com/bwmarrin/discordgo"

func (ctx *DiscordContext) Reply(embed *discordgo.MessageEmbed) {
	ctx.S.InteractionRespond(
		ctx.I.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		},
	)
}

func (ctx *DiscordContext) Error(embed *discordgo.MessageEmbed) {
	embed.Color = 0xe33e32
	ctx.Reply(embed)
}

func (ctx *DiscordContext) Ok(embed *discordgo.MessageEmbed) {
	embed.Color = 0x42f54b
	ctx.Reply(embed)
}
